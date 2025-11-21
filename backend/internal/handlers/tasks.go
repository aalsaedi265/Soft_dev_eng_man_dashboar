package handlers

import (
	"net/http"

	"github.com/aalsa/management_dashboard/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TaskHandler struct {
	db *sqlx.DB
}

func NewTaskHandler(db *sqlx.DB) *TaskHandler {
	return &TaskHandler{db: db}
}

func (h *TaskHandler) GetAll(c *gin.Context) {
	var tasks []models.Task
	query := `SELECT * FROM tasks ORDER BY due_date, created_at DESC`

	if err := h.db.Select(&tasks, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	query := `SELECT * FROM tasks WHERE id = $1`
	if err := h.db.Get(&task, query, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Create(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = uuid.New().String()

	query := `INSERT INTO tasks (id, title, description, project_id, assigned_to, status, priority, due_date)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING created_at, updated_at`

	err := h.db.QueryRow(query, task.ID, task.Title, task.Description,
		task.ProjectID, task.AssignedTo, task.Status, task.Priority, task.DueDate).
		Scan(&task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE tasks SET title = $1, description = $2, assigned_to = $3,
	          status = $4, priority = $5, due_date = $6, updated_at = CURRENT_TIMESTAMP
	          WHERE id = $7`

	result, err := h.db.Exec(query, task.Title, task.Description, task.AssignedTo,
		task.Status, task.Priority, task.DueDate, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (h *TaskHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM tasks WHERE id = $1`
	result, err := h.db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
