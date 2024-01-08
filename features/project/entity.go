package project

import (
	"firly/mytaskapp/features/task"
	"time"
)

type CoreProject struct {
	ID          uint
	Title       string
	Description string
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Tasks       []task.CoreTask `gorm:"foreignKey:ProjectID"`
}

type ProjectDataInterface interface {
	Insert(input CoreProject) (uint, error)
	SelectAll(userID uint) ([]CoreProject, error)
	Select(projectId uint, userID uint) (CoreProject, error)
	Update(projectId uint, userID uint, projectData CoreProject) error
	Delete(projectId uint, userID uint) error
}

type ProjectServiceInterface interface {
	Create(input CoreProject) (uint, error)
	GetAll(userID uint) ([]CoreProject, error)
	GetById(userID uint, projectId uint) (CoreProject, error)
	UpdateById(projectId uint, userID uint, projectData CoreProject) error
	DeleteById(projectId uint, userID uint) error
}
