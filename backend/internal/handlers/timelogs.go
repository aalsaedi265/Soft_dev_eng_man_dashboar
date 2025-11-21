package handlers

import (
	"net/http"

	"github.com/aalsa/management_dashboard/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TimeLogHandler struct {
	db *sqlx.DB
}

func NewTimeLogHandler(db *sqlx.DB) *TimeLogHandler {
	return &TimeLogHandler{db: db}
}

func (h *TimeLogHandler) GetAll(c *gin.Context) {
	var logs []models.TimeLog
	query := `SELECT * FROM time_logs ORDER BY log_date DESC, created_at DESC`

	if err := h.db.Select(&logs, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (h *TimeLogHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	var log models.TimeLog

	query := `SELECT * FROM time_logs WHERE id = $1`
	if err := h.db.Get(&log, query, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time log not found"})
		return
	}

	c.JSON(http.StatusOK, log)
}

func (h *TimeLogHandler) Create(c *gin.Context) {
	var log models.TimeLog
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.ID = uuid.New().String()

	query := `INSERT INTO time_logs (id, employee_id, task_id, hours, log_date, notes)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING created_at`

	err := h.db.QueryRow(query, log.ID, log.EmployeeID, log.TaskID, log.Hours, log.LogDate, log.Notes).Scan(&log.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, log)
}

func (h *TimeLogHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var log models.TimeLog

	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE time_logs SET hours = $1, log_date = $2, notes = $3 WHERE id = $4`
	result, err := h.db.Exec(query, log.Hours, log.LogDate, log.Notes, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time log not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Time log updated"})
}

func (h *TimeLogHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM time_logs WHERE id = $1`
	result, err := h.db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time log not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Time log deleted"})
}

func (h *TimeLogHandler) GetEmployeeHours(c *gin.Context) {
	employeeID := c.Param("id")

	var total float64
	query := `SELECT COALESCE(SUM(hours), 0) FROM time_logs WHERE employee_id = $1`

	if err := h.db.Get(&total, query, employeeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"employee_id": employeeID, "total_hours": total})
}

func (h *TimeLogHandler) GetTaskHours(c *gin.Context) {
	taskID := c.Param("id")

	var total float64
	query := `SELECT COALESCE(SUM(hours), 0) FROM time_logs WHERE task_id = $1`

	if err := h.db.Get(&total, query, taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task_id": taskID, "total_hours": total})
}
