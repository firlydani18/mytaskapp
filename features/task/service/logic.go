package service

import (
	"errors"
	"firly/mytaskapp/features/project"
	"firly/mytaskapp/features/task"
)

type TaskService struct {
	taskRepo    task.TaskDataInterface
	projectRepo project.ProjectDataInterface
}

func New(repo task.TaskDataInterface, projectRepo project.ProjectDataInterface) task.TaskServiceInterface {
	return &TaskService{
		taskRepo:    repo,
		projectRepo: projectRepo,
	}
}

// CreateTask implements task.TaskServiceInterface.
func (s *TaskService) Create(inputTask task.CoreTask, userID uint, projectId uint) (uint, error) {
	// Cek apakah user sudah login atau belum
	if userID == 0 {
		return 0, errors.New("user not logged in")
	}

	// Dapatkan proyek terkait berdasarkan projectID
	project, err := s.projectRepo.Select(projectId, userID)
	if err != nil {
		return 0, err
	}

	// Pastikan userID dalam token sesuai dengan userID dalam proyek
	if userID != project.UserID {
		return 0, errors.New("user does not have access to this project")
	}

	taskID, err := s.taskRepo.Insert(inputTask, userID)
	if err != nil {
		return 0, err
	}
	return taskID, nil
}

// GetAllTask implements task.TaskServiceInterface.
func (s *TaskService) GetAll(userID uint) ([]task.CoreTask, error) {
	result, err := s.taskRepo.SelectAll(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTaskById implements task.TaskServiceInterface.
func (s *TaskService) GetById(taskId uint, userID uint) (task.CoreTask, error) {
	result, err := s.taskRepo.Select(taskId, userID)
	if err != nil {
		return task.CoreTask{}, err
	}
	return result, nil
}

// UpdateTaskById implements task.TaskServiceInterface.
func (s *TaskService) UpdateById(taskId uint, userID uint, taskData task.CoreTask) error {
	err := s.taskRepo.Update(taskId, userID, taskData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTaskById implements task.TaskServiceInterface.
func (s *TaskService) DeleteById(taskId uint, userID uint) error {
	err := s.taskRepo.Delete(taskId, userID)
	if err != nil {
		return err
	}
	return nil
}
