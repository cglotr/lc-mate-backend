package service

import "github.com/cglotr/lc-mate-backend/leetcode"

type UserService interface {
	GetUser(username string) (*leetcode.UserInfo, error)
	UpdateMostOutdatedUser() (*leetcode.UserInfo, error)
}
