package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

// AddPost is the resolver for the addPost field.
func (r *mutationResolver) AddPost(ctx context.Context, input model.AddPostInput) (*model.Post, error) {
	return r.Service.AddPost(ctx, input)
}

// EditPost is the resolver for the editPost field.
func (r *mutationResolver) EditPost(ctx context.Context, postID string, input model.EditPostInput) (*model.Post, error) {
	return r.Service.EditPost(ctx, postID, input)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, postID string) (*model.Post, error) {
	return r.Service.DeletePost(ctx, postID)
}

// Parent is the resolver for the parent field.
func (r *postResolver) Parent(ctx context.Context, obj *model.Post) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, obj.ParentID)
}

// Child is the resolver for the child field.
func (r *postResolver) Child(ctx context.Context, obj *model.Post) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// Attachment is the resolver for the attachment field.
func (r *postResolver) Attachment(ctx context.Context, obj *model.Post) ([]*model.File, error) {
	return r.Service.GetPostAttachmentsByPostID(ctx, obj.ID)
}

// Likes is the resolver for the likes field.
func (r *postResolver) Likes(ctx context.Context, obj *model.Post) (*model.LikedPosts, error) {
	return r.Service.GetLikedPostsByPostID(ctx, obj.ID)
}

// Replies is the resolver for the replies field.
func (r *postResolver) Replies(ctx context.Context, obj *model.Post) (*model.Posts, error) {
	return r.Service.GetPostsByParentID(ctx, obj.ID)
}

// IsLiked is the resolver for the isLiked field.
func (r *postResolver) IsLiked(ctx context.Context, obj *model.Post) (*model.LikedPost, error) {
	return r.Service.IsPostLiked(ctx, obj.ID)
}

// IsSaved is the resolver for the isSaved field.
func (r *postResolver) IsSaved(ctx context.Context, obj *model.Post) (*model.SavedPost, error) {
	return r.Service.IsPostSaved(ctx, obj.ID)
}

// GetPostsByUsername is the resolver for the getPostsByUsername field.
func (r *queryResolver) GetPostsByUsername(ctx context.Context, username string, input *model.PaginationInput) (*model.Posts, error) {
	return r.Service.GetPostsByUsername(ctx, username, input)
}

// GetPostsByTagName is the resolver for the getPostsByTagName field.
func (r *queryResolver) GetPostsByTagName(ctx context.Context, tagName string, input *model.PaginationInput) (*model.Posts, error) {
	return r.Service.GetPostsByTagName(ctx, tagName, input)
}

// GetPostByURL is the resolver for the getPostByUrl field.
func (r *queryResolver) GetPostByURL(ctx context.Context, url string) (*model.Post, error) {
	return r.Service.GetPostByURL(ctx, url)
}

// GetTimelinePosts is the resolver for the getTimelinePosts field.
func (r *queryResolver) GetTimelinePosts(ctx context.Context, input *model.PaginationInput) (*model.Posts, error) {
	return r.Service.GetTimelinePosts(ctx, input)
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
