package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

func (r *mutationResolver) AddPost(ctx context.Context, input model.AddPostInput) (*model.Post, error) {
	return r.Service.AddPost(ctx, input)
}

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

func (r *postResolver) Likes(ctx context.Context, obj *model.Post) (*model.PostLikes, error) {
	return r.Service.GetPostLikesByPostId(ctx, obj.ID)
}

func (r *postResolver) Comments(ctx context.Context, obj *model.Post) (*model.Comments, error) {
	return r.Service.GetComments(ctx, obj.ID)
}

func (r *postResolver) IsLiked(ctx context.Context, obj *model.Post) (*model.PostLike, error) {
	return r.Service.IsPostLiked(ctx, obj.ID)
}

func (r *postResolver) IsSaved(ctx context.Context, obj *model.Post) (*model.PostSave, error) {
	return r.Service.IsPostSaved(ctx, obj.ID)
}

func (r *queryResolver) GetPostsByUsername(ctx context.Context, username string, input *model.PaginationInput) (*model.Posts, error) {
	return r.Service.GetPostsByUsername(ctx, username, input)
}

func (r *queryResolver) GetPostsByTagName(ctx context.Context, tagName string, input *model.PaginationInput) (*model.Posts, error) {
	return r.Service.GetPostsByTagName(ctx, tagName, input)
}

func (r *queryResolver) GetPostByURL(ctx context.Context, url string) (*model.Post, error) {
	return r.Service.GetPostByURL(ctx, url)
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
