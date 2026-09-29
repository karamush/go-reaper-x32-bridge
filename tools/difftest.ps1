# difftest.ps1
#
# Differential test between the reference C emulator (X32-emulator.exe of the
# X32-Behringer project, https://github.com/pmaillot/X32-Behringer) and this Go
# port (emulator/cmd/x32emu). Both are started alone on 127.0.0.1:10023 and
# driven by the same scripted session (emulator/cmd/x32probe), then the two
# captured logs are compared line by line.
#
# The reference binary is NOT part of this repository: build it from the upstream
# C sources, or drop a prebuilt X32-emulator.exe in the repository root, or pass
# -Oracle <path>. When it is missing the script explains how to get it and exits
# successfully, so it stays safe to run anywhere (CI included).
#
# Usage:  powershell -File tools\difftest.ps1 [-Oracle <exe>] [-Root <repo root>]

[CmdletBinding()]
param([string]$Root = '', [string]$Oracle = '')

if (-not $Root) {
    $here = $PSScriptRoot
    if (-not $here) { $here = Split-Path -Parent $MyInvocation.MyCommand.Path }
    $Root = (Resolve-Path (Join-Path $here '..')).Path
}

$ErrorActionPreference = 'Stop'

# Locate the reference emulator binary (see the header: it is not shipped here).
if (-not $Oracle) {
    foreach ($candidate in @((Join-Path $Root 'X32-emulator.exe'),
                             (Join-Path $Root 'X32-Behringer\X32-emulator.exe'))) {
        if (Test-Path $candidate) { $Oracle = $candidate; break }
    }
}
if (-not $Oracle -or -not (Test-Path $Oracle)) {
    Write-Output 'difftest: skipped - the reference emulator binary was not found.'
    Write-Output ''
    Write-Output 'It is not part of this repository (it is the GPLv3 reference implementation'
    Write-Output 'of the upstream X32-Behringer project). Build it from those C sources:'
    Write-Output ''
    Write-Output '    git clone https://github.com/pmaillot/X32-Behringer'
    Write-Output '    cd X32-Behringer && make            # MinGW; MSVC: see its README'
    Write-Output ''
    Write-Output 'or use a prebuilt X32-emulator.exe, then:'
    Write-Output ''
    Write-Output '    powershell -File tools\difftest.ps1 -Oracle <path to X32-emulator.exe>'
    Write-Output ''
    exit 0
}

$build = Join-Path $Root 'build'
$crun = Join-Path $build 'crun'
$gorun = Join-Path $build 'gorun'
foreach ($d in @($build, $crun, $gorun)) { if (-not (Test-Path $d)) { New-Item -ItemType Directory -Force -Path $d | Out-Null } }

# build the Go binaries
Push-Location (Join-Path $Root 'emulator')
& go build -o (Join-Path $build 'x32emu.exe') ./cmd/x32emu
& go build -o (Join-Path $build 'x32probe.exe') ./cmd/x32probe
Pop-Location

function Run-Probe([string]$exe, [string]$workdir, [string]$out, [string]$emuArgs) {
    Remove-Item (Join-Path $workdir '.X32res.rc') -ErrorAction SilentlyContinue
    $p = Start-Process -FilePath $exe -ArgumentList $emuArgs -WorkingDirectory $workdir -PassThru `
        -RedirectStandardOutput (Join-Path $workdir 'emulator.log')
    Start-Sleep -Milliseconds 900
    & (Join-Path $build 'x32probe.exe') -addr 127.0.0.1:10023 > $out 2>&1
    Stop-Process -Id $p.Id -Force -ErrorAction SilentlyContinue
    Start-Sleep -Milliseconds 500
}

Run-Probe $Oracle $crun (Join-Path $build 'c.txt') '-i 127.0.0.1 -v 0'
# The port answers to the name "REAPER" by default (the -name flag), the C
# emulator hard codes "X32 Emulator": use the same name here so that the two
# logs stay byte comparable.
Run-Probe (Join-Path $build 'x32emu.exe') $gorun (Join-Path $build 'go.txt') '-i 127.0.0.1 -v 0 -name "X32 Emulator"'

$c = Get-Content (Join-Path $build 'c.txt')
$g = Get-Content (Join-Path $build 'go.txt')
$d = Compare-Object $c $g

# Meter frames are asynchronous (they are delivered on a 10 s window, not as an
# answer to a request), so an extra or missing frame is a timing artifact and not
# a protocol difference: they are counted apart.
$meter = @($d | Where-Object { $_.InputObject -match '/meters/' })
$other = @($d | Where-Object { $_.InputObject -notmatch '/meters/' })

Write-Output "C emulator : $($c.Count) lines"
Write-Output "Go emulator: $($g.Count) lines"
Write-Output "Differences: $($d.Count) lines ($($other.Count) protocol, $($meter.Count) asynchronous meter frames)"
$other | ForEach-Object { Write-Output ("{0} {1}" -f $_.SideIndicator, $_.InputObject) }
if ($meter.Count -gt 0) {
    Write-Output "--- meter frames (timing) ---"
    $meter | ForEach-Object { Write-Output ("{0} {1}" -f $_.SideIndicator, $_.InputObject) }
}
