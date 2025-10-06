#!/bin/bash

# Simple starter for backend and frontend with basic cleanup and browser open

set -e

if [ ! -d "tp_middleware_core" ] || [ ! -d "tp-middleware-front-2025" ]; then
    echo "Run this script from the Timetable-alerter directory."
    exit 1
fi

cleanup() {
    [ -n "$FRONTEND_PID" ] && kill $FRONTEND_PID 2>/dev/null || true
    [ -n "$BACKEND_PID" ] && kill $BACKEND_PID 2>/dev/null || true
    exit 0
}
trap cleanup INT TERM

cd tp_middleware_core
go run cmd/main.go &
BACKEND_PID=$!

cd ../tp-middleware-front-2025
[ -d node_modules ] || npm install
npm run dev &
FRONTEND_PID=$!

# Try to open browser in Windows/WSL or Linux
URL="http://localhost:5173"
if command -v wslview >/dev/null 2>&1; then
    wslview "$URL" || true
elif command -v powershell.exe >/dev/null 2>&1; then
    powershell.exe Start-Process "$URL" >/dev/null 2>&1 || true
elif command -v explorer.exe >/dev/null 2>&1; then
    explorer.exe "$URL" >/dev/null 2>&1 || true
elif command -v xdg-open >/dev/null 2>&1; then
    xdg-open "$URL" >/dev/null 2>&1 || true
else
    echo "Open $URL in your browser."
fi

echo "Backend:  http://localhost:8080"
echo "Frontend: $URL"
echo "Press Ctrl+C to stop"

wait