package service

import (
	"github.com/cglotr/lc-mate-backend/dao"
	"github.com/cglotr/lc-mate-backend/leetcode"
)

type UserServiceImpl struct {
	userDao     dao.UserDao
	leetcodeApi leetcode.LeetcodeApi
}

func NewUserServiceImpl(userDao dao.UserDao, leetcodeApi leetcode.LeetcodeApi) *UserServiceImpl {
	return &UserServiceImpl{
		userDao:     userDao,
		leetcodeApi: leetcodeApi,
	}
}

func (u *UserServiceImpl) GetUsers(usernames []string) ([]*leetcode.UserInfo, error) {
	userModels, err := u.userDao.QueryUsers(usernames)
	if err != nil {
		return nil, err
	}
	users := []*leetcode.UserInfo{}
	for _, userModel := range userModels {
		user := &leetcode.UserInfo{
			Username: userModel.Username,
			Rating:   userModel.Rating,
			Rank:     userModel.Badge,
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateMostOutdatedUser() (*leetcode.UserInfo, error) {
	userModel, err := u.userDao.QueryMostOutdatedUser()
	if err != nil {
		return nil, err
	}
	err = u.userDao.Upsert(userModel)
	if err != nil {
		return nil, err
	}
	return &leetcode.UserInfo{
		Username: userModel.Username,
		Rating:   userModel.Rating,
		Rank:     userModel.Badge,
	}, nil
}
