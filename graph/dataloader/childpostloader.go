package graph

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

func PrepareChildPostLoader(ctx context.Context, queries *db.Queries) ChildPostLoader {
	return ChildPostLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*model.Post, []error) {
			user, _ := middleware.GetCurrentUserFromCTX(ctx)

			var childPostIDs []pgtype.UUID
			posts, err := queries.GetChildPostsByIDsAndUserID(ctx, db.GetChildPostsByIDsAndUserIDParams{
				UserID:  user.ID,
				ChildID: convertor.StringsToUUIDs(ids),
			})

			for _, value := range posts {
				childPostIDs = append(childPostIDs, value.ID)
			}

			if err != nil {
				return nil, []error{err}
			}

			p := make(map[pgtype.UUID]*model.Post, len(posts))

			for _, post := range posts {
				p[post.ID] = convertor.DBPostToModel(post)
			}

			result := make([]*model.Post, len(ids))

			for i, id := range ids {
				result[i] = p[convertor.StringToUUID(id)]
			}

			return result, nil
		},
	}
}
