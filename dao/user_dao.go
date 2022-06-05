package dao

import "github.com/cglotr/lc-mate-backend/model"

type UserDao interface {
	Upsert(user *model.UserModel) error
	Query(username string) (*model.UserModel, error)
	QueryUsers(usernames []string) ([]*model.UserModel, error)
	QueryMostOutdatedUser() (*model.UserModel, error)
}
