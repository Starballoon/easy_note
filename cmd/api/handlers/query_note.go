package handlers

import (
	"context"
	"easy_note/cmd/api/rpc"
	"easy_note/cmd/notedemo/kitex_gen/notedemo"
	"easy_note/pkg/constants"
	"easy_note/pkg/errno"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func QueryNote(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	var queryVar struct {
		Limit         int64  `json:"limit" form:"limit"`
		Offset        int64  `json:"offset" form:"offset"`
		SearchKeyword string `json:"search_keyword" form:"search_keyword"`
	}
	if err := c.BindQuery(&queryVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}

	if queryVar.Limit < 0 || queryVar.Offset < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &notedemo.QueryNoteRequest{UserId: userID, Offset: queryVar.Offset, Limit: queryVar.Limit}
	if len(queryVar.SearchKeyword) != 0 {
		req.SearchKey = queryVar.SearchKeyword
	}
	notes, total, err := rpc.QueryNotes(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{constants.Total: total, constants.Notes: notes})
}
