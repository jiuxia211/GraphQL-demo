package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        int64
	Uid       int64
	Title     string
	Content   string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateTask(ctx context.Context, task *Task) (*Task, error) {
	err := DB.WithContext(ctx).Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func GetTasksByUserID(ctx context.Context, userID uint) ([]Task, error) {
	var tasks []Task
	result := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&tasks)
	return tasks, result.Error
}

func GetTaskByID(ctx context.Context, id int64) (*Task, error) {
	var task Task
	result := DB.WithContext(ctx).First(&task, id)
	return &task, result.Error
}

func UpdateTask(ctx context.Context, task *Task) (*Task, error) {
	err := DB.WithContext(ctx).Save(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func DeleteTask(ctx context.Context, id int64) error {
	return DB.WithContext(ctx).Delete(&Task{}, id).Error
}
func GetAllTasks(ctx context.Context) ([]*Task, error) {
	var tasks []*Task
	if err := DB.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
