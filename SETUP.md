# Backend Setup Guide

## Prerequisites

✅ **Already Installed:**
- PostgreSQL 18.0
- Go 1.21+
- Node.js 18+

## Quick Start (Windows)

### Option A: Use Batch Scripts (Easiest)

1. **Set up database:**
   ```cmd
   cd backend
   setup-db.bat
   ```

2. **Run migrations:**
   ```cmd
   run-migrations.bat
   ```

3. **Start server:**
   ```cmd
   start-server.bat
   ```

### Option B: Manual Setup

## 1. Create Database

**Method 1: Using psql**
```bash
psql -U postgres
CREATE DATABASE teamdashboard;
\q
```

**Method 2: Using pgAdmin**
- Open pgAdmin
- Right-click on Databases → Create → Database
- Name: `teamdashboard`
- Click Save

## 2. Install Goose (Migration Tool)

### Why Goose?

**Problem:** Managing database schema changes across a team is messy:
- Running SQL files manually is error-prone
- Hard to track which changes have been applied
- No easy way to rollback mistakes
- Team members get out of sync

**Solution:** Goose tracks and applies database migrations automatically.

**What it does:**
1. **Versioned migrations** - Each SQL file is numbered (001, 002, 003...)
2. **Tracks state** - Knows which migrations you've already run
3. **Idempotent** - Safe to run multiple times, only applies new changes
4. **Team sync** - Everyone runs the same migrations in the same order
5. **Rollback support** - Can undo changes if needed

**Alternative approaches we rejected:**
- ❌ Manual SQL scripts - No tracking, easy to miss files
- ❌ ORMs (like GORM) - Over-engineered, hides SQL, harder to debug
- ✅ Goose - Simple, explicit SQL, automated tracking

**Install:**
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Verify installation:
```bash
goose -version
```

## 3. Configure Database Connection

The `.env` file is already created in `backend/` with default settings:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=teamdashboard
DB_SSLMODE=disable

JWT_SECRET=my-super-secret-jwt-key-change-in-production
PORT=8080
```

**Update the password** if your PostgreSQL uses a different password.

## 4. Run Database Migrations

### What happens when you run migrations:

Goose will:
1. Create a tracking table `goose_db_version` in your database
2. Check which migrations have already been applied
3. Run any new migrations in sequential order
4. Record each successful migration

**Run migrations:**
```bash
cd backend
goose -dir migrations postgres "user=postgres password=postgres dbname=teamdashboard sslmode=disable" up
```

You should see:
```
OK    001_create_employees.sql
OK    002_create_projects.sql
OK    003_create_tasks.sql
OK    004_create_time_logs.sql
OK    005_create_users.sql
```

**What just happened:**
- Created 5 tables: `employees`, `projects`, `tasks`, `time_logs`, `users`
- Created 1 join table: `project_assignments` (for many-to-many relationships)
- Added foreign keys and indexes for performance
- Created tracking table: `goose_db_version`

**Common migration commands:**
```bash
# Apply all pending migrations
goose up

# Rollback the last migration
goose down

# Check migration status
goose status

# Reset database (careful!)
goose reset
```

## 5. Dependencies (Already Installed ✅)

Dependencies were already downloaded. If you need to reinstall:
```bash
cd backend
go mod download
```

## 6. Start Backend Server

```bash
cd backend
go run cmd/server/main.go
```

Server will start on http://localhost:8080

## 7. Start Frontend (in new terminal)

```bash
cd frontend/vite-project
npm run dev
```

Frontend will start on http://localhost:5173

## 8. Test the API

### Register a user:
```bash
curl -X POST http://localhost:8080/api/auth/register -H "Content-Type: application/json" -d "{\"email\":\"test@example.com\",\"password\":\"password123\"}"
```

### Login:
```bash
curl -X POST http://localhost:8080/api/auth/login -H "Content-Type: application/json" -d "{\"email\":\"test@example.com\",\"password\":\"password123\"}"
```

Copy the token from the response and use it in subsequent requests:

### Get employees (with auth):
```bash
curl http://localhost:8080/api/employees -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Done!

Your backend is now running with:
- ✅ Tasks CRUD
- ✅ Time Tracking/Logs CRUD
- ✅ Authentication (JWT)
- ✅ Protected routes
- ✅ Employee, Project, and Task management

Next step: Connect the frontend to use these APIs instead of mock data.
