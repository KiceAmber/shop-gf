package controller

import (
	"context"
	"rime-shop-gf/api/backend"
	"rime-shop-gf/api/frontend"
	"rime-shop-gf/internal/model"
	"rime-shop-gf/internal/service"
)

var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *backend.RotationReq) (res *backend.RotationRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationRes{RotationId: out.RotationId}, nil
}

func (a *cRotation) Delete(ctx context.Context, req *backend.RotationDeleteReq) (res *backend.RotationDeleteRes, err error) {
	err = service.Rotation().Delete(ctx, req.Id)
	return
}

func (a *cRotation) Update(ctx context.Context, req *backend.RotationUpdateReq) (res *backend.RotationUpdateRes, err error) {
	err = service.Rotation().Update(ctx, model.RotationUpdateInput{
		Id: req.Id,
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	return &backend.RotationUpdateRes{Id: req.Id}, nil
}

func (a *cRotation) List(ctx context.Context, req *backend.RotationGetListCommonReq) (res *backend.RotationGetListCommonRes, err error) {
	getListRes, err := service.Rotation().GetList(ctx, model.RotationGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &backend.RotationGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// ListFrontend 前台轮播图查询列表
func (a *cRotation) ListFrontend(ctx context.Context, req *frontend.RotationGetListReq) (res *frontend.RotationGetListRes, err error) {
	getListRes, err := service.Rotation().GetList(ctx, model.RotationGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &frontend.RotationGetListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
