package service

import (
	"context"
	"easy_note/cmd/notedemo/dal/db"
	"easy_note/cmd/notedemo/kitex_gen/notedemo"
	"easy_note/cmd/notedemo/pack"
	"easy_note/cmd/notedemo/rpc"
	"easy_note/cmd/userdemo/kitex_gen/userdemo"
)

type MGetNoteService struct {
	ctx context.Context
}

func NewMGetNoteService(ctx context.Context) *MGetNoteService {
	return &MGetNoteService{ctx: ctx}
}

func (s *MGetNoteService) MGetNote(req *notedemo.MGetNoteRequest) ([]*notedemo.Note, error) {
	noteModels, err := db.MGetNotes(s.ctx, req.NoteIds)
	if err != nil {
		return nil, err
	}
	uIds := pack.UserIds(noteModels)
	userMap, err := rpc.MGetUser(s.ctx, &userdemo.MGetUserRequest{UserIds: uIds})
	if err != nil {
		return nil, err
	}
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := userMap[notes[i].UserId]; ok {
			notes[i].UserName = u.UserName
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, nil
}
