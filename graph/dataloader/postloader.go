package graph

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

func PreparePostLoader(ctx context.Context, queries *db.Queries) PostLoader {
	return PostLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*model.Post, []error) {
			var postIDs []pgtype.UUID
			posts, err := queries.GetPostsByIDs(ctx, convertor.StringsToUUIDs(ids))

			for _, value := range posts {
				postIDs = append(postIDs, value.ID)
			}

			if err != nil {
				return nil, []error{err}
			}

			p := make(map[pgtype.UUID]*model.Post, len(posts))

			for _, post := range posts {
				p[post.ID] = convertor.DBPostToModel(post)
			}

			files, _ := queries.GetFilesByPostIDs(ctx, postIDs)

			for _, value := range p {
				var postAttachments []*db.File
				for _, fileValue := range files {
					if fileValue.PostID == value.ID {
						postAttachments = append(postAttachments, &db.File{
							ID:     fileValue.ID,
							Name:   fileValue.Name,
							Height: fileValue.Height,
							Width:  fileValue.Width,
						})
					}
				}
				p[value.ID].Attachment = postAttachments
			}

			result := make([]*model.Post, len(ids))

			for i, id := range ids {
				result[i] = p[convertor.StringToUUID(id)]
			}

			return result, nil
		},
	}
}
