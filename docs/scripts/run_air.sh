#!/bin/sh
# This script starts the air live-reloader for the MCPGo project.

# Navigate to the project root directory relative to the script's location
cd "$(dirname "$0")/../.."

# Execute the air command
~/go/bin/air
