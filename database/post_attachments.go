package database

import (
	"context"

	"github.com/plogto/core/db"
)

type PostAttachments struct {
	Queries *db.Queries
}

func (p *PostAttachments) CreatePostAttachment(ctx context.Context, arg db.CreatePostAttachmentParams) (*db.PostAttachment, error) {
	postAttachment, err := p.Queries.CreatePostAttachment(ctx, arg)

	if err != nil {
		return nil, err
	}

	return postAttachment, nil
}
