package model

type RoleCreateUpdateBase struct {
	Name        string
	Description string
}

// 创建内容
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// 返回的结果
type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}
