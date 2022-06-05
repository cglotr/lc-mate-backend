package dao

import "github.com/cglotr/lc-mate-backend/model"

type UserDao interface {
	Upsert(user *model.UserModel) error
	QueryUsers(usernames []string) ([]*model.UserModel, error)
	QueryMostOutdatedUser() (*model.UserModel, error)
	OutdateUser(username string) error
	DeleteUser(username string) error
}

func ErrorUserNotFound() string { return "error user not found" }
