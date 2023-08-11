package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	t "xitulu/types"
	"xitulu/utils"
)

type User struct {
	Id         int    `json:"id,omitempty" gorm:"column:id;primary"`
	Username   string `json:"username" validate:"alphanum,min=5,max=20" gorm:"column:username"`
	Password   string `json:"password" validate:"sha256" gorm:"column:password"`
	Nickname   string `json:"nickname,omitempty" validate:"min=2,max=20" gorm:"column:nickname;default:''"`
	AvatarUrl  string `json:"avatarUrl,omitempty" validate:"url" gorm:"column:avatar_url;default:''"`
	Email      string `json:"email,omitempty" validate:"email" gorm:"column:email"`
	Status     int    `json:"status" gorm:"column:status"`
	CreateDate string `json:"createDate" gorm:"column:create_date"`
	Token      string `json:"token,omitempty" gorm:"column:token;default:null"`
}

// SelectAll 查询所有用户数据
func (u *User) SelectAll() (*map[string]any, error) {
	var users []t.UserRes
	var count int64
	if err := db.Table("users").Where("status = 1").Find(&users).Count(&count).Error; err != nil {
		log.Fatalln("SelectUsersAllError:", err)
		return nil, err
	} else {
		finalResult := map[string]any{
			"list":  users,
			"count": count,
		}
		return &finalResult, nil
	}
}

// SelectOne 根据用户名查询用户
func (u *User) SelectOne(user *User) (*User, error) {
	var data User
	result := db.Table("users").Where("username = ? AND status = 1", user.Username).Take(&data)
	err := result.Error
	if err != nil {
		log.Fatalln("SelectUserFirstError:", err)
		return nil, err
	}
	return &data, err
}

// Insert 新增用户
func (u *User) Insert(user *User) (*t.UserRes, error) {
	var userM User
	userRes := t.UserRes{Id: 0, UserBase: t.UserBase{Username: user.Username, Nickname: user.Nickname, Email: user.Email}, UserStatus: t.UserStatus{Status: 1}}
	check := db.Table("users").Where("username = ?", user.Username).Take(&userM)
	if check.Error == nil {
		return &userRes, errors.New("用户名已存在")
	}
	fmt.Printf("%+v\n", userM)

	createDate := utils.GetMysqlNow()
	user.CreateDate = createDate
	user.Status = 1
	result := db.Table("users").Create(&user)
	err := result.Error
	if err != nil {
		log.Fatalln("InsertUserError:", err)
		return &userRes, err
	}
	userRes.Id = user.Id
	token, errT := u.UpdateUserUuid(user.Id, false)
	if errT != nil {
		return &userRes, errT
	}
	userRes.Token = token
	return &userRes, nil
}

// UpdateUserUuid 生成并更新用户token
func (u *User) UpdateUserUuid(id int, empty bool) (string, error) {
	if !empty {
		uuids := uuid.New()
		str := uuids.String()
		// TODO 事务&token碰撞检查
		db.Table("users").Where("id = ?", id).Update("token", str)
		return str, nil
	} else {
		result := db.Table("users").Where("id = ?", id).Update("token", nil)
		if result.RowsAffected == 0 {
			return "", errors.New("用户不存在")
		} else {
			return "", nil
		}
	}
}

//func (u *User) CheckUuid(uuid) error {
//
//}

// Update 更新用户数据
func (u *User) Update(data *User) error {
	result := db.
		Table("users").
		Model(&data).
		Where("id = ?", data.Id).
		Omit("id", "username", "create_date").
		Updates(&data)
	if result.RowsAffected == 0 {
		return errors.New("ID错误")
	}
	return result.Error
}

// Delete 软删除用户
func (u *User) Delete(id int) error {
	result := db.Table("users").
		Where("id = ?", id).
		Update("status", 0)
	if result.RowsAffected == 0 {
		return errors.New("ID错误")
	}
	return result.Error
}
