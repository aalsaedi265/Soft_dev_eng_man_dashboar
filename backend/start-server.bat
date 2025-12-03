@echo off
echo ========================================
echo Starting Backend Server
echo ========================================

echo.
echo Server will start on http://localhost:8080
echo.

go run cmd/server/main.go
