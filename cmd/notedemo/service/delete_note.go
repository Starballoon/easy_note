package service

import (
	"context"
	"easy_note/cmd/notedemo/dal/db"
	"easy_note/cmd/notedemo/kitex_gen/notedemo"
)

type DelNoteService struct {
	ctx context.Context
}

func NewDelNoteService(ctx context.Context) *DelNoteService {
	return &DelNoteService{
		ctx: ctx,
	}
}

func (s *DelNoteService) DelNote(req *notedemo.DeleteNoteRequest) error {
	return db.DeleteNote(s.ctx, req.NoteId, req.UserId)
}
