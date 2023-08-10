package types

type UserBase struct {
	Username string `json:"username" gorm:"column:username"`
	Nickname string `json:"nickname,omitempty" gorm:"column:nickname"`
	Email    string `json:"email,omitempty" gorm:"column:email"`
}

type UserPassword struct {
	Password string `json:"password" gorm:"column:password"`
}

type UserStatus struct {
	CreateDate string `json:"createDate" gorm:"column:create_date"`
	Status     int    `json:"status" gorm:"column:status"`
}

type UserRes struct {
	Id int `json:"id" gorm:"column:id;primary"`
	UserBase
	UserStatus
	Token string `json:"token,omitempty" gorm:"column:token"`
}

type UserModel struct {
	UserRes
	UserPassword
}

type UserAdd struct {
	Id int `json:"id,omitempty" gorm:"column:id;primary"`
	UserBase
	UserStatus
	UserPassword
}

type UserLogin struct {
	UserBase
	UserPassword
}
