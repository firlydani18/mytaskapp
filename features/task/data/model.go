package data

import (
	"firly/mytaskapp/features/task"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title            string `gorm:"not null"`
	ProjectID        uint   `gorm:"not null"`
	CompletionStatus string `gorm:"enum('Not Completed', 'Completed')"`
}

// Mapping CoreTask to Task Model
func MapCoreTaskToTask(core task.CoreTask) Task {
	return Task{
		Title:            core.Title,
		ProjectID:        core.ProjectID,
		CompletionStatus: core.CompletionStatus,
	}
}

func MapTaskToCoreTask(model Task) task.CoreTask {
	return task.CoreTask{
		ID:               model.ID,
		Title:            model.Title,
		ProjectID:        model.ProjectID,
		CompletionStatus: model.CompletionStatus,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		DeletedAt:        model.DeletedAt.Time,
	}
}

// mapping Task Model to CoreTask
func ListMapTaskToCoreTask(models []Task) []task.CoreTask {
	var coreTasks []task.CoreTask
	for _, model := range models {
		coreTasks = append(coreTasks, MapTaskToCoreTask(model))
	}
	return coreTasks
}
