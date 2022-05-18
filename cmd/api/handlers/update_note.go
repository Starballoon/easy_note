package handlers

import (
	"context"
	"easy_note/cmd/api/rpc"
	"easy_note/cmd/notedemo/kitex_gen/notedemo"
	"easy_note/pkg/constants"
	"easy_note/pkg/errno"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateNote(c *gin.Context) {
	var noteVar NoteParam
	if err := c.ShouldBind(&noteVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	noteIDStr := c.Param(constants.NoteID)
	noteID, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if noteID <= 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &notedemo.UpdateNoteRequest{NoteId: noteID, UserId: userID}
	if len(noteVar.Title) != 0 {
		req.Title = noteVar.Title
	}
	if len(noteVar.Content) != 0 {
		req.Content = noteVar.Content
	}
	if err = rpc.UpdateNote(context.Background(), req); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
