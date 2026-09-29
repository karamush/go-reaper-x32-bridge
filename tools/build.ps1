# build.ps1
#
# Builds every command of the two Go modules into bin/ (or -Out <dir>):
#
#   x32reaper   the console and the REAPER bridge in one process (recommended)
#   x32emu      the console alone
#   x32bridge   the bridge alone (real console, or an emulator started separately)
#   x32listen   OSC dump tool (either side; -meters decodes meter frames)
#   x32probe    scripted OSC session, the driver of tools/difftest.ps1
#
# The two modules are wired together by the go.work file at the repository root,
# so this script can be run from anywhere.
#
# Usage:
#   powershell -File tools\build.ps1
#   powershell -File tools\build.ps1 -Out C:\temp\x32 -Release
#   powershell -File tools\build.ps1 -TargetOS linux -Arch amd64 -Release

[CmdletBinding()]
param(
    [string]$Root = '',
    [string]$Out = '',
    [AllowEmptyString()]
    [ValidateSet('windows', 'linux', 'darwin', 'freebsd')]
    [string]$TargetOS = '',
    [string]$Arch = '',
    [switch]$Release
)

$ErrorActionPreference = 'Stop'

if (-not $Root) {
    $here = $PSScriptRoot
    if (-not $here) { $here = Split-Path -Parent $MyInvocation.MyCommand.Path }
    if (-not $here) { $here = (Get-Location).Path }
    $Root = (Resolve-Path (Join-Path $here '..')).Path
}
if (-not $Out) { $Out = Join-Path $Root 'bin' }
if (-not (Test-Path $Out)) { New-Item -ItemType Directory -Force -Path $Out | Out-Null }
# Absolute, always: the build below changes the current directory, so a relative
# -Out would otherwise be resolved against each module directory.
$Out = (Resolve-Path $Out).Path

$commands = @(
    [pscustomobject]@{ Module = 'emulator'; Package = 'cmd/x32emu';    Name = 'x32emu' },
    [pscustomobject]@{ Module = 'emulator'; Package = 'cmd/x32probe';  Name = 'x32probe' },
    [pscustomobject]@{ Module = 'bridge';   Package = 'cmd/x32reaper'; Name = 'x32reaper' },
    [pscustomobject]@{ Module = 'bridge';   Package = 'cmd/x32bridge'; Name = 'x32bridge' },
    [pscustomobject]@{ Module = 'bridge';   Package = 'cmd/x32listen'; Name = 'x32listen' }
)

$saved = @{}
foreach ($v in @('GOOS', 'GOARCH')) { $saved[$v] = [System.Environment]::GetEnvironmentVariable($v) }

$ext = '.exe'
try {
    if ($TargetOS) { $env:GOOS = $TargetOS }
    if ($Arch) { $env:GOARCH = $Arch }
    if ($env:GOOS) { if ($env:GOOS -ne 'windows') { $ext = '' } }
    elseif ($env:OS -ne 'Windows_NT') { $ext = '' }

    $flags = @('-trimpath')
    if ($Release) { $flags += @('-ldflags', '-s -w') }

    Write-Output ("building {0} binaries into {1}" -f $(if ($Release) { 'release' } else { 'debug' }), $Out)
    foreach ($c in $commands) {
        # NB: keep this name distinct from $Out - PowerShell variable names are
        # case insensitive, so $out would overwrite the output directory.
        $target = Join-Path $Out ($c.Name + $ext)
        Push-Location (Join-Path $Root $c.Module)
        try {
            & go build @flags -o $target ('./' + $c.Package)
            if ($LASTEXITCODE -ne 0) { throw ("go build failed for " + $c.Name) }
        } finally {
            Pop-Location
        }
        if (-not (Test-Path $target)) { throw ("go build produced no file for " + $c.Name) }
        $kb = [math]::Round((Get-Item $target).Length / 1KB, 0)
        Write-Output ("  {0,-11} {1,7} kB" -f $c.Name, $kb)
    }
} finally {
    foreach ($v in @('GOOS', 'GOARCH')) {
        if ($saved[$v]) { Set-Item -Path ('Env:' + $v) -Value $saved[$v] }
        else { Remove-Item -Path ('Env:' + $v) -ErrorAction SilentlyContinue }
    }
}
