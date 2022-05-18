package service

import (
	"context"
	"easy_note/cmd/userdemo/dal/db"
	"easy_note/cmd/userdemo/kitex_gen/userdemo"
	"easy_note/cmd/userdemo/pack"
)

type MGetUserService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

func (s *MGetUserService) MGetUser(req *userdemo.MGetUserRequest) ([]*userdemo.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
