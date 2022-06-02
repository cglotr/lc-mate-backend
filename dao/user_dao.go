package dao

import "github.com/cglotr/lc-mate-backend/model"

type UserDao interface {
	Upsert(user *model.UserModel) error
	ReadAll() ([]*model.UserModel, error)
	Query(username string) (*model.UserModel, error)
	QueryMostOutdatedUser() (*model.UserModel, error)
}
