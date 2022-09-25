// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AddTicketMessageInput struct {
	Message    string    `json:"message"`
	Attachment []*string `json:"attachment"`
}

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

type CreateTicketInput struct {
	Subject    string    `json:"subject"`
	Message    string    `json:"message"`
	Attachment []*string `json:"attachment"`
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

type OAuthGoogleInput struct {
	Credential     string  `json:"credential"`
	InvitationCode *string `json:"invitationCode"`
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

type TicketMessages struct {
	TotalCount *int                  `json:"totalCount"`
	Ticket     *Ticket               `json:"ticket"`
	Edges      []*TicketMessagesEdge `json:"edges"`
	PageInfo   *PageInfo             `json:"pageInfo"`
}

type TicketMessagesEdge struct {
	Cursor string         `json:"cursor"`
	Node   *TicketMessage `json:"node"`
}

type Tickets struct {
	TotalCount *int           `json:"totalCount"`
	Edges      []*TicketsEdge `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
}

type TicketsEdge struct {
	Cursor string  `json:"cursor"`
	Node   *Ticket `json:"node"`
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

type CreditTransactionTemplateName string

const (
	CreditTransactionTemplateNameInviteUser               CreditTransactionTemplateName = "INVITE_USER"
	CreditTransactionTemplateNameRegisterByInvitationCode CreditTransactionTemplateName = "REGISTER_BY_INVITATION_CODE"
)

var AllCreditTransactionTemplateName = []CreditTransactionTemplateName{
	CreditTransactionTemplateNameInviteUser,
	CreditTransactionTemplateNameRegisterByInvitationCode,
}

func (e CreditTransactionTemplateName) IsValid() bool {
	switch e {
	case CreditTransactionTemplateNameInviteUser, CreditTransactionTemplateNameRegisterByInvitationCode:
		return true
	}
	return false
}

func (e CreditTransactionTemplateName) String() string {
	return string(e)
}

func (e *CreditTransactionTemplateName) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CreditTransactionTemplateName(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CreditTransactionTemplateName", str)
	}
	return nil
}

func (e CreditTransactionTemplateName) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CreditTransactionType string

const (
	CreditTransactionTypeOrder      CreditTransactionType = "ORDER"
	CreditTransactionTypeTransfer   CreditTransactionType = "TRANSFER"
	CreditTransactionTypeCommission CreditTransactionType = "COMMISSION"
	CreditTransactionTypeFund       CreditTransactionType = "FUND"
)

var AllCreditTransactionType = []CreditTransactionType{
	CreditTransactionTypeOrder,
	CreditTransactionTypeTransfer,
	CreditTransactionTypeCommission,
	CreditTransactionTypeFund,
}

func (e CreditTransactionType) IsValid() bool {
	switch e {
	case CreditTransactionTypeOrder, CreditTransactionTypeTransfer, CreditTransactionTypeCommission, CreditTransactionTypeFund:
		return true
	}
	return false
}

func (e CreditTransactionType) String() string {
	return string(e)
}

func (e *CreditTransactionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CreditTransactionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CreditTransactionType", str)
	}
	return nil
}

func (e CreditTransactionType) MarshalGQL(w io.Writer) {
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

type TicketPermission string

const (
	TicketPermissionOpen       TicketPermission = "OPEN"
	TicketPermissionClose      TicketPermission = "CLOSE"
	TicketPermissionApprove    TicketPermission = "APPROVE"
	TicketPermissionSolve      TicketPermission = "SOLVE"
	TicketPermissionNewMessage TicketPermission = "NEW_MESSAGE"
)

var AllTicketPermission = []TicketPermission{
	TicketPermissionOpen,
	TicketPermissionClose,
	TicketPermissionApprove,
	TicketPermissionSolve,
	TicketPermissionNewMessage,
}

func (e TicketPermission) IsValid() bool {
	switch e {
	case TicketPermissionOpen, TicketPermissionClose, TicketPermissionApprove, TicketPermissionSolve, TicketPermissionNewMessage:
		return true
	}
	return false
}

func (e TicketPermission) String() string {
	return string(e)
}

func (e *TicketPermission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TicketPermission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TicketPermission", str)
	}
	return nil
}

func (e TicketPermission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TicketStatus string

const (
	TicketStatusOpen     TicketStatus = "OPEN"
	TicketStatusClosed   TicketStatus = "CLOSED"
	TicketStatusApproved TicketStatus = "APPROVED"
	TicketStatusSolved   TicketStatus = "SOLVED"
)

var AllTicketStatus = []TicketStatus{
	TicketStatusOpen,
	TicketStatusClosed,
	TicketStatusApproved,
	TicketStatusSolved,
}

func (e TicketStatus) IsValid() bool {
	switch e {
	case TicketStatusOpen, TicketStatusClosed, TicketStatusApproved, TicketStatusSolved:
		return true
	}
	return false
}

func (e TicketStatus) String() string {
	return string(e)
}

func (e *TicketStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TicketStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TicketStatus", str)
	}
	return nil
}

func (e TicketStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleSuperAdmin UserRole = "SUPER_ADMIN"
	UserRoleAdmin      UserRole = "ADMIN"
	UserRoleUser       UserRole = "USER"
)

var AllUserRole = []UserRole{
	UserRoleSuperAdmin,
	UserRoleAdmin,
	UserRoleUser,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleSuperAdmin, UserRoleAdmin, UserRoleUser:
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
