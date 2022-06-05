package dao

import "github.com/cglotr/lc-mate-backend/model"

type UserDao interface {
	Upsert(user *model.UserModel) error
	QueryUsers(usernames []string) ([]*model.UserModel, error)
	QueryMostOutdatedUser() (*model.UserModel, error)
	MoveBackUpdatedAtOneDay(username string) error
}

func ErrUserNotFound() string { return "user not found" }
