# Development Guidelines

This document outlines the core principles and practices for developing the Team Dashboard project.

## Core Principles

### 1. Security First - No Hardcoded Secrets

**Rule:** All sensitive data must be stored in environment variables, never in code.

**Implementation:**
- Use `.env` files for local development
- Reference secrets via `os.Getenv()` (Go) or `process.env` (Node.js)
- Always add `.env` to `.gitignore`
- Scripts (batch files, shell scripts) must read from `.env` instead of hardcoding values

**Examples:**
```go
// ✅ GOOD
password := os.Getenv("DB_PASSWORD")

// ❌ BAD
password := "Thegreatone"
```

```bash
# ✅ GOOD
for /f "tokens=1,2 delims==" %%a in (.env) do (
    if "%%a"=="DB_PASSWORD" set DB_PASSWORD=%%b
)

# ❌ BAD
set DB_PASSWORD=Thegreatone
```

### 2. KISS Principle (Keep It Simple, Stupid)

**Rule:** Don't over-engineer. Build only what's needed right now.

**Guidelines:**
- Keep files consolidated until they exceed 100-150 lines
- Three similar lines of code is better than a premature abstraction
- Don't create helpers/utilities for one-time operations
- Avoid feature flags or abstractions for hypothetical future requirements

**Examples:**
```typescript
// ✅ GOOD - Simple and direct
const employees = await fetchEmployees();
const projects = await fetchProjects();
const tasks = await fetchTasks();

// ❌ BAD - Over-engineered
class DataFetcher<T> {
  constructor(private endpoint: string) {}
  async fetch(): Promise<T[]> { /* ... */ }
}
const employeeFetcher = new DataFetcher<Employee>('/employees');
```

### 3. No Unnecessary Features

**Rule:** Only implement what the user explicitly requests. Don't add "nice to have" features.

**Don't Add:**
- Extra error handling for scenarios that can't happen
- Comments or docstrings to unchanged code
- Configurability that isn't needed yet
- Validation at internal boundaries (trust your own code)

**Do Add:**
- Validation at system boundaries (user input, external APIs)
- Error handling for actual failure modes
- Comments only where logic isn't self-evident

### 4. Clean Code Hygiene

**Rule:** Delete unused code completely. No remnants.

**Guidelines:**
- ❌ Don't rename unused variables with `_` prefix
- ❌ Don't leave `// removed` comments
- ❌ Don't add backwards-compatibility shims unless explicitly needed
- ✅ Delete unused imports, variables, functions completely
- ✅ Remove mock data when switching to real APIs

**Examples:**
```typescript
// ❌ BAD
const _oldFunction = () => { /* ... */ }; // Not used anymore
const mockData = []; // Removed, now using API

// ✅ GOOD
// (Just delete them entirely)
```

### 5. Focus on "Why" Not "What"

**Rule:** Commit messages and comments explain purpose, not implementation.

**Examples:**
```bash
# ✅ GOOD
git commit -m "Add JWT authentication to secure API endpoints"

# ❌ BAD
git commit -m "Added new middleware function that checks token in header"
```

## Code Organization

### File Structure
```
management_dashboard/
├── backend/
│   ├── cmd/server/main.go        # Entry point
│   ├── internal/
│   │   ├── db/db.go              # Database connection
│   │   ├── handlers/             # HTTP handlers
│   │   ├── middleware/           # Auth, CORS, etc.
│   │   └── models/               # Data structures
│   ├── migrations/               # SQL migrations
│   ├── .env                      # Secrets (NOT in git)
│   ├── go.mod
│   └── *.bat                     # Windows setup scripts
│
├── frontend/vite-project/
│   └── src/
│       ├── App.tsx               # All components (until too large)
│       ├── App.css               # All styles
│       └── main.tsx              # Entry point
│
├── DEVELOPMENT_GUIDELINES.md    # This file
├── SETUP.md                      # Setup instructions
└── ReadMe                        # Project overview
```

### When to Split Files

**Don't split until:**
- A single file exceeds 150 lines
- A component is reused in multiple places
- Code becomes genuinely hard to navigate

**Premature file splitting adds complexity without benefit.**

## Technology Stack

### Backend
- **Language:** Go 1.21+
- **Framework:** Gin (HTTP server)
- **Database:** PostgreSQL 18+ with sqlx (no ORM)
- **Migrations:** Goose
- **Auth:** JWT with bcrypt
- **Environment:** godotenv for `.env` loading

### Frontend
- **Framework:** React 18 with TypeScript
- **Build Tool:** Vite
- **Routing:** React Router
- **HTTP:** Native Fetch API (no axios)
- **Styling:** Plain CSS

