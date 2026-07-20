#!/usr/bin/env bash
# Universal Bootstrap Framework - Entrypoint
set -euo pipefail

# Find the directory of the script
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

echo "🚀 Initializing Universal Bootstrap Framework..."

# Invoke the core engine (Shell version)
if [ -f "$DIR/core/discovery/discover.sh" ]; then
    bash "$DIR/core/discovery/discover.sh"
else
    echo "❌ Error: Core discovery engine not found at $DIR/core/discovery/discover.sh"
    exit 1
fi

echo "✅ Bootstrap complete."
