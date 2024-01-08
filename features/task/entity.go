package task

import "time"

type CoreTask struct {
	ID               uint
	Title            string
	ProjectID        uint
	CompletionStatus string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
type TaskDataInterface interface {
	Insert(input CoreTask, taskId uint) (uint, error)
	SelectAll(userID uint) ([]CoreTask, error)
	Select(taskId uint, userID uint) (CoreTask, error)
	Update(taskId uint, userID uint, taskData CoreTask) error
	Delete(taskId, userID uint) error
}

type TaskServiceInterface interface {
	Create(input CoreTask, userID uint, projectId uint) (uint, error)
	GetAll(userID uint) ([]CoreTask, error)
	GetById(taskId uint, userID uint) (CoreTask, error)
	UpdateById(taskId uint, userID uint, taskData CoreTask) error
	DeleteById(taskId uint, userID uint) error
}
