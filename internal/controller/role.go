package controller

import (
	"context"
	"rime-shop-gf/api/backend"
	"rime-shop-gf/internal/model"
	"rime-shop-gf/internal/service"
)

var Role = cRole{}

type cRole struct {
}

func (c *cRole) Create(ctx context.Context, req *backend.RoleReq) (res *backend.RoleRes, err error) {
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name:        req.Name,
			Description: req.Description,
		},
	})

	if err != nil {
		return nil, err
	}
	return &backend.RoleRes{
		RoleId: out.RoleId,
	}, nil
}
