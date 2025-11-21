package handlers

import (
	"net/http"

	"github.com/aalsa/management_dashboard/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProjectHandler struct {
	db *sqlx.DB
}

func NewProjectHandler(db *sqlx.DB) *ProjectHandler {
	return &ProjectHandler{db: db}
}

func (h *ProjectHandler) GetAll(c *gin.Context) {
	var projects []models.Project
	query := `SELECT * FROM projects ORDER BY start_date DESC`

	if err := h.db.Select(&projects, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	query := `SELECT * FROM projects WHERE id = $1`
	if err := h.db.Get(&project, query, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) Create(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project.ID = uuid.New().String()

	query := `INSERT INTO projects (id, name, description, start_date, end_date, status, budget)
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING created_at, updated_at`

	err := h.db.QueryRow(query, project.ID, project.Name, project.Description,
		project.StartDate, project.EndDate, project.Status, project.Budget).
		Scan(&project.CreatedAt, &project.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *ProjectHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE projects SET name = $1, description = $2, start_date = $3,
	          end_date = $4, status = $5, budget = $6, updated_at = CURRENT_TIMESTAMP
	          WHERE id = $7`

	result, err := h.db.Exec(query, project.Name, project.Description, project.StartDate,
		project.EndDate, project.Status, project.Budget, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated"})
}

func (h *ProjectHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM projects WHERE id = $1`
	result, err := h.db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
