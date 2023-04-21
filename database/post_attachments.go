package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type PostAttachments struct {
	Queries *db.Queries
}

func (p *PostAttachments) CreatePostAttachment(ctx context.Context, postID, fileID pgtype.UUID) (*db.PostAttachment, error) {
	postAttachment, _ := p.Queries.CreatePostAttachment(ctx, db.CreatePostAttachmentParams{
		PostID: postID,
		FileID: fileID,
	})

	return postAttachment, nil
}
