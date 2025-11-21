# Quick Setup Guide

## 1. Install PostgreSQL (if not already installed)

Download and install PostgreSQL from: https://www.postgresql.org/download/windows/

## 2. Create Database

Open Command Prompt and run:
```bash
psql -U postgres
CREATE DATABASE teamdashboard;
\q
```

Or use pgAdmin to create a database named `teamdashboard`.

## 3. Install Goose (Migration Tool)

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## 4. Run Database Migrations

```bash
cd backend
goose -dir migrations postgres "user=postgres password=postgres dbname=teamdashboard sslmode=disable" up
```

## 5. Install Go Dependencies

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
