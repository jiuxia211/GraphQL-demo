package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateUser(ctx context.Context, user *User) (*User, error) {
	err := DB.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(ctx context.Context, id int64) (*User, error) {
	var user User
	result := DB.WithContext(ctx).First(&user, id)
	return &user, result.Error
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	result := DB.WithContext(ctx).Where("username = ?", username).First(&user)
	return &user, result.Error
}
