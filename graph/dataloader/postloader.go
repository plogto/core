package graph

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
)

func PreparePostLoader(ctx context.Context, queries *db.Queries) PostLoader {
	return PostLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*db.Post, []error) {
			posts, err := queries.GetPostsByIDs(ctx, convertor.StringsToUUIDs(ids))

			if err != nil {
				return nil, []error{err}
			}

			p := make(map[uuid.UUID]*db.Post, len(posts))

			for _, post := range posts {
				p[post.ID] = post
			}

			result := make([]*db.Post, len(ids))

			for i, id := range ids {
				result[i] = p[uuid.MustParse(id)]
			}

			return result, nil
		},
	}
}
