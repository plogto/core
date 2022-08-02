// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
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

type Connections struct {
	TotalCount *int               `json:"totalCount"`
	Edges      []*ConnectionsEdge `json:"edges"`
	PageInfo   *PageInfo          `json:"pageInfo"`
}

type ConnectionsEdge struct {
	Cursor string      `json:"cursor"`
	Node   *Connection `json:"node"`
}

type EditUserInput struct {
	Username        *string          `json:"username"`
	BackgroundColor *BackgroundColor `json:"backgroundColor"`
	PrimaryColor    *PrimaryColor    `json:"primaryColor"`
	Avatar          *string          `json:"avatar"`
	Background      *string          `json:"background"`
	FullName        *string          `json:"fullName"`
	Email           *string          `json:"email"`
	Bio             *string          `json:"bio"`
	IsPrivate       *bool            `json:"isPrivate"`
}

type LikedPosts struct {
	TotalCount *int              `json:"totalCount"`
	Edges      []*LikedPostsEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
}

type LikedPostsEdge struct {
	Cursor string     `json:"cursor"`
	Node   *LikedPost `json:"node"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Notifications struct {
	TotalCount               *int                 `json:"totalCount"`
	Edges                    []*NotificationsEdge `json:"edges"`
	UnreadNotificationsCount *int                 `json:"unreadNotificationsCount"`
	PageInfo                 *PageInfo            `json:"pageInfo"`
}

type NotificationsEdge struct {
	Cursor string        `json:"cursor"`
	Node   *Notification `json:"node"`
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage *bool  `json:"hasNextPage"`
}

type PageInfoInput struct {
	First *int    `json:"first"`
	After *string `json:"after"`
}

type Posts struct {
	TotalCount *int         `json:"totalCount"`
	Edges      []*PostsEdge `json:"edges"`
	PageInfo   *PageInfo    `json:"pageInfo"`
}

type PostsEdge struct {
	Cursor string `json:"cursor"`
	Node   *Post  `json:"node"`
}

type RegisterInput struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SavedPosts struct {
	TotalCount *int              `json:"totalCount"`
	Edges      []*SavedPostsEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
}

type SavedPostsEdge struct {
	Cursor string     `json:"cursor"`
	Node   *SavedPost `json:"node"`
}

type Search struct {
	User *Users `json:"user"`
	Tag  *Tags  `json:"tag"`
}

type Tags struct {
	Edges []*TagsEdge `json:"edges"`
}

type TagsEdge struct {
	Node *Tag `json:"node"`
}

type Test struct {
	Content *string `json:"content"`
}

type TestInput struct {
	Content *string `json:"content"`
}

type Users struct {
	Edges []*UsersEdge `json:"edges"`
}

type UsersEdge struct {
	Node *User `json:"node"`
}

type AddPostInput struct {
	ParentID   *string  `json:"parentId"`
	Content    *string  `json:"content"`
	Status     *string  `json:"status"`
	Attachment []string `json:"attachment"`
}

type EditPostInput struct {
	Content *string `json:"content"`
	Status  *string `json:"status"`
}

type BackgroundColor string

const (
	BackgroundColorLight BackgroundColor = "LIGHT"
	BackgroundColorDim   BackgroundColor = "DIM"
	BackgroundColorDark  BackgroundColor = "DARK"
)

var AllBackgroundColor = []BackgroundColor{
	BackgroundColorLight,
	BackgroundColorDim,
	BackgroundColorDark,
}

func (e BackgroundColor) IsValid() bool {
	switch e {
	case BackgroundColorLight, BackgroundColorDim, BackgroundColorDark:
		return true
	}
	return false
}

func (e BackgroundColor) String() string {
	return string(e)
}

func (e *BackgroundColor) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BackgroundColor(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BackgroundColor", str)
	}
	return nil
}

func (e BackgroundColor) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PrimaryColor string

const (
	PrimaryColorBlue   PrimaryColor = "BLUE"
	PrimaryColorGreen  PrimaryColor = "GREEN"
	PrimaryColorRed    PrimaryColor = "RED"
	PrimaryColorPurple PrimaryColor = "PURPLE"
	PrimaryColorOrange PrimaryColor = "ORANGE"
	PrimaryColorYellow PrimaryColor = "YELLOW"
)

var AllPrimaryColor = []PrimaryColor{
	PrimaryColorBlue,
	PrimaryColorGreen,
	PrimaryColorRed,
	PrimaryColorPurple,
	PrimaryColorOrange,
	PrimaryColorYellow,
}

func (e PrimaryColor) IsValid() bool {
	switch e {
	case PrimaryColorBlue, PrimaryColorGreen, PrimaryColorRed, PrimaryColorPurple, PrimaryColorOrange, PrimaryColorYellow:
		return true
	}
	return false
}

func (e PrimaryColor) String() string {
	return string(e)
}

func (e *PrimaryColor) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PrimaryColor(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PrimaryColor", str)
	}
	return nil
}

func (e PrimaryColor) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
