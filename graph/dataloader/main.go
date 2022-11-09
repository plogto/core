package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/graph/model"
)

func DataloaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := PrepareUserLoader(db)
		postLoader := PreparePostLoader(db)
		tagLoader := PrepareTagLoader(db)

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
func PrepareUserLoader(db *pg.DB) UserLoader {
	return UserLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*model.User, []error) {
			var users []*model.User

			err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[string]*model.User, len(users))

			for _, user := range users {
				u[user.ID] = user
			}

			result := make([]*model.User, len(ids))

			for i, id := range ids {
				result[i] = u[id]
			}

			return result, nil
		},
	}
}

func PreparePostLoader(db *pg.DB) PostLoader {
	return PostLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*model.Post, []error) {
			var posts []*model.Post

			err := db.Model(&posts).Where("id in (?)", pg.In(ids)).Select()

			if err != nil {
				return nil, []error{err}
			}

			p := make(map[string]*model.Post, len(posts))

			for _, post := range posts {
				p[post.ID] = post
			}

			result := make([]*model.Post, len(ids))

			for i, id := range ids {
				result[i] = p[id]
			}

			return result, nil
		},
	}
}

func PrepareTagLoader(db *pg.DB) TagLoader {
	return TagLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*model.Tag, []error) {
			var tags []*model.Tag

			err := db.Model(&tags).Where("id in (?)", pg.In(ids)).Select()

			if err != nil {
				return nil, []error{err}
			}

			t := make(map[string]*model.Tag, len(tags))

			for _, post := range tags {
				t[post.ID] = post
			}

			result := make([]*model.Tag, len(ids))

			for i, id := range ids {
				result[i] = t[id]
			}

			return result, nil
		},
	}
}
