package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/plog-core/graph/generated"
	"github.com/favecode/plog-core/graph/model"
)

func (r *mutationResolver) AddPostComment(ctx context.Context, input model.CommentPostInput) (*model.PostComment, error) {
	return r.Service.AddPostComment(ctx, input)
}

func (r *postCommentResolver) Parent(ctx context.Context, obj *model.PostComment) (*model.PostComment, error) {
	return r.Service.GetPostCommentByID(ctx, obj.ParentID)
}

func (r *postCommentResolver) Children(ctx context.Context, obj *model.PostComment) (*model.PostComments, error) {
	return r.Service.GetChildrenPostComments(ctx, obj.PostID, obj.ID)
}

func (r *postCommentResolver) User(ctx context.Context, obj *model.PostComment) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

func (r *postCommentResolver) Post(ctx context.Context, obj *model.PostComment) (*model.Post, error) {
	return r.Service.GetPostsByID(ctx, obj.PostID)
}

func (r *queryResolver) GetPostComments(ctx context.Context, postID string, input *model.PaginationInput) (*model.PostComments, error) {
	return r.Service.GetPostComments(ctx, postID)
}

// PostComment returns generated.PostCommentResolver implementation.
func (r *Resolver) PostComment() generated.PostCommentResolver { return &postCommentResolver{r} }

type postCommentResolver struct{ *Resolver }
