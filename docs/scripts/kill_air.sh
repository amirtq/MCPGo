#!/bin/bash

# Define the path to the MCPGo binary relative to the project root
MCPGO_BIN_PATH="bin/mcpgo"

# Find the PID of the mcpgo process from the full command path
MCPGO_PID=$(ps -axo pid=,command= | grep "$MCPGO_BIN_PATH" | grep -v grep | awk '{print $1}')

if [ -z "$MCPGO_PID" ]; then
  echo "No mcpgo process found."
  exit 0
fi

# Find the parent PID (which should be the specific air process)
AIR_PID=$(ps -o ppid= -p "$MCPGO_PID" | tr -d ' ')

# Combine unique PIDs
PIDS=$(printf "%s\n%s\n" "$MCPGO_PID" "$AIR_PID" | sort -u | awk 'NF')

if [ -z "$PIDS" ]; then
  echo "No matching processes found to kill."
  exit 0
fi

echo "Will kill the following processes:"
for PID in $PIDS; do
  ps -p "$PID" -o pid=,command=
done

# Try graceful kill first
kill $PIDS 2>/dev/null
sleep 1

# Force kill if still running
for PID in $PIDS; do
  if kill -0 "$PID" 2>/dev/null; then
    echo "Force killing PID: $PID"
    kill -9 "$PID" 2>/dev/null
  fi
done

echo "✅ Done."
