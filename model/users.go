package model

import (
	"errors"
	"log"
	t "xitulu/types"
	u "xitulu/util"

	"github.com/google/uuid"
)

func SelectUsersAll() (map[string]any, error) {
	var users []t.UserRes
	var count int64
	result := orm.Table("users").Find(&users).Count(&count)
	err := result.Error

	if err != nil {
		log.Fatalln("SelectUsersAllError:", err)
		return nil, err
	}
	finalResult := map[string]any{
		"list":  users,
		"count": count,
	}
	return finalResult, nil
}

func SelectUserFirst(user t.UserLogin) (t.UserModel, error) {
	var dbUser t.UserModel
	result := orm.Table("users").Where("username = ? AND status = 1", user.Username).First(&dbUser)
	err := result.Error
	if err != nil {
		log.Fatalln("SelectUserFirstError:", err)
		return dbUser, err
	}
	return dbUser, err
}

func InsertUser(user t.UserAdd) error {
	createDate := u.GetMysqlNow()
	user.CreateDate = createDate
	user.Status = 1
	result := orm.Table("users").Create(user)
	err := result.Error
	if err != nil {
		log.Fatalln("InsertUserError:", err)
		return err
	}
	return nil
}

func UpdateUserUuid(id int, empty bool) (string, error) {
	if !empty {
		uuids := uuid.New()
		str := uuids.String()
		orm.Table("users").Where("id = ?", id).Update("token", str)
		return str, nil
	} else {
		result := orm.Table("users").Where("id = ?", id).Update("token", nil)
		if result.RowsAffected == 0 {
			return "", errors.New("用户不存在")
		} else {
			return "", nil
		}
	}
}
