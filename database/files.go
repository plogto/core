package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type Files struct {
	Queries *db.Queries
}

func (f *Files) CreateFile(ctx context.Context, arg db.CreateFileParams) (*db.File, error) {
	file, err := f.Queries.CreateFile(ctx, arg)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *Files) GetFilesByPostID(ctx context.Context, postID uuid.UUID) ([]*db.File, error) {
	files, err := f.Queries.GetFilesByPostID(ctx, postID)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func (f *Files) GetFilesByTicketMessageID(ctx context.Context, ticketMessageID uuid.UUID) ([]*db.File, error) {
	files, err := f.Queries.GetFilesByTicketMessageID(ctx, ticketMessageID)

	if err != nil {
		return nil, err
	}

	return files, nil

}

func (f *Files) GetFileByHash(ctx context.Context, hash string) (*db.File, error) {
	// FIXME: use dataloader
	file, err := f.Queries.GetFileByHash(ctx, hash)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *Files) GetFileByName(ctx context.Context, name string) (*db.File, error) {
	// FIXME: use dataloader
	file, err := f.Queries.GetFileByName(ctx, name)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *Files) GetFileByID(ctx context.Context, id uuid.UUID) (*db.File, error) {
	// FIXME: use dataloader
	file, err := f.Queries.GetFileByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return file, nil
}
