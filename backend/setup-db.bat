@echo off
echo ========================================
echo PostgreSQL Database Setup
echo ========================================

REM Load environment variables from .env file
for /f "tokens=1,2 delims==" %%a in (.env) do (
    if "%%a"=="DB_PASSWORD" set DB_PASSWORD=%%b
)

echo.
echo Creating database 'teamdashboard'...
set PGPASSWORD=%DB_PASSWORD%
psql -U postgres -c "CREATE DATABASE teamdashboard;"

echo.
echo Database created successfully!
echo.
echo Next step: Run migrations with run-migrations.bat
echo.
pause
