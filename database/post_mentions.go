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

func (p *PostMentions) CreatePostMention(ctx context.Context, userID, postID string) (*db.PostMention, error) {
	UserID, _ := uuid.Parse(userID)
	PostID, _ := uuid.Parse(postID)

	postMention, err := p.Queries.CreatePostMention(ctx, db.CreatePostMentionParams{
		UserID: UserID,
		PostID: PostID,
	})

	if err != nil {
		return nil, err
	}

	return postMention, nil
}

func (p *PostMentions) DeletePostMention(ctx context.Context, userID, postID string) ([]*db.PostMention, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	UserID, _ := uuid.Parse(userID)
	PostID, _ := uuid.Parse(postID)
	postMentions, err := p.Queries.DeletePostMention(ctx, db.DeletePostMentionParams{
		PostID:    PostID,
		UserID:    UserID,
		DeletedAt: DeletedAt,
	})

	if err != nil {
		return nil, err
	}

	return postMentions, nil

}

func (p *PostMentions) DeletePostMentionsByPostID(ctx context.Context, postID string) ([]*db.PostMention, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	PostID, _ := uuid.Parse(postID)
	postMentions, err := p.Queries.DeletePostMentionsByPostID(ctx, db.DeletePostMentionsByPostIDParams{
		PostID:    PostID,
		DeletedAt: DeletedAt,
	})

	if err != nil {
		return nil, err
	}

	return postMentions, nil
}
