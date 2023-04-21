package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type Files struct {
	Queries *db.Queries
}

func (f *Files) CreateFile(ctx context.Context, arg db.CreateFileParams) (*db.File, error) {
	return f.Queries.CreateFile(ctx, arg)
}

func (f *Files) GetFilesByPostID(ctx context.Context, postID pgtype.UUID) ([]*db.File, error) {
	files, _ := f.Queries.GetFilesByPostID(ctx, postID)

	return files, nil
}

func (f *Files) GetFilesByTicketMessageID(ctx context.Context, ticketMessageID pgtype.UUID) ([]*db.File, error) {
	files, _ := f.Queries.GetFilesByTicketMessageID(ctx, ticketMessageID)

	return files, nil
}

func (f *Files) GetFileByHash(ctx context.Context, hash string) (*db.File, error) {
	// TODO: use dataloader
	return f.Queries.GetFileByHash(ctx, hash)
}

func (f *Files) GetFileByName(ctx context.Context, name string) (*db.File, error) {
	// TODO: use dataloader
	return f.Queries.GetFileByName(ctx, name)
}

func (f *Files) GetFileByID(ctx context.Context, id pgtype.UUID) (*db.File, error) {
	// TODO: use dataloader
	return f.Queries.GetFileByID(ctx, id)
}
