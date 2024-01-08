package data

import (
	"errors"
	"firly/mytaskapp/features/task"
	"time"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) task.TaskDataInterface {
	return &taskQuery{
		db: database,
	}
}

// Insert implements task.TaskDataInterface.
func (r *taskQuery) Insert(input task.CoreTask, userID uint) (uint, error) {
	var dataTask []Task

	newTask := MapCoreTaskToTask(input)
	//simpan ke db
	tx := r.db.Model(&dataTask).Create(&newTask)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("task not found")
	}
	return newTask.ID, nil
}

// SelectAll implements task.TaskDataInterface.
func (r *taskQuery) SelectAll(userID uint) ([]task.CoreTask, error) {
	var dataTask []Task
	tx := r.db.Joins("JOIN projects ON tasks.project_id = projects.id").
		Where("projects.user_id = ?", userID).
		First(&dataTask)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("task not found")
	}
	//mapping Task Model to CoreTask
	coreTaskSlice := ListMapTaskToCoreTask(dataTask)
	return coreTaskSlice, nil
}

// Select implements task.TaskDataInterface.
func (r *taskQuery) Select(taskId uint, userID uint) (task.CoreTask, error) {
	var taskData Task
	tx := r.db.Joins("JOIN projects ON tasks.project_id = projects.id").
		Where("tasks.id = ? AND projects.user_id = ?", taskId, userID).
		First(&taskData)
	if tx.Error != nil {
		return task.CoreTask{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task.CoreTask{}, errors.New("task not found")
	}
	//Mapping Task to CorePtask
	coreTask := MapTaskToCoreTask(taskData)
	return coreTask, nil
}

// Update implements task.TaskDataInterface.
func (r *taskQuery) Update(taskId uint, userID uint, taskData task.CoreTask) error {
	var task Task
	tx := r.db.Joins("JOIN projects ON tasks.project_id = projects.id").
		Where("tasks.id = ? AND projects.user_id = ?", taskId, userID).
		First(&task)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("task not found")
	}

	//Mapping Task to CoreTask
	updatedTask := MapCoreTaskToTask(taskData)
	tx = r.db.Model(&task).Updates(updatedTask)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	// Set update the task
	tx = r.db.Model(&task).Updates(updatedTask)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements task.TaskDataInterface.
func (r *taskQuery) Delete(taskId uint, userID uint) error {
	var task Task
	tx := r.db.Joins("JOIN projects ON tasks.project_id = projects.id").
		Where("tasks.id = ? AND projects.user_id = ?", taskId, userID).
		First(&task)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("task not found")
	}
	// Set deleted_at and update the task
	tx = r.db.Model(&task).Update("deleted_at", time.Now())
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
