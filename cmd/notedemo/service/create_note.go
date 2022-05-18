package service

import (
	"context"
	"easy_note/cmd/notedemo/dal/db"
	"easy_note/cmd/notedemo/kitex_gen/notedemo"
)

type CreateNoteService struct {
	ctx context.Context
}

func NewCreateNoteService(ctx context.Context) *CreateNoteService {
	return &CreateNoteService{ctx: ctx}
}

func (s *CreateNoteService) CreateNote(req *notedemo.CreateNoteRequest) error {
	noteModel := &db.Note{
		UserID:  req.UserId,
		Title:   req.Title,
		Content: req.Title,
	}
	return db.CreateNote(s.ctx, []*db.Note{noteModel})
}
