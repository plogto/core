package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type PostMentions struct {
	Queries *db.Queries
}

func (p *PostMentions) CreatePostMention(ctx context.Context, userID, postID uuid.UUID) (*db.PostMention, error) {
	postMention, err := p.Queries.CreatePostMention(ctx, db.CreatePostMentionParams{
		UserID: userID,
		PostID: postID,
	})

	if err != nil {
		return nil, err
	}

	return postMention, nil
}

func (p *PostMentions) DeletePostMention(ctx context.Context, userID, postID uuid.UUID) ([]*db.PostMention, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	postMentions, err := p.Queries.DeletePostMention(ctx, db.DeletePostMentionParams{
		PostID:    postID,
		UserID:    userID,
		DeletedAt: DeletedAt,
	})

	if err != nil {
		return nil, err
	}

	return postMentions, nil

}

func (p *PostMentions) DeletePostMentionsByPostID(ctx context.Context, postID uuid.UUID) ([]*db.PostMention, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	postMentions, err := p.Queries.DeletePostMentionsByPostID(ctx, db.DeletePostMentionsByPostIDParams{
		PostID:    postID,
		DeletedAt: DeletedAt,
	})

	if err != nil {
		return nil, err
	}

	return postMentions, nil
}
