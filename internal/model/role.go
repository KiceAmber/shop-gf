package model

import "github.com/gogf/gf/v2/os/gtime"

type RoleCreateUpdateBase struct {
	Name string
	Desc string
}

// 创建内容
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// 返回的结果
type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}

type RoleUpdateInput struct {
	RoleCreateUpdateBase
	Id uint
}

type RoleGetListInput struct {
	Page int
	Size int
	Sort int
}

type RoleGetListOutput struct {
	List  []RoleGetListOutputItem `json:"list" description:"返回列表"`
	Page  int                     `json:"list" description:"返回列表"`
	Size  int                     `json:"list" description:"返回列表"`
	Total int                     `json:"list" description:"返回列表"`
}

type RoleGetListOutputItem struct {
	Id       uint        `json:"id"`
	Name     string      `json:"name"`
	RoleIds  string      `json:"role_ids"`
	IsRole   uint        `json:"is_role"`
	CreateAt *gtime.Time `json:"create_at"`
	UpdateAt *gtime.Time `json:"update_at"`
}
