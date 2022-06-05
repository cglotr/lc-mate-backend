package dao

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/cglotr/lc-mate-backend/model"
	"github.com/cglotr/lc-mate-backend/util"
)

type UserDaoImpl struct {
	db *sql.DB
}

func NewUserDaoImpl(db *sql.DB) *UserDaoImpl {
	return &UserDaoImpl{
		db: db,
	}
}

func (u *UserDaoImpl) Upsert(user *model.UserModel) error {
	result, err := u.db.Exec(
		`
		UPDATE
			user SET rating = ?,
			badge = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			username = ?
		;
		`,
		user.Rating,
		user.Badge,
		user.Username,
	)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count > 0 {
		return nil
	} else {
		_, err = u.db.Exec(
			`
			INSERT INTO user (
				username,
				rating,
				badge
			)
			VALUES (
				?,
				?,
				?
			);
			`,
			user.Username,
			user.Rating,
			user.Badge,
		)
		if err != nil {
			return err
		}
		return nil
	}
}

func (u *UserDaoImpl) QueryMostOutdatedUser() (*model.UserModel, error) {
	rows, err := u.db.Query(
		`
		SELECT
			username,
			rating,
			badge
		FROM user
		ORDER BY updated_at ASC
		LIMIT 1
		;
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.New(ErrUserNotFound())
	}
	user := model.UserModel{}
	err = rows.Scan(&user.Username, &user.Rating, &user.Badge)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDaoImpl) QueryUsers(usernames []string) ([]*model.UserModel, error) {
	query := fmt.Sprintf(
		`
		SELECT
			username,
			rating,
			badge
		FROM user
		WHERE
			username IN (%v)
		;
		`,
		util.GetUsernamesQuery(usernames),
	)
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []*model.UserModel{}
	for rows.Next() {
		user := model.UserModel{}
		err = rows.Scan(&user.Username, &user.Rating, &user.Badge)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *UserDaoImpl) MoveBackUpdatedAtOneDay(username string) error {
	_, err := u.db.Exec(
		`
		UPDATE user
		SET updated_at = NOW() - INTERVAL 1 MONTH
		WHERE username = ?
		;
		`,
		username,
	)
	return err
}
