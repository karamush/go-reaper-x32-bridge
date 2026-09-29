# gen_tables.ps1
# Extracts the X32 OSC command catalogue (tables, enum string arrays, set arrays)
# from the C sources of the X32-Behringer repository (X32.c v0.88 + *.h tables,
# https://github.com/pmaillot/X32-Behringer) and generates:
#   docs/x32_commands.json                      - machine readable catalogue
#   docs/x32_enums.json                         - all enum string arrays
#   emulator/internal/x32/types_gen.go          - Type constants (parity with C enum types)
#   emulator/internal/x32/enums_gen.go          - enum string arrays
#   emulator/internal/x32/tables_gen.go         - command tables
#   emulator/internal/x32/sets_gen.go           - table sets (*set arrays)
#
# The C sources are NOT part of this repository (only the generated files are):
# clone the upstream project and point -Src at the checkout. Default is
# X32-Behringer\ in the repository root.
#
# Usage:  powershell -File tools\gen_tables.ps1 [-Src <checkout>] [-RepoRoot <path>]

[CmdletBinding()]
param([string]$RepoRoot = '', [string]$Src = '')

if (-not $RepoRoot) {
    $here = $PSScriptRoot
    if (-not $here) { $here = Split-Path -Parent $MyInvocation.MyCommand.Path }
    if (-not $here) { $here = (Get-Location).Path }
    $RepoRoot = (Resolve-Path (Join-Path $here '..')).Path
}

$ErrorActionPreference = 'Stop'
$src = $Src
if (-not $src) { $src = Join-Path $RepoRoot 'X32-Behringer' }
if (-not (Test-Path (Join-Path $src 'X32.c'))) {
    Write-Host ''
    Write-Host ("X32.c was not found in '{0}'." -f $src)
    Write-Host 'The C sources of the reference emulator are not part of this repository'
    Write-Host '(only the files generated from them are). Clone the upstream project:'
    Write-Host ''
    Write-Host '    git clone https://github.com/pmaillot/X32-Behringer'
    Write-Host '    powershell -File tools\gen_tables.ps1 -Src C:\src\X32-Behringer'
    Write-Host ''
    exit 1
}
$docs = Join-Path $RepoRoot 'docs'
$go   = Join-Path $RepoRoot 'emulator\internal\x32'
foreach ($p in @($docs, $go)) { if (-not (Test-Path $p)) { New-Item -ItemType Directory -Force -Path $p | Out-Null } }

$tableFiles = @('X32Channel.h','X32CfgMain.h','X32PrefStat.h','X32Auxin.h','X32Fxrtn.h','X32Bus.h',
                'X32Mtx.h','X32Dca.h','X32Fx.h','X32Output.h','X32Headamp.h','X32Show.h','X32Misc.h','X32Libs.h')

