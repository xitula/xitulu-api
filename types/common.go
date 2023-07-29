/*
@Title
@Description 通用的自定义数据类型
@Auther xitu
*/
package types

// 通用请求参数结构体
type ReqParams map[string]string

// API接口标准返回对象结构体
type Res struct {
	Code    int         `json:"code"`           // 错误码
	Message string      `json:"message"`        // 错误信息
	Data    interface{} `json:"data,omitempty"` // 返回数据，可选
}

// 代办事项
type Todo struct {
	Id             int    `json:"id" gorm:"column:id;primarykey"`                          // ID
	Uid            int    `json:"uid" gorm:"column:uid"`                                   // 用户ID
	Content        string `json:"content" gorm:"column:content"`                           // 内容
	Description    string `json:"description,omitempty" gorm:"column:description"`         // 描述，可选
	CreateDate     string `json:"createDate,omitempty" gorm:"column:create_date"`          // 创建日期
	Done           int    `json:"done,omitempty" gorm:"column:done"`                       // 是否已完成
	LastUpdateDate string `json:"lastUpdateDate,omitempty" gorm:"column:last_update_date"` // 最后更新日期，可选
}

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
	Id int `json:"id" gorm:"column:id;primarykey"`
	UserBase
	UserStatus
	Token string `json:"token,omitempty" gorm:"column:token"`
}

type UserModel struct {
	UserRes
	UserPassword
}

type UserAdd struct {
	Id int `json:"id,omitempty" gorm:"column:id;primarykey"`
	UserBase
	UserStatus
	UserPassword
}

type UserLogin struct {
	UserBase
	UserPassword
}
