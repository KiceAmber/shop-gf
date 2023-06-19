package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/backend/role/add" method:"post" desc:"添加角色" tags:"role"`
	Name   string `json:"name" v:"required#名称必须不为空" dc:"角色名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleRes struct {
	RoleId int `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" method:"post" desc:"修改角色" tags:"role"`
	Id     uint   `json:"id" v:"required#角色ID必须不为空" desc:"角色ID"`
	Name   string `json:"name" v:"required#名称必须不为空" dc:"角色名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleUpdateRes struct {
	Id uint `json:"id"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"delete" desc:"删除角色" tags:"role"`
	Id     uint `json:"id" v:"min:1#请选择要删除的角色" desc:"角色ID"`
}

type RoleDeleteRes struct {
}

type RoleGetListCommonReq struct {
	g.Meta `path:"/backend/role/list" method:"get" tags:"role"`
	CommonPaginationReq
}

type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"页码"`
	Size  int         `json:"size" description:"每页包含的数量"`
	Total int         `json:"total" description:"总共数量"`
}
