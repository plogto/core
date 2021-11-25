// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}

type ChangePasswordInput struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type CommentLikes struct {
	CommentLikes []*CommentLike `json:"commentLikes"`
	Pagination   *Pagination    `json:"pagination"`
}

type CommentPostInput struct {
	ParentID *string `json:"parentId"`
	PostID   string  `json:"postId"`
	Content  string  `json:"content"`
}

type Comments struct {
	Comments   []*Comment  `json:"comments"`
	Pagination *Pagination `json:"pagination"`
}

type Connections struct {
	Connections []*Connection `json:"connections"`
	Pagination  *Pagination   `json:"pagination"`
}

type EditUserInput struct {
	FullName  *string `json:"fullName"`
	Email     *string `json:"email"`
	Bio       *string `json:"bio"`
	IsPrivate *bool   `json:"isPrivate"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Pagination struct {
	TotalDocs  int  `json:"totalDocs"`
	TotalPages int  `json:"totalPages"`
	Limit      int  `json:"limit"`
	Page       int  `json:"page"`
	NextPage   *int `json:"nextPage"`
}

type PaginationInput struct {
	Page  *int `json:"page"`
	Limit *int `json:"limit"`
}

type PostLikes struct {
	PostLikes  []*PostLike `json:"postLikes"`
	Pagination *Pagination `json:"pagination"`
}

type PostSaves struct {
	PostSaves  []*PostSave `json:"postSaves"`
	Pagination *Pagination `json:"pagination"`
}

type Posts struct {
	Posts      []*Post     `json:"posts"`
	Pagination *Pagination `json:"pagination"`
}

type RegisterInput struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Search struct {
	User *Users `json:"user"`
	Tag  *Tags  `json:"tag"`
}

type Tags struct {
	Tags       []*Tag      `json:"tags"`
	Pagination *Pagination `json:"pagination"`
}

type Test struct {
	Content *string `json:"content"`
}

type TestInput struct {
	Content *string `json:"content"`
}

type Users struct {
	Users      []*User     `json:"users"`
	Pagination *Pagination `json:"pagination"`
}

type AddPostInput struct {
	Content    string  `json:"content"`
	Status     *int    `json:"status"`
	Attachment *string `json:"attachment"`
}
