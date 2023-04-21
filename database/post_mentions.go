package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type PostMentions struct {
	Queries *db.Queries
}

func (p *PostMentions) CreatePostMention(ctx context.Context, userID, postID pgtype.UUID) (*db.PostMention, error) {
	postMention, _ := p.Queries.CreatePostMention(ctx, db.CreatePostMentionParams{
		UserID: userID,
		PostID: postID,
	})

	return postMention, nil
}

func (p *PostMentions) DeletePostMention(ctx context.Context, userID, postID pgtype.UUID) ([]*db.PostMention, error) {
	DeletedAt := time.Now()

	postMentions, _ := p.Queries.DeletePostMention(ctx, db.DeletePostMentionParams{
		PostID:    postID,
		UserID:    userID,
		DeletedAt: &DeletedAt,
	})

	return postMentions, nil

}

func (p *PostMentions) DeletePostMentionsByPostID(ctx context.Context, postID pgtype.UUID) ([]*db.PostMention, error) {
	DeletedAt := time.Now()

	postMentions, _ := p.Queries.DeletePostMentionsByPostID(ctx, db.DeletePostMentionsByPostIDParams{
		PostID:    postID,
		DeletedAt: &DeletedAt,
	})

	return postMentions, nil
}
