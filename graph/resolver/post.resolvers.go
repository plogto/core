package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

// AddPost is the resolver for the addPost field.
func (r *mutationResolver) AddPost(ctx context.Context, input model.AddPostInput) (*model.Post, error) {
	return r.Service.AddPost(ctx, input)
}

// EditPost is the resolver for the editPost field.
func (r *mutationResolver) EditPost(ctx context.Context, postID string, input model.EditPostInput) (*model.Post, error) {
	return r.Service.EditPost(ctx, convertor.StringToUUID(postID), input)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, postID string) (*model.Post, error) {
	return r.Service.DeletePost(ctx, convertor.StringToUUID(postID))
}

// Parent is the resolver for the parent field.
func (r *postResolver) Parent(ctx context.Context, obj *model.Post) (*model.Post, error) {
	if !obj.ParentID.Valid {
		return nil, nil
	} else {
		return r.Service.GetPostByID(ctx, obj.ParentID)
	}
}

// Child is the resolver for the child field.
func (r *postResolver) Child(ctx context.Context, obj *model.Post) (*model.Post, error) {
	if !obj.ChildID.Valid {
		return nil, nil
	} else {
		return r.Service.GetPostByID(ctx, obj.ChildID)
	}
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (*db.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// Content is the resolver for the content field.
func (r *postResolver) Content(ctx context.Context, obj *model.Post) (*string, error) {
	return r.Service.GetPostContentByPostID(ctx, obj.ID)
}

// Attachment is the resolver for the attachment field.
func (r *postResolver) Attachment(ctx context.Context, obj *model.Post) ([]*db.File, error) {
	return obj.Attachment, nil
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
func (r *postResolver) IsLiked(ctx context.Context, obj *model.Post) (*db.LikedPost, error) {
	return r.Service.IsPostLiked(ctx, obj.ID)
}

// IsSaved is the resolver for the isSaved field.
func (r *postResolver) IsSaved(ctx context.Context, obj *model.Post) (*db.SavedPost, error) {
	return r.Service.IsPostSaved(ctx, obj.ID)
}

// IsReposted is the resolver for the isReposted field.
func (r *postResolver) IsReposted(ctx context.Context, obj *model.Post) (*model.Post, error) {
	if !obj.ID.Valid {
		return nil, nil
	} else {
		return r.Service.GetChildPostByPostID(ctx, obj.ID)
	}
}

// Cursor is the resolver for the cursor field.
func (r *postsEdgeResolver) Cursor(ctx context.Context, obj *model.PostsEdge) (string, error) {
	return util.ConvertCreateAtToCursor(obj.Node.CreatedAt), nil
}

// Node is the resolver for the node field.
func (r *postsEdgeResolver) Node(ctx context.Context, obj *model.PostsEdge) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, obj.Node.ID)
}

// GetPostsByUsername is the resolver for the getPostsByUsername field.
func (r *queryResolver) GetPostsByUsername(ctx context.Context, username string, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	return r.Service.GetPostsByUsername(ctx, username, pageInfo)
}

// GetPostsWithAttachmentByUsername is the resolver for the getPostsWithAttachmentByUsername field.
func (r *queryResolver) GetPostsWithAttachmentByUsername(ctx context.Context, username string, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	return r.Service.GetPostsWithAttachmentByUsername(ctx, username, pageInfo)
}

// GetRepliesByUsername is the resolver for the getRepliesByUsername field.
func (r *queryResolver) GetRepliesByUsername(ctx context.Context, username string, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	return r.Service.GetRepliesByUsername(ctx, username, pageInfo)
}

// GetPostsByTagName is the resolver for the getPostsByTagName field.
func (r *queryResolver) GetPostsByTagName(ctx context.Context, tagName string, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	return r.Service.GetPostsByTagName(ctx, tagName, pageInfo)
}

// GetPostByURL is the resolver for the getPostByUrl field.
func (r *queryResolver) GetPostByURL(ctx context.Context, url string) (*model.Post, error) {
	return r.Service.GetPostByURL(ctx, url)
}

// GetTimelinePosts is the resolver for the getTimelinePosts field.
func (r *queryResolver) GetTimelinePosts(ctx context.Context, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	return r.Service.GetTimelinePosts(ctx, pageInfo)
}

// GetExplorePosts is the resolver for the getExplorePosts field.
func (r *queryResolver) GetExplorePosts(ctx context.Context, input *model.GetExplorePostsInput, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	return r.Service.GetExplorePosts(ctx, input, pageInfo)
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// PostsEdge returns generated.PostsEdgeResolver implementation.
func (r *Resolver) PostsEdge() generated.PostsEdgeResolver { return &postsEdgeResolver{r} }

type postResolver struct{ *Resolver }
type postsEdgeResolver struct{ *Resolver }
