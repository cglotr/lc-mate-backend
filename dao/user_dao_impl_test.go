package dao

import (
	"testing"

	"github.com/cglotr/lc-mate-backend/model"
	"github.com/cglotr/lc-mate-backend/util"
	"github.com/stretchr/testify/assert"
)

func TestUpsert(t *testing.T) {
	db := util.SpinUpTestDb()
	userDaoImpl := NewUserDaoImpl(db)
	userDaoImpl.Upsert(&model.UserModel{
		Username: "awice",
		Rating:   3000,
		Badge:    "Guardian",
	})
	users, err := userDaoImpl.QueryUsers([]string{"awice"})

	assert.Nil(t, err)
	assert.Equal(t, "awice", users[0].Username)
	assert.Equal(t, 3000, users[0].Rating)

	userDaoImpl.Upsert(&model.UserModel{
		Username: "awice",
		Rating:   3333,
		Badge:    "Guardian",
	})
	users, err = userDaoImpl.QueryUsers([]string{"awice"})

	assert.Nil(t, err)
	assert.Equal(t, "awice", users[0].Username)
	assert.Equal(t, 3333, users[0].Rating)

	userDaoImpl.DeleteUser("awice")
	users, err = userDaoImpl.QueryUsers([]string{"awice"})

	assert.Nil(t, err)
	assert.Equal(t, 0, len(users))
}