### Database
- **Primary Keys:** UUIDs (distributed-friendly)
- **Constraints:** Foreign keys for data integrity
- **Indexes:** On commonly queried fields
- **Migrations:** Versioned SQL files managed by Goose

## Development Workflow

### 1. Environment Setup
```bash
# Backend
cd backend
# Ensure .env exists with correct credentials
setup-db.bat         # Create database
run-migrations.bat   # Apply schema
start-server.bat     # Run server (localhost:8080)

# Frontend
cd frontend/vite-project
npm install
npm run dev          # Run dev server (localhost:5173)
```

### 2. Database Migrations

**Creating a migration:**
```bash
cd backend
goose -dir migrations create my_migration_name sql
```

**Running migrations:**
```bash
run-migrations.bat  # Uses credentials from .env
```

**Never:**
- Edit existing migration files
- Run SQL manually outside of migrations
- Skip the migration system

### 3. API Development

**Order of Implementation:**
1. Define types/models
2. Write SQL migrations
3. Implement handlers
4. Add routes to main.go
5. Test with curl/Postman
6. Connect frontend

**Handler Pattern:**
```go
func (h *Handler) GetAll(c *gin.Context) {
    var items []Item
    err := h.db.Select(&items, "SELECT * FROM items")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, items)
}
```

### 4. Frontend Development

**State Management:**
- Use `useState` for component state
- Use `useEffect` for data fetching
- Use `useContext` for authentication
- No Redux/Zustand until complexity demands it

**API Calls:**
```typescript
const { token } = useAuth();
const response = await fetch(`${API_URL}/endpoint`, {
  headers: { Authorization: `Bearer ${token}` }
});
const data = await response.json();
```

## Security Guidelines

### Authentication Flow
1. User registers/logs in via `/api/auth/register` or `/api/auth/login`
2. Backend returns JWT token
3. Frontend stores token in `localStorage`
4. All subsequent requests include `Authorization: Bearer <token>` header
5. Backend middleware validates token before processing requests

### Password Handling
- **Never** log passwords
- **Always** use bcrypt for hashing
- **Never** return password hashes in API responses
- Minimum length: 6 characters (enforced in frontend)

### CORS Configuration
```go
AllowOrigins:     []string{"http://localhost:5173"},
AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
AllowCredentials: true,
```

## Git Workflow

### Commit Messages
- Keep to 1-2 sentences
- Focus on "why" not "what"
- Use imperative mood: "Add feature" not "Added feature"

### What to Commit
- ✅ Source code
- ✅ Configuration templates
- ✅ Migration files
- ✅ Documentation

### What NOT to Commit
- ❌ `.env` files
- ❌ `node_modules/`
- ❌ Binary files
- ❌ IDE-specific settings
- ❌ Temporary files

## Testing Guidelines

### Manual Testing Checklist
1. Register new user
2. Login with credentials
3. View all entities (employees, projects, tasks)
4. Create new entity
5. Update existing entity
6. Delete entity
7. Verify dashboard statistics update
8. Logout and verify redirect
9. Try accessing protected route without token

### API Testing
```bash
# Health check
curl http://localhost:8080/api/health

# Register
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Get employees (with token)
curl http://localhost:8080/api/employees \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Common Pitfalls to Avoid

1. **Don't hardcode credentials** - Always use environment variables
2. **Don't create files prematurely** - Start with one file, split when needed
3. **Don't add unused features** - Only build what's requested
4. **Don't leave commented code** - Delete it completely
5. **Don't trust user input** - Always validate at system boundaries
6. **Don't skip migrations** - Every schema change needs a migration
7. **Don't ignore errors** - Handle them appropriately (log, return, recover)

## Performance Considerations

### Database
- Add indexes on foreign keys
- Use `SELECT` only needed columns (not `SELECT *` in production)
- Batch operations when possible

### Frontend
- Fetch data on component mount
- Show loading states
- Handle empty states gracefully
- Debounce expensive operations (search, etc.)

### Backend
- Use connection pooling (sqlx default)
- Set appropriate timeouts
- Return paginated results for large datasets (when needed)

## Next Steps After Phase 1

Once CRUD operations work end-to-end:
1. Add form validation
2. Implement "Add New" forms
3. Implement "Edit" functionality
4. Add proper error messages
5. Improve loading states
6. Add WebSocket for real-time updates (Phase 2)
7. Implement file uploads (Phase 2)
8. Add analytics queries (Phase 2)

## Questions?

When in doubt:
1. Keep it simple
2. Don't add what's not requested
3. Delete unused code
4. Protect secrets
5. Focus on making it work first, optimize later
