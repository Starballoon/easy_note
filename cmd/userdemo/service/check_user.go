package service

import (
	"context"
	"crypto/md5"
	"easy_note/cmd/userdemo/dal/db"
	"easy_note/cmd/userdemo/kitex_gen/userdemo"
	"easy_note/pkg/errno"
	"fmt"
	"io"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (s *CheckUserService) CheckUser(req *userdemo.CheckUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.UserName
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(u.ID), nil
}
