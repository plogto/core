package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/plog-core/graph/generated"
	"github.com/favecode/plog-core/graph/model"
)

func (r *mutationResolver) LikePost(ctx context.Context, postID string) (*model.PostLike, error) {
	return r.Service.LikePost(ctx, postID)
}

func (r *mutationResolver) UnlikePost(ctx context.Context, postID string) (*model.PostLike, error) {
	return r.Service.UnlikePost(ctx, postID)
}

func (r *postLikeResolver) User(ctx context.Context, obj *model.PostLike) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

func (r *postLikeResolver) Post(ctx context.Context, obj *model.PostLike) (*model.Post, error) {
	return r.Service.GetPostsByID(ctx, obj.PostID)
}

func (r *queryResolver) GetPostLikesByPostID(ctx context.Context, postID string, input *model.PaginationInput) (*model.PostLikes, error) {
	return r.Service.GetPostLikesByPostId(ctx, postID)
}

// PostLike returns generated.PostLikeResolver implementation.
func (r *Resolver) PostLike() generated.PostLikeResolver { return &postLikeResolver{r} }

type postLikeResolver struct{ *Resolver }
