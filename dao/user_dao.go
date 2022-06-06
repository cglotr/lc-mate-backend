package dao

import "github.com/cglotr/lc-mate-backend/model"

//go:generate mockgen -package=dao -mock_names=UserDao=UserDaoMock -source=./user_dao.go -destination=./user_dao_mock.go
type UserDao interface {
	Upsert(user *model.UserModel) error
	QueryUsers(usernames []string) ([]*model.UserModel, error)
	QueryMostOutdatedUser() (*model.UserModel, error)
	OutdateUser(username string) error
	DeleteUser(username string) error
}

func ErrorUserNotFound() string { return "error user not found" }
