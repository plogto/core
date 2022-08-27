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

type CreateCreditTransactionInput struct {
	SenderID    string  `json:"senderId"`
	ReceiverID  string  `json:"receiverId"`
	Amount      float64 `json:"amount"`
	Description *string `json:"description"`
}

type CreditTransactions struct {
	TotalCount *int                      `json:"totalCount"`
	Edges      []*CreditTransactionsEdge `json:"edges"`
	PageInfo   *PageInfo                 `json:"pageInfo"`
}

type CreditTransactionsEdge struct {
	Cursor string             `json:"cursor"`
	Node   *CreditTransaction `json:"node"`
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

type InvitedUsers struct {
	TotalCount *int                `json:"totalCount"`
	Edges      []*InvitedUsersEdge `json:"edges"`
	PageInfo   *PageInfo           `json:"pageInfo"`
}

type InvitedUsersEdge struct {
	Cursor string       `json:"cursor"`
	Node   *InvitedUser `json:"node"`
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
	FullName       string  `json:"fullName"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	InvitationCode *string `json:"invitationCode"`
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

type CreditTransactionDescriptionVariableType string

const (
	CreditTransactionDescriptionVariableTypeUser CreditTransactionDescriptionVariableType = "USER"
	CreditTransactionDescriptionVariableTypeTag  CreditTransactionDescriptionVariableType = "TAG"
)

var AllCreditTransactionDescriptionVariableType = []CreditTransactionDescriptionVariableType{
	CreditTransactionDescriptionVariableTypeUser,
	CreditTransactionDescriptionVariableTypeTag,
}

func (e CreditTransactionDescriptionVariableType) IsValid() bool {
	switch e {
	case CreditTransactionDescriptionVariableTypeUser, CreditTransactionDescriptionVariableTypeTag:
		return true
	}
	return false
}

func (e CreditTransactionDescriptionVariableType) String() string {
	return string(e)
}

func (e *CreditTransactionDescriptionVariableType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CreditTransactionDescriptionVariableType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CreditTransactionDescriptionVariableType", str)
	}
	return nil
}

func (e CreditTransactionDescriptionVariableType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CreditTransactionStatus string

const (
	CreditTransactionStatusApproved CreditTransactionStatus = "APPROVED"
	CreditTransactionStatusPending  CreditTransactionStatus = "PENDING"
	CreditTransactionStatusFailed   CreditTransactionStatus = "FAILED"
	CreditTransactionStatusCanceled CreditTransactionStatus = "CANCELED"
)

var AllCreditTransactionStatus = []CreditTransactionStatus{
	CreditTransactionStatusApproved,
	CreditTransactionStatusPending,
	CreditTransactionStatusFailed,
	CreditTransactionStatusCanceled,
}

func (e CreditTransactionStatus) IsValid() bool {
	switch e {
	case CreditTransactionStatusApproved, CreditTransactionStatusPending, CreditTransactionStatusFailed, CreditTransactionStatusCanceled:
		return true
	}
	return false
}

func (e CreditTransactionStatus) String() string {
	return string(e)
}

func (e *CreditTransactionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CreditTransactionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CreditTransactionStatus", str)
	}
	return nil
}

func (e CreditTransactionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CreditTransactionTypeName string

const (
	CreditTransactionTypeNameInviteUser               CreditTransactionTypeName = "INVITE_USER"
	CreditTransactionTypeNameRegisterByInvitationCode CreditTransactionTypeName = "REGISTER_BY_INVITATION_CODE"
)

var AllCreditTransactionTypeName = []CreditTransactionTypeName{
	CreditTransactionTypeNameInviteUser,
	CreditTransactionTypeNameRegisterByInvitationCode,
}

func (e CreditTransactionTypeName) IsValid() bool {
	switch e {
	case CreditTransactionTypeNameInviteUser, CreditTransactionTypeNameRegisterByInvitationCode:
		return true
	}
	return false
}

func (e CreditTransactionTypeName) String() string {
	return string(e)
}

func (e *CreditTransactionTypeName) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CreditTransactionTypeName(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CreditTransactionTypeName", str)
	}
	return nil
}

func (e CreditTransactionTypeName) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PostStatus string

const (
	PostStatusPublic  PostStatus = "PUBLIC"
	PostStatusPrivate PostStatus = "PRIVATE"
)

var AllPostStatus = []PostStatus{
	PostStatusPublic,
	PostStatusPrivate,
}

func (e PostStatus) IsValid() bool {
	switch e {
	case PostStatusPublic, PostStatusPrivate:
		return true
	}
	return false
}

func (e PostStatus) String() string {
	return string(e)
}

func (e *PostStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PostStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PostStatus", str)
	}
	return nil
}

func (e PostStatus) MarshalGQL(w io.Writer) {
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

type UserRole string

const (
	UserRoleAdmin UserRole = "ADMIN"
	UserRoleUser  UserRole = "USER"
)

var AllUserRole = []UserRole{
	UserRoleAdmin,
	UserRoleUser,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleAdmin, UserRoleUser:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
