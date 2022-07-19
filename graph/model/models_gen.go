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
	Connections []*Connection `json:"connections"`
	Pagination  *Pagination   `json:"pagination"`
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
	LikedPosts []*LikedPost `json:"likedPosts"`
	Pagination *Pagination  `json:"pagination"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Notifications struct {
	Notifications            []*Notification `json:"notifications"`
	UnreadNotificationsCount *int            `json:"unreadNotificationsCount"`
	Pagination               *Pagination     `json:"pagination"`
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
