package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

func DataloaderMiddleware(queries *db.Queries, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := PrepareUserLoader(r.Context(), queries)
		postLoader := PreparePostLoader(r.Context(), queries)
		tagLoader := PrepareTagLoader(r.Context(), queries)

		ctxUser := context.WithValue(r.Context(), constants.USER_LOADER_KEY, &userLoader)
		ctxPost := context.WithValue(ctxUser, constants.POST_LOADER_KEY, &postLoader)
		ctxTag := context.WithValue(ctxPost, constants.TAG_LOADER_KEY, &tagLoader)

		ctx := r.WithContext(ctxTag)
		next.ServeHTTP(w, ctx)
	})
}

// get functions
func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(constants.USER_LOADER_KEY).(*UserLoader)
}

func GetPostLoader(ctx context.Context) *PostLoader {
	return ctx.Value(constants.POST_LOADER_KEY).(*PostLoader)
}

func GetTagLoader(ctx context.Context) *TagLoader {
	return ctx.Value(constants.TAG_LOADER_KEY).(*TagLoader)
}

// prepare functions
func PrepareUserLoader(ctx context.Context, queries *db.Queries) UserLoader {
	return UserLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*db.User, []error) {

			users, err := queries.GetUsersByIDs(ctx, convertor.StringsToUUIDs(ids))

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[uuid.UUID]*db.User, len(users))

			for _, user := range users {
				u[user.ID] = user
			}

			result := make([]*db.User, len(ids))

			for i, id := range ids {
				result[i] = u[uuid.MustParse(id)]
			}

			return result, nil
		},
	}
}

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

func PrepareTagLoader(ctx context.Context, queries *db.Queries) TagLoader {
	return TagLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*model.Tag, []error) {

			tags, err := queries.GetTagByIDs(ctx, convertor.StringsToUUIDs(ids))

			if err != nil {
				return nil, []error{err}
			}

			t := make(map[uuid.UUID]*model.Tag, len(tags))

			for _, tag := range tags {
				t[tag.ID] = convertor.DBTagToModel(tag)
			}

			result := make([]*model.Tag, len(ids))

			for i, id := range ids {
				result[i] = t[uuid.MustParse(id)]
			}

			return result, nil
		},
	}
}
