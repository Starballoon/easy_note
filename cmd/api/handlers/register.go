package handlers

import (
	"context"
	"easy_note/cmd/api/rpc"
	"easy_note/cmd/userdemo/kitex_gen/userdemo"
	"easy_note/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerVar UserParam
	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord},
	)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
