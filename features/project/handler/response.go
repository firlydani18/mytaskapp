package handler

import (
	"firly/mytaskapp/features/project"
	_taskData "firly/mytaskapp/features/task"
	"time"
)

type ProjectResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	UserID      uint      `json:"user_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Tasks       []_taskData.CoreTask
}

// Mapping CorePrject to ProjectResponsee
func MapCoreProjToProjRes(core project.CoreProject) ProjectResponse {
	return ProjectResponse{
		ID:          core.ID,
		Title:       core.Title,
		UserID:      core.UserID,
		Description: core.Description,
		CreatedAt:   core.CreatedAt,
		Tasks:       core.Tasks,
	}
}
