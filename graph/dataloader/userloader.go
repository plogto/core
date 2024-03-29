package graph

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
)

func PrepareUserLoader(ctx context.Context, queries *db.Queries) UserLoader {
	return UserLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*db.User, []error) {
			users, err := queries.GetUsersByIDs(ctx, convertor.StringsToUUIDs(ids))

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[pgtype.UUID]*db.User, len(users))

			for _, user := range users {
				u[user.ID] = user
			}

			result := make([]*db.User, len(ids))

			for i, id := range ids {
				result[i] = u[convertor.StringToUUID(id)]
			}

			return result, nil
		},
	}
}
