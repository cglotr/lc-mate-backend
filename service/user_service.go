package service

import "github.com/cglotr/lc-mate-backend/leetcode"

//go:generate mockgen -package=service -mock_names=UserService=UserServiceMock -source=./user_service.go -destination=./user_service_mock.go
type UserService interface {
	GetUsers(usernames []string) ([]*leetcode.UserInfo, error)
	UpdateMostOutdatedUser() (*leetcode.UserInfo, error)
}

func ErrorInvalidUser() string { return "error invalid user" }
