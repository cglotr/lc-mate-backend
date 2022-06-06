package dao_test

import (
	"testing"

	"github.com/cglotr/lc-mate-backend/dao"
	"github.com/cglotr/lc-mate-backend/model"
	"github.com/cglotr/lc-mate-backend/util"
	"github.com/stretchr/testify/assert"
)

func TestQueryMostOutdatedUser(t *testing.T) {
	db := util.SpinUpTestDb()
	userDaoImpl := dao.NewUserDaoImpl(db)
	userDaoImpl.Upsert(&model.UserModel{
		Username: "numb3r5",
		Rating:   0,
		Badge:    "Guardian",
	})
	userDaoImpl.Upsert(&model.UserModel{
		Username: "awice",
		Rating:   0,
		Badge:    "Guardian",
	})

	// second upsert to populate updated_at. 'numb3r5' is currently the most outdated
	userDaoImpl.Upsert(&model.UserModel{
		Username: "numb3r5",
		Rating:   4000,
		Badge:    "Guardian",
	})
	userDaoImpl.Upsert(&model.UserModel{
		Username: "awice",
		Rating:   4000,
		Badge:    "Guardian",
	})

	users, err := userDaoImpl.QueryUsers([]string{"awice", "numb3r5"})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))

	// after this, 'awice' will be the most outdated
	err = userDaoImpl.OutdateUser("awice")

	assert.Nil(t, err)

	user, err := userDaoImpl.QueryMostOutdatedUser()

	assert.Nil(t, err)
	assert.Equal(t, "awice", user.Username)
}

func TestUpsert(t *testing.T) {
	db := util.SpinUpTestDb()
	userDaoImpl := dao.NewUserDaoImpl(db)
	userDaoImpl.Upsert(&model.UserModel{
		Username: "awice",
		Rating:   4000,
		Badge:    "Guardian",
	})
	users, err := userDaoImpl.QueryUsers([]string{"awice"})

	assert.Nil(t, err)
	assert.Equal(t, "awice", users[0].Username)
	assert.Equal(t, 4000, users[0].Rating)

	userDaoImpl.Upsert(&model.UserModel{
		Username: "awice",
		Rating:   5000,
		Badge:    "Guardian",
	})
	users, err = userDaoImpl.QueryUsers([]string{"awice"})

	assert.Nil(t, err)
	assert.Equal(t, "awice", users[0].Username)
	assert.Equal(t, 5000, users[0].Rating)

	userDaoImpl.DeleteUser("awice")
	users, err = userDaoImpl.QueryUsers([]string{"awice"})

	assert.Nil(t, err)
	assert.Equal(t, 0, len(users))
}