function Write-Text([string]$path, [string]$text) {
    [System.IO.File]::WriteAllText($path, $text, (New-Object System.Text.UTF8Encoding($false)))
}
function GoStr([string]$s) { return '"' + ($s.Replace('\','\\').Replace('"','\"')) + '"' }
function GoFlag([string]$s) {
    $fl = $s.Replace(' ', '')
    switch ($fl) {
        'F_FND'            { return 'FFND' }
        'F_XET'            { return 'FXET' }
        'F_GET'            { return 'FGET' }
        'F_SET'            { return 'FSET' }
        'F_NPR'            { return 'FNPR' }
        'F_GET|F_SET'      { return 'FGET|FSET' }
        default            { return 'F0 /*' + $fl + '*/' }
    }
}

# ---------------------------------------------------------------- enums
$sb = New-Object System.Text.StringBuilder
foreach ($f in $tableFiles) { [void]$sb.AppendLine((Get-Content (Join-Path $src $f) -Raw)) }
[void]$sb.AppendLine((Get-Content (Join-Path $src 'X32.c') -Raw))
$enumText = $sb.ToString()

$enumNames = New-Object System.Collections.ArrayList
$enumMap   = @{}
foreach ($m in [regex]::Matches($enumText, 'char\s*\*\s*(?<n>\w+)\s*\[\s*\]\s*=\s*\{(?<b>.*?)\}\s*;', 'Singleline')) {
    $name = $m.Groups['n'].Value
    if ($enumMap.ContainsKey($name)) { continue }
    if ($m.Groups['b'].Value -match '=') { continue }
    $vals = New-Object System.Collections.ArrayList
    foreach ($s in [regex]::Matches($m.Groups['b'].Value, '"((?:[^"\\]|\\.)*)"')) {
        [void]$vals.Add($s.Groups[1].Value.Replace('\"','"').Replace('\\','\'))
    }
    if ($vals.Count -eq 0) { continue }
    $enumMap[$name] = $vals
    [void]$enumNames.Add($name)
}

# ---------------------------------------------------------------- enum types { ... }
$typeNames = New-Object System.Collections.ArrayList
$tb = [regex]::Match((Get-Content (Join-Path $src 'X32.c') -Raw), 'enum\s+types\s*\{(?<b>.*?)\};', 'Singleline').Groups['b'].Value
foreach ($part in ($tb -split ',')) {
    $t = ($part -replace '//.*','').Trim()
    if ($t -match '^[A-Z]\w*$') { [void]$typeNames.Add($t) }
}

# ---------------------------------------------------------------- command tables
$tables    = New-Object System.Collections.ArrayList   # @{Name;File;Entries=@(...)}
$sets      = New-Object System.Collections.ArrayList   # @{Name;Tables=@(names)}
$warnings  = New-Object System.Collections.ArrayList
$byType    = @{}
$byFile    = @{}

foreach ($f in $tableFiles) {
    $txt = Get-Content (Join-Path $src $f) -Raw
    foreach ($m in [regex]::Matches($txt, 'X32command\s+(?<n>\w+)\s*\[\s*\]\s*=\s*\{(?<b>.*?)\}\s*;', 'Singleline')) {
        $name = $m.Groups['n'].Value
        $entries = New-Object System.Collections.ArrayList
        foreach ($e in [regex]::Matches($m.Groups['b'].Value,
                  '\{\s*"(?<a>[^"]*)"\s*,\s*\{\s*(?<t>\w+)\s*\}\s*,\s*(?<fl>[^,]+?)\s*,\s*\{\s*(?<c>-?\d*)\s*\}\s*,\s*(?<nd>[^}]+?)\s*\}', 'Singleline')) {
            $node = $e.Groups['nd'].Value.Trim()
            if ($node -eq 'NULL') { $node = $null }
            elseif ($node -notmatch '^\w+$') { [void]$warnings.Add("$f/$name : odd node field '$node'"); $node = $null }
            elseif (-not $enumMap.ContainsKey($node)) { [void]$warnings.Add("$f/$name : unknown enum '$node'"); $node = $null }
            $cnt = 0; if ($e.Groups['c'].Value -ne '') { $cnt = [int]$e.Groups['c'].Value }
            [void]$entries.Add([pscustomobject]@{ Addr = $e.Groups['a'].Value; Type = $e.Groups['t'].Value; Flags = $e.Groups['fl'].Value.Trim(); Count = $cnt; Node = $node })
            $t = $e.Groups['t'].Value
            if (-not $byType.ContainsKey($t)) { $byType[$t] = 0 }
            $byType[$t] = $byType[$t] + 1
        }
        if ($entries.Count -eq 0) { [void]$warnings.Add("$f/$name : no entries parsed"); continue }
        [void]$tables.Add([pscustomobject]@{ Name = $name; File = $f; Entries = $entries })
        if (-not $byFile.ContainsKey($f)) { $byFile[$f] = 0 }
        $byFile[$f] = $byFile[$f] + $entries.Count
    }
    foreach ($m in [regex]::Matches($txt, 'X32command\s*\*\s*(?<n>\w+)\s*\[\s*\d*\s*\]\s*=\s*\{(?<b>.*?)\}\s*;', 'Singleline')) {
        $list = New-Object System.Collections.ArrayList
        foreach ($part in ($m.Groups['b'].Value -split ',')) {
            $id = ($part -replace '//.*','').Trim()
            if ($id -match '^\w+$') { [void]$list.Add($id) }
        }
        if ($list.Count -gt 0) { [void]$sets.Add([pscustomobject]@{ Name = $m.Groups['n'].Value; Tables = $list }) }
    }
}

$totalEntries = 0
foreach ($t in $tables) { $totalEntries += $t.Entries.Count }

# === EMIT ===

$hdr = @"
// Code generated by tools/gen_tables.ps1; DO NOT EDIT.
// Source: X32-Behringer by Patrick-Gilles Maillot (GPLv3) - X32.c v0.88 and the
// command table headers. https://github.com/pmaillot/X32-Behringer

package x32

"@

# ---------------------------------------------------------------- types_gen.go
$o = New-Object System.Text.StringBuilder
[void]$o.AppendLine($hdr)
[void]$o.AppendLine('// Type is the OSC value/command type of a table entry (parity with "enum types" in X32.c).')
[void]$o.AppendLine('type Type uint8')
[void]$o.AppendLine('')
[void]$o.AppendLine('const (')
for ($i = 0; $i -lt $typeNames.Count; $i++) {
    [void]$o.AppendLine(("`tT{0} Type = {1}" -f $typeNames[$i], $i))
}
[void]$o.AppendLine(')')
[void]$o.AppendLine('')
[void]$o.AppendLine('var typeNames = map[Type]string{')
for ($i = 0; $i -lt $typeNames.Count; $i++) {
    [void]$o.AppendLine(("`tT{0}: {1}," -f $typeNames[$i], (GoStr $typeNames[$i])))
}
[void]$o.AppendLine('}')
[void]$o.AppendLine('')
[void]$o.AppendLine('// String returns the C name of a type (for dumping / debugging).')
[void]$o.AppendLine('func (t Type) String() string {')
[void]$o.AppendLine("`tif s, ok := typeNames[t]; ok {")
[void]$o.AppendLine("`t`treturn s")
[void]$o.AppendLine("`t}")
[void]$o.AppendLine("`treturn `"UNKNOWN`"")
[void]$o.AppendLine('}')
Write-Text (Join-Path $go 'types_gen.go') $o.ToString()

# ---------------------------------------------------------------- enums_gen.go
$o = New-Object System.Text.StringBuilder
[void]$o.AppendLine($hdr)
[void]$o.AppendLine('// Enum is a named list of enum strings. Values are kept verbatim (most carry a')
[void]$o.AppendLine('// leading space) because they are concatenated as-is into /node replies.')
[void]$o.AppendLine('type Enum struct {')
[void]$o.AppendLine("`tName   string")
[void]$o.AppendLine("`tValues []string")
[void]$o.AppendLine('}')
[void]$o.AppendLine('')
[void]$o.AppendLine('var Enums = []Enum{')
for ($i = 0; $i -lt $enumNames.Count; $i++) {
    $n = $enumNames[$i]
    $vals = ($enumMap[$n] | ForEach-Object { GoStr $_ }) -join ', '
    [void]$o.AppendLine(("`t{{Name: {0}, Values: []string{{{1}}}}}," -f (GoStr $n), $vals))
}
[void]$o.AppendLine('}')
[void]$o.AppendLine('')
[void]$o.AppendLine('// enums indexes the array above by its C name.')
[void]$o.AppendLine('var enums = map[string]*Enum{')
for ($i = 0; $i -lt $enumNames.Count; $i++) {
    [void]$o.AppendLine(("`t{0}: &Enums[{1}]," -f (GoStr $enumNames[$i]), $i))
}
[void]$o.AppendLine('}')
Write-Text (Join-Path $go 'enums_gen.go') $o.ToString()

# ---------------------------------------------------------------- tables_gen.go
$o = New-Object System.Text.StringBuilder
[void]$o.AppendLine($hdr)
[void]$o.AppendLine('// Command tables: mirror of the arrays found in the upstream X32-Behringer headers.')
[void]$o.AppendLine('// Literal order: {Addr, Type, Flags, Count, Enum}.')
foreach ($t in $tables) {
    [void]$o.AppendLine('')
    [void]$o.AppendLine(("// Table {0} (from {1}, {2} entries)" -f $t.Name, $t.File, $t.Entries.Count))
    [void]$o.AppendLine(("var {0} = Table{{" -f $t.Name))
    foreach ($e in $t.Entries) {
        $en = 'nil'
        if ($e.Node) { $en = 'enums[' + (GoStr $e.Node) + ']' }
        [void]$o.AppendLine(("`t{{{0}, T{1}, {2}, {3}, {4}, 0, 0, nil}}," -f (GoStr $e.Addr), $e.Type, (GoFlag $e.Flags), $e.Count, $en))
    }
    [void]$o.AppendLine('}')
}
Write-Text (Join-Path $go 'tables_gen.go') $o.ToString()

# ---------------------------------------------------------------- sets_gen.go
$o = New-Object System.Text.StringBuilder
[void]$o.AppendLine($hdr)
[void]$o.AppendLine('// Table sets: mirror of the "*set" arrays of the upstream X32-Behringer headers.')
foreach ($s in $sets) {
    [void]$o.AppendLine('')
    [void]$o.AppendLine(("var {0} = []Table{{{1}}}" -f $s.Name, (($s.Tables | ForEach-Object { "`n`t$_," }) -join '')))
}
Write-Text (Join-Path $go 'sets_gen.go') $o.ToString()

# === GOFMT ===
# The generated Go files are committed and CI runs `gofmt -l`, so they must be
# gofmt clean: gofmt aligns the type/enum/flag blocks that are emitted raw above.
$genFiles = @('types_gen.go', 'enums_gen.go', 'tables_gen.go', 'sets_gen.go') |
    ForEach-Object { Join-Path $go $_ }
if (Get-Command gofmt -ErrorAction SilentlyContinue) {
    & gofmt -w @genFiles
} else {
    Write-Warning ('gofmt not found in PATH: run "gofmt -w {0}" before committing' -f ($genFiles -join ' '))
}

# === JSON ===

# ---------------------------------------------------------------- docs/x32_commands.json
$j = New-Object System.Text.StringBuilder
[void]$j.AppendLine('{')
[void]$j.AppendLine(('  "source": {0},' -f (GoStr 'X32-Behringer (https://github.com/pmaillot/X32-Behringer) / X32.c v0.88 + command table headers')))
[void]$j.AppendLine(('  "generated": {0},' -f (GoStr (Get-Date -Format 'yyyy-MM-dd HH:mm:ss'))))
[void]$j.AppendLine(('  "totals": {{"tables": {0}, "entries": {1}, "enums": {2}}},' -f $tables.Count, $totalEntries, $enumNames.Count))
[void]$j.AppendLine('  "types": [')
$tnames = @($byType.Keys | Sort-Object)
for ($i = 0; $i -lt $tnames.Count; $i++) {
    $comma = if ($i -lt $tnames.Count - 1) { ',' } else { '' }
    [void]$j.AppendLine(('    {{"type": {0}, "entries": {1}}}{2}' -f (GoStr $tnames[$i]), $byType[$tnames[$i]], $comma))
}
[void]$j.AppendLine('  ],')
[void]$j.AppendLine('  "tables": [')
for ($ti = 0; $ti -lt $tables.Count; $ti++) {
    $t = $tables[$ti]
    [void]$j.AppendLine(('    {{"name": {0}, "file": {1}, "entries": [' -f (GoStr $t.Name), (GoStr $t.File)))
    for ($ei = 0; $ei -lt $t.Entries.Count; $ei++) {
        $e = $t.Entries[$ei]
        $node = if ($e.Node) { GoStr $e.Node } else { 'null' }
        $comma = if ($ei -lt $t.Entries.Count - 1) { ',' } else { '' }
        [void]$j.AppendLine(('      {{"a": {0}, "t": {1}, "f": {2}, "n": {3}, "e": {4}}}{5}' -f (GoStr $e.Addr), (GoStr $e.Type), (GoStr $e.Flags), $e.Count, $node, $comma))
    }
    $comma = if ($ti -lt $tables.Count - 1) { ',' } else { '' }
    [void]$j.AppendLine(('    ]}}{0}' -f $comma))
}
[void]$j.AppendLine('  ]')
[void]$j.AppendLine('}')
Write-Text (Join-Path $docs 'x32_commands.json') $j.ToString()

# ---------------------------------------------------------------- docs/x32_enums.json
$e2 = New-Object System.Text.StringBuilder
[void]$e2.AppendLine('{')
for ($i = 0; $i -lt $enumNames.Count; $i++) {
    $n = $enumNames[$i]
    $vals = ($enumMap[$n] | ForEach-Object { GoStr $_ }) -join ', '
    $comma = if ($i -lt $enumNames.Count - 1) { ',' } else { '' }
    [void]$e2.AppendLine(('  {0}: [{1}]{2}' -f (GoStr $n), $vals, $comma))
}
[void]$e2.AppendLine('}')
Write-Text (Join-Path $docs 'x32_enums.json') $e2.ToString()

# === REPORT ===

Write-Output '== X32 table extraction report =='
Write-Output ("source          : {0}" -f $src)
Write-Output ("tables          : {0}" -f $tables.Count)
Write-Output ("entries (total) : {0}" -f $totalEntries)
Write-Output ("enum arrays     : {0}" -f $enumNames.Count)
Write-Output ("types           : {0}" -f $typeNames.Count)
Write-Output ("sets            : {0}" -f $sets.Count)
Write-Output '-- entries by type --'
foreach ($n in $tnames) { Write-Output ("   {0,-10} {1}" -f $n, $byType[$n]) }
Write-Output '-- entries by file --'
foreach ($n in ($byFile.Keys | Sort-Object)) { Write-Output ("   {0,-16} {1}" -f $n, $byFile[$n]) }
Write-Output '-- warnings --'
if ($warnings.Count -eq 0) { Write-Output '   none' } else { foreach ($w in $warnings) { Write-Output ('   ' + $w) } }
Write-Output '-- set arrays --'
foreach ($s in $sets) { Write-Output ("   {0} = {1}" -f $s.Name, ($s.Tables -join ',')) }

