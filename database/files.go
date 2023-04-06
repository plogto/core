package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/util"
)

type Files struct {
	Queries *db.Queries
}

func (f *Files) CreateFile(ctx context.Context, arg db.CreateFileParams) (*db.File, error) {
	return util.HandleDBResponse(f.Queries.CreateFile(ctx, arg))
}

func (f *Files) GetFilesByPostID(ctx context.Context, postID uuid.UUID) ([]*db.File, error) {
	files, _ := f.Queries.GetFilesByPostID(ctx, postID)

	return files, nil
}

func (f *Files) GetFilesByTicketMessageID(ctx context.Context, ticketMessageID uuid.UUID) ([]*db.File, error) {
	files, _ := f.Queries.GetFilesByTicketMessageID(ctx, ticketMessageID)

	return files, nil
}

func (f *Files) GetFileByHash(ctx context.Context, hash string) (*db.File, error) {
	// TODO: use dataloader
	return util.HandleDBResponse(f.Queries.GetFileByHash(ctx, hash))
}

func (f *Files) GetFileByName(ctx context.Context, name string) (*db.File, error) {
	// TODO: use dataloader
	return util.HandleDBResponse(f.Queries.GetFileByName(ctx, name))
}

func (f *Files) GetFileByID(ctx context.Context, id uuid.UUID) (*db.File, error) {
	// TODO: use dataloader
	return util.HandleDBResponse(f.Queries.GetFileByID(ctx, id))
}
