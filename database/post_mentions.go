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
	postMention, _ := p.Queries.CreatePostMention(ctx, db.CreatePostMentionParams{
		UserID: userID,
		PostID: postID,
	})

	return postMention, nil
}

func (p *PostMentions) DeletePostMention(ctx context.Context, userID, postID uuid.UUID) ([]*db.PostMention, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	postMentions, _ := p.Queries.DeletePostMention(ctx, db.DeletePostMentionParams{
		PostID:    postID,
		UserID:    userID,
		DeletedAt: DeletedAt,
	})

	return postMentions, nil

}

func (p *PostMentions) DeletePostMentionsByPostID(ctx context.Context, postID uuid.UUID) ([]*db.PostMention, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	postMentions, _ := p.Queries.DeletePostMentionsByPostID(ctx, db.DeletePostMentionsByPostIDParams{
		PostID:    postID,
		DeletedAt: DeletedAt,
	})

	return postMentions, nil
}
