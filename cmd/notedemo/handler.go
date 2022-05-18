package main

import (
	"context"
	"easy_note/cmd/notedemo/kitex_gen/notedemo"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *notedemo.CreateNoteResponse) (resp *notedemo.CreateNoteRequest, err error) {
	// TODO: Your code here...
	return
}

// MGetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) MGetNote(ctx context.Context, req *notedemo.MGetNoteResponse) (resp *notedemo.MGetNoteRequest, err error) {
	// TODO: Your code here...
	return
}

// DeleteNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteNote(ctx context.Context, req *notedemo.DeleteNoteResponse) (resp *notedemo.DeleteNoteRequest, err error) {
	// TODO: Your code here...
	return
}

// QueryNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) QueryNote(ctx context.Context, req *notedemo.QueryNoteResponse) (resp *notedemo.QueryNoteRequest, err error) {
	// TODO: Your code here...
	return
}

// UpdateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req *notedemo.UpdateNoteResponse) (resp *notedemo.UpdateNoteRequest, err error) {
	// TODO: Your code here...
	return
}
