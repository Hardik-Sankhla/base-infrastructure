# Universal Bootstrap Framework - Entrypoint
$ErrorActionPreference = "Stop"

$DIR = $PSScriptRoot

Write-Host "🚀 Initializing Universal Bootstrap Framework..." -ForegroundColor Cyan

# Invoke the core engine (PowerShell version)
$discoverPath = Join-Path $DIR "core\discovery\discover.ps1"
if (Test-Path $discoverPath) {
    & $discoverPath
} else {
    Write-Error "❌ Core discovery engine not found at $discoverPath"
    exit 1
}

Write-Host "✅ Bootstrap complete." -ForegroundColor Green
