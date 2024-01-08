package service

import (
	"firly/mytaskapp/features/project"
)

type projectService struct {
	projectRepo project.ProjectDataInterface
	// validate    *validator.Validate
}

func New(repo project.ProjectDataInterface) project.ProjectServiceInterface {
	return &projectService{
		projectRepo: repo,
	}
}

// Create implements project.ProjectServiceInterface.
func (s *projectService) Create(input project.CoreProject) (uint, error) {
	projectID, err := s.projectRepo.Insert(input)
	if err != nil {
		return 0, err
	}
	return projectID, nil
}

// GetAll implements project.ProjectServiceInterface.
func (s *projectService) GetAll(userID uint) ([]project.CoreProject, error) {
	result, err := s.projectRepo.SelectAll(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetById implements project.ProjectServiceInterface.
func (s *projectService) GetById(projectId uint, userID uint) (project.CoreProject, error) {
	result, err := s.projectRepo.Select(projectId, userID)
	if err != nil {
		return project.CoreProject{}, err
	}
	return result, nil
}

// UpdateById implements project.ProjectServiceInterface.
func (s *projectService) UpdateById(projectId uint, userID uint, projectData project.CoreProject) error {
	err := s.projectRepo.Update(projectId, userID, projectData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteById implements project.ProjectServiceInterface.
func (s *projectService) DeleteById(projectId uint, userID uint) error {
	err := s.projectRepo.Delete(projectId, userID)
	if err != nil {
		return err
	}
	return nil
}
