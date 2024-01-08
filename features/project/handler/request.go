package handler

import "firly/mytaskapp/features/project"

type ProjectRequest struct {
	Title       string `json:"title" form:"title"`
	UserID      uint   `json:"user_id" form:"user_id"`
	Description string `json:"description" form:"description"`
}

// Mapping dari struct ProjectRequest To struct Core Project
func MapProjReqToCoreProject(req ProjectRequest) project.CoreProject {
	return project.CoreProject{
		Title:       req.Title,
		UserID:      req.UserID,
		Description: req.Description,
	}
}
