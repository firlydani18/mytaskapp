package data

import (
	"errors"
	"firly/mytaskapp/features/project"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) project.ProjectDataInterface {
	return &projectQuery{
		db: database,
	}
}

// Insert implements project.ProjectDataInterface.
func (r *projectQuery) Insert(input project.CoreProject) (uint, error) {
	newProject := MapCoreProjectToProject(input)

	//simpan ke db
	tx := r.db.Create(&newProject)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("project not found")
	}
	return newProject.ID, nil
}

// SelectAll implements project.ProjectDataInterface.
func (r *projectQuery) SelectAll(userID uint) ([]project.CoreProject, error) {
	var dataProject []Project
	tx := r.db.Where("user_id", userID).Find(&dataProject)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("project not found")
	}
	//mapping Project Model to CoreProject
	coreProjectSlice := ListMapProjectToCoreProject(dataProject)
	return coreProjectSlice, nil
}

// Select implements project.ProjectDataInterface.
func (r *projectQuery) Select(projectId uint, userID uint) (project.CoreProject, error) {
	var projectData Project
	tx := r.db.Where("id = ? AND user_id = ?", projectId, userID).First(&projectData)
	if tx.Error != nil {
		return project.CoreProject{}, tx.Error
	}
	tx = r.db.Preload("Tasks").First(&projectData, projectId)
	if tx.Error != nil {
		return project.CoreProject{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return project.CoreProject{}, errors.New("project not found")
	}
	//Mapping Project to CorePproject
	coreProject := MapProjectToCoreProject(projectData)
	return coreProject, nil
}

// Update implements project.ProjectDataInterface.
func (r *projectQuery) Update(projectId uint, userID uint, projectData project.CoreProject) error {
	var project Project
	tx := r.db.Where("id = ? AND user_id = ?", projectId, userID).First(&project)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("project not found")
	}

	//Mapping Project to CoreProject
	updatedProject := MapCoreProjectToProject(projectData)

	// Lakukan pembaruan data proyek dalam database
	tx = r.db.Model(&project).Updates(updatedProject)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}

// Delete implements project.ProjectDataInterface.
func (r *projectQuery) Delete(projectId uint, userID uint) error {
	var project Project
	tx := r.db.Where("id = ? AND user_id = ?", projectId, userID).Delete(&project)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("project not found")
	}
	return nil
}
