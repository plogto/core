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

		ctxUser := context.WithValue(r.Context(), constants.USER_LOADER_KEY, &userLoader)

		ctx := r.WithContext(ctxUser)
		next.ServeHTTP(w, ctx)
	})
}

func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(constants.USER_LOADER_KEY).(*UserLoader)
}

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
