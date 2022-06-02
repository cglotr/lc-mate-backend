package dao

import (
	"database/sql"

	"github.com/cglotr/lc-mate-backend/model"
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
		`,
		user.Rating,
		user.Badge,
		user.Username,
	)
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
			)
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

func (u *UserDaoImpl) ReadAll() ([]*model.UserModel, error) {
	rows, err := u.db.Query(
		`
		SELECT
			username,
			rating,
			badge
		FROM user
		ORDER BY
			username ASC
		`,
	)
	if err != nil {
		return nil, err
	}
	result := []*model.UserModel{}
	for rows.Next() {
		user := model.UserModel{}
		err = rows.Scan(&user.Username, &user.Rating, &user.Badge)
		if err != nil {
			return nil, err
		}
		result = append(result, &user)
	}
	return result, nil
}

func (u *UserDaoImpl) Query(username string) (*model.UserModel, error) {
	rows, err := u.db.Query(
		`
		SELECT
			username,
			rating,
			badge
		FROM user
		WHERE
			username = ?
		`,
		username,
	)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	user := model.UserModel{}
	err = rows.Scan(&user.Username, &user.Rating, &user.Badge)
	if err != nil {
		return nil, err
	}
	return &user, nil
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
		`,
	)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	user := model.UserModel{}
	err = rows.Scan(&user.Username, &user.Rating, &user.Badge)
	if err != nil {
		return nil, err
	}
	return &user, nil
}