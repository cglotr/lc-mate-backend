package service

import "github.com/cglotr/lc-mate-backend/leetcode"

type UserService interface {
	GetUsers(usernames []string) ([]*leetcode.UserInfo, error)
	UpdateMostOutdatedUser() (*leetcode.UserInfo, error)
}

func ErrorInvalidUser() string { return "error invalid user" }
