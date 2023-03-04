package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {

	var taskList []entity.Task
	err := r.db.WithContext(ctx).Where("user_id = ?", id).Find(&taskList).Error

	return taskList, err
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {

	err = r.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {

	var task entity.Task
	err := r.db.WithContext(ctx).Find(&task, id).Error

	return task, err
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {

	var taskList []entity.Task
	err := r.db.WithContext(ctx).Where("category_id = ?", catId).Find(&taskList).Error

	return taskList, err
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {

	return r.db.WithContext(ctx).Model(&task).Updates(&task).Error
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {

	return r.db.WithContext(ctx).Delete(entity.Task{}, id).Error
}
