package handler

import "firly/mytaskapp/features/task"

type TaskRequest struct {
	Title            string `json:"title" form:"title"`
	ProjectID        uint   `json:"project_id" form:"project_id"`
	CompletionStatus string `json:"completion_status" form:"completion_status"`
}

// Mapping dari struct TaskRequest To struct Core Task
func MapTaskReqToCoreTask(req TaskRequest) task.CoreTask {
	return task.CoreTask{
		Title:            req.Title,
		ProjectID:        req.ProjectID,
		CompletionStatus: req.CompletionStatus,
	}
}
