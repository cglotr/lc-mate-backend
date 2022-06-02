package service

import (
	"github.com/cglotr/lc-mate-backend/dao"
	"github.com/cglotr/lc-mate-backend/leetcode"
	"github.com/cglotr/lc-mate-backend/model"
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

func (u *UserServiceImpl) GetUser(username string) (*leetcode.UserInfo, error) {
	user, err := u.userDao.Query(username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		userInfo := leetcode.UserInfo{
			Username: user.Username,
			Rating:   user.Rating,
			Rank:     user.Badge,
		}
		return &userInfo, nil
	} else {
		userInfo, err := u.leetcodeApi.GetUserInfo(username)
		if err != nil {
			return nil, err
		}
		u.userDao.Upsert(&model.UserModel{
			Username: userInfo.Username,
			Rating:   userInfo.Rating,
			Badge:    userInfo.Rank,
		})
		return userInfo, nil
	}
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

func (u *UserServiceImpl) GetUsers() ([]*leetcode.UserInfo, error) {
	userModels, err := u.userDao.ReadAll()
	if err != nil {
		return nil, err
	}
	users := []*leetcode.UserInfo{}
	for _, userModel := range userModels {
		users = append(users, &leetcode.UserInfo{
			Username: userModel.Username,
			Rating:   userModel.Rating,
			Rank:     userModel.Badge,
		})
	}
	return users, nil
}
