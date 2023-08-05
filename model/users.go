package model

import (
	"errors"
	"fmt"
	"log"
	t "xitulu/types"
	u "xitulu/utils"

	"github.com/google/uuid"
)

// 查询所有用户数据
func SelectUsersAll() (*map[string]any, error) {
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
	return &finalResult, nil
}

// 根据用户名查询用户
func SelectUserFirst(user *t.UserLogin) (*t.UserModel, error) {
	var dbUser t.UserModel
	result := orm.Table("users").Where("username = ? AND status = 1", user.Username).First(&dbUser)
	err := result.Error
	if err != nil {
		log.Fatalln("SelectUserFirstError:", err)
		return &dbUser, err
	}
	return &dbUser, err
}

// 新增用户
func InsertUser(user *t.UserAdd) (*t.UserRes, error) {
	var userM t.UserModel
	userRes := t.UserRes{Id: 0, UserBase: t.UserBase{Username: user.Username, Nickname: user.Nickname, Email: user.Email}, UserStatus: t.UserStatus{Status: 1}}
	check := orm.Table("users").Where("username = ?", user.Username).Take(&userM)
	if check.Error == nil {
		return &userRes, errors.New("用户名已存在")
	}
	fmt.Printf("%+v\n", userM)

	createDate := u.GetMysqlNow()
	user.CreateDate = createDate
	user.Status = 1
	result := orm.Table("users").Create(&user)
	err := result.Error
	if err != nil {
		log.Fatalln("InsertUserError:", err)
		return &userRes, err
	}
	userRes.Id = user.Id
	token, errT := UpdateUserUuid(user.Id, false)
	if errT != nil {
		return &userRes, errT
	}
	userRes.Token = token
	return &userRes, nil
}

// 更新用户
func UpdateUser(user *t.UserModel) error {
	// TODO
	return nil
}

// 删除用户数据
func DeleteUser(id int) error {
	// TODO
	return nil
}

// 生成并更新用户token
func UpdateUserUuid(id int, empty bool) (string, error) {
	if !empty {
		uuids := uuid.New()
		str := uuids.String()
		// TODO 事务&token碰撞检查
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
