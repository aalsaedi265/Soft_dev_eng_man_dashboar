@echo off
echo ========================================
echo Running Database Migrations
echo ========================================

REM Load environment variables from .env file
for /f "tokens=1,2 delims==" %%a in (.env) do (
    if "%%a"=="DB_HOST" set DB_HOST=%%b
    if "%%a"=="DB_PORT" set DB_PORT=%%b
    if "%%a"=="DB_USER" set DB_USER=%%b
    if "%%a"=="DB_PASSWORD" set DB_PASSWORD=%%b
    if "%%a"=="DB_NAME" set DB_NAME=%%b
    if "%%a"=="DB_SSLMODE" set DB_SSLMODE=%%b
)

REM Install goose if not already installed
echo Checking for goose...
where goose >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo Installing goose migration tool...
    go install github.com/pressly/goose/v3/cmd/goose@latest
)

echo.
echo Running migrations...
goose -dir migrations postgres "host=%DB_HOST% port=%DB_PORT% user=%DB_USER% password=%DB_PASSWORD% dbname=%DB_NAME% sslmode=%DB_SSLMODE%" up

echo.
echo Migrations completed!
pause
