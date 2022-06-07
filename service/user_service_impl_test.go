package service_test

import (
	"errors"
	"testing"

	"github.com/cglotr/lc-mate-backend/dao"
	"github.com/cglotr/lc-mate-backend/leetcode"
	"github.com/cglotr/lc-mate-backend/model"
	"github.com/cglotr/lc-mate-backend/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMostOutdatedUserUpsertError(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := dao.NewUserDaoMock(ctrl)
	leetcodeApiMock := leetcode.NewLeetcodeApiMock(ctrl)
	userServiceImpl := service.NewUserServiceImpl(userDaoMock, leetcodeApiMock)

	userDaoMock.EXPECT().
		QueryMostOutdatedUser().
		Return(&model.UserModel{Username: "awice"}, nil)
	leetcodeApiMock.EXPECT().
		GetUserInfo(gomock.Eq("awice")).
		Return(&leetcode.UserInfo{Username: "awice"}, nil)
	userDaoMock.EXPECT().
		Upsert(gomock.Eq(&model.UserModel{Username: "awice"})).
		Return(errors.New(dao.ErrorUserNotFound()))

	_, err := userServiceImpl.UpdateMostOutdatedUser()

	assert.Equal(t, "error user not found", err.Error())
}

func TestErrorDeletingInvalidUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := dao.NewUserDaoMock(ctrl)
	leetcodeApiMock := leetcode.NewLeetcodeApiMock(ctrl)
	userServiceImpl := service.NewUserServiceImpl(userDaoMock, leetcodeApiMock)

	userDaoMock.EXPECT().
		QueryMostOutdatedUser().
		Return(&model.UserModel{Username: "awice"}, nil)
	leetcodeApiMock.EXPECT().
		GetUserInfo(gomock.Eq("awice")).
		Return(nil, errors.New(service.ErrorInvalidUser()))
	userDaoMock.EXPECT().
		DeleteUser(gomock.Eq("awice")).
		Return(errors.New(service.ErrorInvalidUser()))

	_, err := userServiceImpl.UpdateMostOutdatedUser()

	assert.Equal(t, "error invalid user", err.Error())
}

func TestUpdatingUserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := dao.NewUserDaoMock(ctrl)
	leetcodeApiMock := leetcode.NewLeetcodeApiMock(ctrl)
	userServiceImpl := service.NewUserServiceImpl(userDaoMock, leetcodeApiMock)

	userDaoMock.EXPECT().
		QueryMostOutdatedUser().
		Return(nil, errors.New(dao.ErrorUserNotFound()))

	_, err := userServiceImpl.UpdateMostOutdatedUser()

	assert.Equal(t, "error user not found", err.Error())
}

func TestUserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := dao.NewUserDaoMock(ctrl)
	leetcodeApiMock := leetcode.NewLeetcodeApiMock(ctrl)
	userServiceImpl := service.NewUserServiceImpl(userDaoMock, leetcodeApiMock)

	userDaoMock.EXPECT().
		QueryUsers(gomock.Eq([]string{"awice"})).
		Return(nil, errors.New(dao.ErrorUserNotFound()))

	_, err := userServiceImpl.GetUsers([]string{"awice"})

	assert.Equal(t, "error user not found", err.Error())
}

func TestDeletingInvalidUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := dao.NewUserDaoMock(ctrl)
	leetcodeApiMock := leetcode.NewLeetcodeApiMock(ctrl)
	userServiceImpl := service.NewUserServiceImpl(userDaoMock, leetcodeApiMock)

	userDaoMock.EXPECT().
		QueryMostOutdatedUser().
		Return(&model.UserModel{Username: "awice"}, nil)
	leetcodeApiMock.EXPECT().
		GetUserInfo(gomock.Eq("awice")).
		Return(nil, errors.New(service.ErrorInvalidUser()))
	userDaoMock.EXPECT().
		DeleteUser(gomock.Eq("awice")).
		Return(nil)

	userInfo, err := userServiceImpl.UpdateMostOutdatedUser()

	assert.Nil(t, userInfo)
	assert.Equal(t, "error invalid user", err.Error())
}

func TestUpdateMostOutdatedUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := dao.NewUserDaoMock(ctrl)
	leetcodeApiMock := leetcode.NewLeetcodeApiMock(ctrl)
	userServiceImpl := service.NewUserServiceImpl(userDaoMock, leetcodeApiMock)

	userDaoMock.EXPECT().
		QueryMostOutdatedUser().
		Return(&model.UserModel{Username: "awice"}, nil)
	leetcodeApiMock.EXPECT().
		GetUserInfo(gomock.Eq("awice")).
		Return(&leetcode.UserInfo{Username: "awice"}, nil)
	userDaoMock.EXPECT().
		Upsert(gomock.Eq(&model.UserModel{Username: "awice"})).
		Return(nil)

	userInfo, err := userServiceImpl.UpdateMostOutdatedUser()

	assert.Nil(t, err)
	assert.Equal(t, "awice", userInfo.Username)
	assert.Equal(t, 0, userInfo.Rating)
	assert.Equal(t, "", userInfo.Rank)
}

func TestMissingUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := dao.NewUserDaoMock(ctrl)
	leetcodeApiMock := leetcode.NewLeetcodeApiMock(ctrl)
	userServiceImpl := service.NewUserServiceImpl(userDaoMock, leetcodeApiMock)

	userDaoMock.EXPECT().
		QueryUsers(gomock.Eq([]string{"awice", "numb3r5"})).
		Return([]*model.UserModel{}, nil).
		Times(1)
	userDaoMock.EXPECT().
		Upsert(gomock.Eq(&model.UserModel{Username: "awice"})).
		Return(nil).
		Times(1)
	userDaoMock.EXPECT().
		Upsert(gomock.Eq(&model.UserModel{Username: "numb3r5"})).
		Return(errors.New(dao.ErrorUserNotFound())).
		Times(1)
	userDaoMock.EXPECT().
		OutdateUser(gomock.Eq("awice")).Return(errors.New(dao.ErrorUserNotFound())).
		Times(1)
	userDaoMock.EXPECT().
		OutdateUser(gomock.Eq("numb3r5")).Return(nil).
		Times(1)

	users, err := userServiceImpl.GetUsers([]string{"awice", "numb3r5"})

	assert.Nil(t, err)
	assert.Equal(t, 0, len(users))
}

func TestExistingUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := dao.NewUserDaoMock(ctrl)
	leetcodeApiMock := leetcode.NewLeetcodeApiMock(ctrl)
	userServiceImpl := service.NewUserServiceImpl(userDaoMock, leetcodeApiMock)

	userDaoMock.EXPECT().
		QueryUsers(gomock.Eq([]string{"awice", "numb3r5"})).
		Return([]*model.UserModel{
			{
				Username: "awice",
			},
			{
				Username: "numb3r5",
			},
		}, nil).
		Times(1)

	users, err := userServiceImpl.GetUsers([]string{"awice", "numb3r5"})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}
