package service

import (
	"github.com/cglotr/lc-mate-backend/dao"
	"github.com/cglotr/lc-mate-backend/leetcode"
	"github.com/cglotr/lc-mate-backend/model"
	"github.com/hooligram/kifu"
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
	missingUsernames := []string{}
	for _, username := range usernames {
		existing := false
		for _, user := range users {
			if user.Username == username {
				existing = true
			}
		}
		if !existing {
			missingUsernames = append(missingUsernames, username)
		}
	}
	if len(missingUsernames) > 0 {
		kifu.Info("Missing usernames: %v", missingUsernames)
		for _, missingUsername := range missingUsernames {
			err := u.userDao.Upsert(&model.UserModel{
				Username: missingUsername,
			})
			kifu.Error(err.Error())

			// this will prioritize user to be updated by the cron
			err = u.userDao.MoveBackUpdatedAtOneDay(missingUsername)
			kifu.Error(err.Error())
		}
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
