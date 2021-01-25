package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type Dao struct{}

func New() *Dao {
	return &Dao{}
}

func (d *Dao) FindByID(id int) (user *model.User, err error) {

	err = Db.Table("t_user").Where("id = ?", userID).Find(user).Error
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrRecordNotFound
	}
	// ...... 这里可以补充其他错误处理
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("find by user_id: %v error", userID))
	}
	return
}
