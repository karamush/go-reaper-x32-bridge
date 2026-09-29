# gen_appendix.ps1
#
# Reads docs/x32_commands.json (produced by gen_tables.ps1) and writes
# docs/APPENDIX_COMMANDS.md: the complete command catalogue, grouped by
# subsystem and by node group, ready to be used as a reference.
#
# Usage:  powershell -File tools\gen_appendix.ps1 [-RepoRoot <path>]

[CmdletBinding()]
param([string]$RepoRoot = '')

if (-not $RepoRoot) {
    $here = $PSScriptRoot
    if (-not $here) { $here = Split-Path -Parent $MyInvocation.MyCommand.Path }
    $RepoRoot = (Resolve-Path (Join-Path $here '..')).Path
}

$ErrorActionPreference = 'Stop'
$json = Get-Content (Join-Path $RepoRoot 'docs\x32_commands.json') -Raw | ConvertFrom-Json

function Write-Text([string]$path, [string]$text) {
    [System.IO.File]::WriteAllText($path, $text, (New-Object System.Text.UTF8Encoding($false)))
}

$sb = New-Object System.Text.StringBuilder
[void]$sb.AppendLine('# Appendix: X32 command catalogue')
[void]$sb.AppendLine('')
[void]$sb.AppendLine('Generated from `docs/x32_commands.json`, itself extracted by `tools/gen_tables.ps1`')
[void]$sb.AppendLine('from the command table headers of the upstream X32-Behringer project')
[void]$sb.AppendLine('(<https://github.com/pmaillot/X32-Behringer>, Patrick-Gilles Maillot, GPLv3).')
[void]$sb.AppendLine('')
[void]$sb.AppendLine(("- total tables: {0}" -f $json.totals.tables))
[void]$sb.AppendLine(("- total entries: {0} (leaf commands + group headers)" -f $json.totals.entries))
[void]$sb.AppendLine(("- enum string arrays: {0}" -f $json.totals.enums))
[void]$sb.AppendLine('')
[void]$sb.AppendLine('## Entries per type')
[void]$sb.AppendLine('')
[void]$sb.AppendLine('| type | entries |')
[void]$sb.AppendLine('| --- | --- |')
foreach ($t in $json.types) { [void]$sb.AppendLine(("| {0} | {1} |" -f $t.type, $t.entries)) }
[void]$sb.AppendLine('')
[void]$sb.AppendLine('## Entries per table')
[void]$sb.AppendLine('')
[void]$sb.AppendLine('| table | source file | entries |')
[void]$sb.AppendLine('| --- | --- | --- |')
foreach ($t in $json.tables) {
    [void]$sb.AppendLine(("| {0} | {1} | {2} |" -f $t.name, $t.file, $t.entries.Count))
}
[void]$sb.AppendLine('')
[void]$sb.AppendLine('## Command tables')
[void]$sb.AppendLine('')
[void]$sb.AppendLine('Group headers are marked with `F_FND`; `n=` is the item count stored in the')
[void]$sb.AppendLine('table (used by the `/node` renderer when it is not zero). Leaves carry their')
[void]$sb.AppendLine('value type and their enum array name (if any).')
[void]$sb.AppendLine('')

foreach ($t in $json.tables) {
    [void]$sb.AppendLine(("### {0} ({1}, {2} entries)" -f $t.name, $t.file, $t.entries.Count))
    [void]$sb.AppendLine('')
    [void]$sb.AppendLine('```')
    foreach ($e in $t.entries) {
        if ($e.f -eq 'F_FND') {
            [void]$sb.AppendLine(("{0}  <{1}> n={2}" -f $e.a, $e.t, $e.n))
        } else {
            $en = ''
            if ($e.e) { $en = ' enum=' + $e.e }
            [void]$sb.AppendLine(("    {0}  {1} {2}{3}" -f $e.a, $e.t, $e.f, $en))
        }
    }
    [void]$sb.AppendLine('```')
    [void]$sb.AppendLine('')
}

Write-Text (Join-Path $RepoRoot 'docs\APPENDIX_COMMANDS.md') $sb.ToString()
Write-Output ("docs/APPENDIX_COMMANDS.md written ({0} tables)" -f $json.tables.Count)
