package handler

import (
	"firly/mytaskapp/features/task"
	"time"
)

type TaskResponse struct {
	Title            string    `json:"title"`
	ProjectID        uint      `json:"project_id"`
	CompletionStatus string    `json:"completion_status"`
	CreatedAt        time.Time `json:"created_at"`
}

// Mapping CorePrject to TaskResponsee
func MapCoreTaskToTaskRes(core task.CoreTask) TaskResponse {
	return TaskResponse{
		Title:     core.Title,
		ProjectID: core.ProjectID,
		CreatedAt: core.CreatedAt,
	}
}
