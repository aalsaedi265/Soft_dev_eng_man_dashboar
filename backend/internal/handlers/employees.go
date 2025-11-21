package handlers

import (
	"net/http"

	"github.com/aalsa/management_dashboard/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type EmployeeHandler struct {
	db *sqlx.DB
}

func NewEmployeeHandler(db *sqlx.DB) *EmployeeHandler {
	return &EmployeeHandler{db: db}
}

func (h *EmployeeHandler) GetAll(c *gin.Context) {
	var employees []models.Employee
	query := `SELECT * FROM employees ORDER BY full_name`

	if err := h.db.Select(&employees, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	query := `SELECT * FROM employees WHERE id = $1`
	if err := h.db.Get(&employee, query, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.ID = uuid.New().String()

	query := `INSERT INTO employees (id, email, full_name, role, department, hire_date, status)
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING created_at, updated_at`

	err := h.db.QueryRow(query, employee.ID, employee.Email, employee.FullName,
		employee.Role, employee.Department, employee.HireDate, employee.Status).
		Scan(&employee.CreatedAt, &employee.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE employees SET email = $1, full_name = $2, role = $3,
	          department = $4, hire_date = $5, status = $6, updated_at = CURRENT_TIMESTAMP
	          WHERE id = $7`

	result, err := h.db.Exec(query, employee.Email, employee.FullName, employee.Role,
		employee.Department, employee.HireDate, employee.Status, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated"})
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM employees WHERE id = $1`
	result, err := h.db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
