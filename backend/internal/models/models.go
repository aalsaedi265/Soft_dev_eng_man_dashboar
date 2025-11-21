package models

import "time"

type Employee struct {
	ID         string    `db:"id" json:"id"`
	Email      string    `db:"email" json:"email"`
	FullName   string    `db:"full_name" json:"full_name"`
	Role       string    `db:"role" json:"role"`
	Department string    `db:"department" json:"department"`
	HireDate   string    `db:"hire_date" json:"hire_date"`
	Status     string    `db:"status" json:"status"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

type Project struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	StartDate   string    `db:"start_date" json:"start_date"`
	EndDate     *string   `db:"end_date" json:"end_date"`
	Status      string    `db:"status" json:"status"`
	Budget      *float64  `db:"budget" json:"budget"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Task struct {
	ID          string     `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	ProjectID   string     `db:"project_id" json:"project_id"`
	AssignedTo  *string    `db:"assigned_to" json:"assigned_to"`
	Status      string     `db:"status" json:"status"`
	Priority    string     `db:"priority" json:"priority"`
	DueDate     *string    `db:"due_date" json:"due_date"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	CompletedAt *time.Time `db:"completed_at" json:"completed_at"`
}

type TimeLog struct {
	ID         string    `db:"id" json:"id"`
	EmployeeID string    `db:"employee_id" json:"employee_id"`
	TaskID     string    `db:"task_id" json:"task_id"`
	Hours      float64   `db:"hours" json:"hours"`
	LogDate    string    `db:"log_date" json:"log_date"`
	Notes      string    `db:"notes" json:"notes"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

type User struct {
	ID           string     `db:"id" json:"id"`
	EmployeeID   *string    `db:"employee_id" json:"employee_id"`
	Email        string     `db:"email" json:"email"`
	PasswordHash string     `db:"password_hash" json:"-"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	LastLogin    *time.Time `db:"last_login" json:"last_login"`
}
