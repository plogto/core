package graph

import (
	"context"
	"net/http"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/db"
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
