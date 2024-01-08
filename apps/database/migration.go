package database

import (
	_projectData "firly/mytaskapp/features/project/data"
	_taskData "firly/mytaskapp/features/task/data"
	_userData "firly/mytaskapp/features/user/data"

	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
	db.AutoMigrate(&_projectData.Project{})
	db.AutoMigrate(&_taskData.Task{})
}
