package service

import (
	"context"
	"easy_note/cmd/notedemo/dal/db"
	"easy_note/cmd/notedemo/kitex_gen/notedemo"
	"easy_note/cmd/notedemo/pack"
	"easy_note/cmd/notedemo/rpc"
	"easy_note/cmd/userdemo/kitex_gen/userdemo"
)

type QueryNoteService struct {
	ctx context.Context
}

func NewQueryNoteService(ctx context.Context) *QueryNoteService {
	return &QueryNoteService{ctx: ctx}
}

func (s *QueryNoteService) QueryNoteService(req *notedemo.QueryNoteRequest) ([]*notedemo.Note, int64, error) {
	noteModels, total, err := db.QueryNote(s.ctx, req.UserId, &req.SearchKey, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, 0, err
	}
	userMap, err := rpc.MGetUser(s.ctx, &userdemo.MGetUserRequest{UserIds: []int64{req.UserId}})
	if err != nil {
		return nil, 0, err
	}
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := userMap[notes[i].UserId]; ok {
			notes[i].UserName = u.UserName
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, total, nil
}
