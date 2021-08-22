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

type GetUserPostsByUsernameInput struct {
	Page  *int `json:"page"`
	Limit *int `json:"limit"`
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

type Posts struct {
	Posts      []*Post     `json:"posts"`
	Pagination *Pagination `json:"pagination"`
}

type RegisterInput struct {
	Fullname *string `json:"fullname"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type Test struct {
	Content *string `json:"content"`
}

type TestInput struct {
	Content *string `json:"content"`
}

type AddPostInput struct {
	Content string `json:"content"`
	Status  *int   `json:"status"`
}
