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
	Id             int    `json:"id" gorm:"column:id;primarykey"`                  // ID
	Contant        string `json:"contant" gorm:"column:content"`                   // 内容
	Description    string `json:"description,omitempty" gorm:"column:description"` // 描述，可选
	CreateDate     string `json:"createDate" gorm:"column:create_date"`            // 创建日期
	Done           int    `json:"done" gorm:"column:done"`                         // 是否已完成
	LastUpdateDate string `json:"lastUpdateDate" gorm:"column:last_update_date"`   // 最后更新日期，可选
}
