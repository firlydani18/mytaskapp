package data

import (
	"firly/mytaskapp/features/project"
	"firly/mytaskapp/features/project/data"
	_projectData "firly/mytaskapp/features/project/data"
	"firly/mytaskapp/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string         `gorm:"column:name;not null"`
	Email       string         `gorm:"column:email;not null;unique"`
	Password    string         `gorm:"column:password;not null"`
	PhoneNumber string         `gorm:"column:phone_number;"`
	Projects    []data.Project `gorm:"foreignKey:UserID"`
}

// mapping coreUser to User
func MapCoreUsertoUser(coreUser user.CoreUser) User {
	var projects []_projectData.Project
	for _, coreProject := range coreUser.Projects {
		projects = append(projects, data.MapCoreProjectToProject(coreProject))
	}
	return User{
		Name:        coreUser.Name,
		Email:       coreUser.Email,
		Password:    coreUser.Password,
		PhoneNumber: coreUser.PhoneNumber,
		Projects:    projects,
	}
}

// Mapping User to CoreUser
func MapUserToCoreUser(dataModel User) user.CoreUser {
	var coreProjects []project.CoreProject
	for _, projectModel := range dataModel.Projects {
		coreProjects = append(coreProjects, data.MapProjectToCoreProject(projectModel))
	}
	return user.CoreUser{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		PhoneNumber: dataModel.PhoneNumber,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
		DeletedAt:   dataModel.DeletedAt.Time,
		Projects:    coreProjects,
	}
}

// Mapping dari []User ke []CoreUser
func ListMapUserToCoreUser(users []User) []user.CoreUser {
	var coreUsers []user.CoreUser
	for _, userModel := range users {
		coreUsers = append(coreUsers, MapUserToCoreUser(userModel))
	}
	return coreUsers
}
