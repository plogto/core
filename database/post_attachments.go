package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type PostAttachments struct {
	Queries *db.Queries
}

func (p *PostAttachments) CreatePostAttachment(ctx context.Context, postID, fileID uuid.UUID) (*db.PostAttachment, error) {
	postAttachment, err := p.Queries.CreatePostAttachment(ctx, db.CreatePostAttachmentParams{
		PostID: postID,
		FileID: fileID,
	})

	if err != nil {
		return nil, err
	}

	return postAttachment, nil
}
