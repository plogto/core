package graph

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

func PrepareTagLoader(ctx context.Context, queries *db.Queries) TagLoader {
	return TagLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*model.Tag, []error) {
			tags, err := queries.GetTagByIDs(ctx, convertor.StringsToUUIDs(ids))

			if err != nil {
				return nil, []error{err}
			}

			t := make(map[pgtype.UUID]*model.Tag, len(tags))

			for _, tag := range tags {
				t[tag.ID] = convertor.DBTagToModel(tag)
			}

			result := make([]*model.Tag, len(ids))

			for i, id := range ids {
				result[i] = t[convertor.StringToUUID(id)]
			}

			return result, nil
		},
	}
}
