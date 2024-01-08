package data

import (
	"firly/mytaskapp/features/project"
	"firly/mytaskapp/features/task"
	_taskData "firly/mytaskapp/features/task/data"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title       string           `gorm:"column:title;not null"`
	Description string           `gorm:"column:description"`
	UserID      uint             `gorm:"column:user_id;not null"`
	Tasks       []_taskData.Task `gorm:"foreignKey:ProjectID"`
}

// Mapping CoreProject to Project Model
func MapCoreProjectToProject(core project.CoreProject) Project {
	//convert CoreTask TO Task Model
	var taskModels []_taskData.Task
	for _, task := range core.Tasks {
		taskModels = append(taskModels, _taskData.MapCoreTaskToTask(task))
	}
	return Project{
		Title:       core.Title,
		UserID:      core.UserID,
		Description: core.Description,
		Tasks:       taskModels,
	}
}

// Mapping Project Model to CoreProject
func MapProjectToCoreProject(model Project) project.CoreProject {
	var coreTasks []task.CoreTask
	for _, task := range model.Tasks {
		coreTasks = append(coreTasks, _taskData.MapTaskToCoreTask(task))
	}
	return project.CoreProject{
		ID:          model.ID,
		Title:       model.Title,
		UserID:      model.UserID,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt.Time,
		Tasks:       coreTasks,
	}
}

// mapping Project Model to CoreProject
func ListMapProjectToCoreProject(models []Project) []project.CoreProject {
	var coreProjects []project.CoreProject
	for _, model := range models {
		var coreTasks []task.CoreTask
		for _, task := range model.Tasks {
			coreTasks = append(coreTasks, _taskData.MapTaskToCoreTask(task))
		}
		coreProjects = append(coreProjects, project.CoreProject{
			ID:          model.ID,
			Title:       model.Title,
			Description: model.Description,
			UserID:      model.UserID,
			CreatedAt:   model.CreatedAt,
			UpdatedAt:   model.UpdatedAt,
			DeletedAt:   model.DeletedAt.Time,
			Tasks:       coreTasks,
		})
	}
	return coreProjects
}
