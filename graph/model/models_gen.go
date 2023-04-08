// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type AddPostInput struct {
	ParentID   *uuid.UUID  `json:"parentId,omitempty"`
	Content    *string     `json:"content,omitempty"`
	Status     *PostStatus `json:"status,omitempty"`
	Attachment []string    `json:"attachment,omitempty"`
}

type AddTicketMessageInput struct {
	Message    string    `json:"message"`
	Attachment []*string `json:"attachment,omitempty"`
}

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *db.User   `json:"user"`
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
	TotalCount int64              `json:"totalCount"`
	Edges      []*ConnectionsEdge `json:"edges"`
	PageInfo   *PageInfo          `json:"pageInfo"`
}

type ConnectionsEdge struct {
	Cursor string         `json:"cursor"`
	Node   *db.Connection `json:"node,omitempty"`
}

type CreateTicketInput struct {
	Subject    string    `json:"subject"`
	Message    string    `json:"message"`
	Attachment []*string `json:"attachment,omitempty"`
}

type CreditTransactions struct {
	TotalCount int64                     `json:"totalCount"`
	Edges      []*CreditTransactionsEdge `json:"edges"`
	PageInfo   *PageInfo                 `json:"pageInfo"`
}

type CreditTransactionsEdge struct {
	Cursor string                `json:"cursor"`
	Node   *db.CreditTransaction `json:"node,omitempty"`
}

type EditPostInput struct {
	Content *string     `json:"content,omitempty"`
	Status  *PostStatus `json:"status,omitempty"`
}

type EditUserInput struct {
	Username        *string          `json:"username,omitempty"`
	BackgroundColor *BackgroundColor `json:"backgroundColor,omitempty"`
	PrimaryColor    *PrimaryColor    `json:"primaryColor,omitempty"`
	Avatar          *string          `json:"avatar,omitempty"`
	Background      *string          `json:"background,omitempty"`
	FullName        *string          `json:"fullName,omitempty"`
	Email           *string          `json:"email,omitempty"`
	Bio             *string          `json:"bio,omitempty"`
	IsPrivate       *bool            `json:"isPrivate,omitempty"`
}

type GetExplorePostsInput struct {
	IsAttachment *bool `json:"isAttachment,omitempty"`
}

type InvitedUser struct {
	ID        uuid.UUID  `json:"id"`
	Inviter   *db.User   `json:"inviter"`
	Invitee   *db.User   `json:"invitee"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type InvitedUsers struct {
	TotalCount int64               `json:"totalCount"`
	Edges      []*InvitedUsersEdge `json:"edges"`
	PageInfo   *PageInfo           `json:"pageInfo"`
}

type InvitedUsersEdge struct {
	Cursor string       `json:"cursor"`
	Node   *InvitedUser `json:"node,omitempty"`
}

type LikedPosts struct {
	TotalCount int64             `json:"totalCount"`
	Edges      []*LikedPostsEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
}

type LikedPostsEdge struct {
	Cursor string        `json:"cursor"`
	Node   *db.LikedPost `json:"node,omitempty"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Notifications struct {
	TotalCount               int64                `json:"totalCount"`
	Edges                    []*NotificationsEdge `json:"edges"`
	UnreadNotificationsCount int64                `json:"unreadNotificationsCount"`
	PageInfo                 *PageInfo            `json:"pageInfo"`
}

type NotificationsEdge struct {
	Cursor string           `json:"cursor"`
	Node   *db.Notification `json:"node,omitempty"`
}

type OAuthGoogleInput struct {
	Credential     string  `json:"credential"`
	InvitationCode *string `json:"invitationCode,omitempty"`
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type PageInfoInput struct {
	First *int    `json:"first,omitempty"`
	After *string `json:"after,omitempty"`
}

type Posts struct {
	TotalCount int64        `json:"totalCount"`
	Edges      []*PostsEdge `json:"edges"`
	PageInfo   *PageInfo    `json:"pageInfo"`
}

type PostsEdge struct {
	Cursor string   `json:"cursor"`
	Node   *db.Post `json:"node,omitempty"`
}

type RegisterInput struct {
	FullName       string  `json:"fullName"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	InvitationCode *string `json:"invitationCode,omitempty"`
}

type SavedPosts struct {
	TotalCount int64             `json:"totalCount"`
	Edges      []*SavedPostsEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
}

type SavedPostsEdge struct {
	Cursor string        `json:"cursor"`
	Node   *db.SavedPost `json:"node,omitempty"`
}

type Search struct {
	User *Users `json:"user,omitempty"`
	Tag  *Tags  `json:"tag,omitempty"`
}

type Tags struct {
	Edges []*TagsEdge `json:"edges"`
}

type TagsEdge struct {
	Node *Tag `json:"node,omitempty"`
}

type Test struct {
	Content *string `json:"content,omitempty"`
}

type TestInput struct {
	Content *string `json:"content,omitempty"`
}

type TicketMessages struct {
	TotalCount int64                 `json:"totalCount"`
	Ticket     *db.Ticket            `json:"ticket,omitempty"`
	Edges      []*TicketMessagesEdge `json:"edges"`
	PageInfo   *PageInfo             `json:"pageInfo"`
}

type TicketMessagesEdge struct {
	Cursor string            `json:"cursor"`
	Node   *db.TicketMessage `json:"node,omitempty"`
}

type Tickets struct {
	TotalCount int64          `json:"totalCount"`
	Edges      []*TicketsEdge `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
}

type TicketsEdge struct {
	Cursor string     `json:"cursor"`
	Node   *db.Ticket `json:"node,omitempty"`
}

type Users struct {
	Edges []*UsersEdge `json:"edges"`
}

type UsersEdge struct {
	Node *db.User `json:"node,omitempty"`
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

type CreditTransactionDescriptionVariableKey string

const (
	CreditTransactionDescriptionVariableKeyTicket      CreditTransactionDescriptionVariableKey = "ticket"
	CreditTransactionDescriptionVariableKeyInvitedUser CreditTransactionDescriptionVariableKey = "invited_user"
	CreditTransactionDescriptionVariableKeyInviterUser CreditTransactionDescriptionVariableKey = "inviter_user"
)

var AllCreditTransactionDescriptionVariableKey = []CreditTransactionDescriptionVariableKey{
	CreditTransactionDescriptionVariableKeyTicket,
	CreditTransactionDescriptionVariableKeyInvitedUser,
	CreditTransactionDescriptionVariableKeyInviterUser,
}

func (e CreditTransactionDescriptionVariableKey) IsValid() bool {
	switch e {
	case CreditTransactionDescriptionVariableKeyTicket, CreditTransactionDescriptionVariableKeyInvitedUser, CreditTransactionDescriptionVariableKeyInviterUser:
		return true
	}
	return false
}

func (e CreditTransactionDescriptionVariableKey) String() string {
	return string(e)
}

func (e *CreditTransactionDescriptionVariableKey) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CreditTransactionDescriptionVariableKey(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CreditTransactionDescriptionVariableKey", str)
	}
	return nil
}

func (e CreditTransactionDescriptionVariableKey) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CreditTransactionDescriptionVariableType string

const (
	CreditTransactionDescriptionVariableTypeTicket CreditTransactionDescriptionVariableType = "ticket"
	CreditTransactionDescriptionVariableTypeUser   CreditTransactionDescriptionVariableType = "user"
	CreditTransactionDescriptionVariableTypeTag    CreditTransactionDescriptionVariableType = "tag"
)

var AllCreditTransactionDescriptionVariableType = []CreditTransactionDescriptionVariableType{
	CreditTransactionDescriptionVariableTypeTicket,
	CreditTransactionDescriptionVariableTypeUser,
	CreditTransactionDescriptionVariableTypeTag,
}

func (e CreditTransactionDescriptionVariableType) IsValid() bool {
	switch e {
	case CreditTransactionDescriptionVariableTypeTicket, CreditTransactionDescriptionVariableTypeUser, CreditTransactionDescriptionVariableTypeTag:
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
	CreditTransactionTemplateNameApproveTicket            CreditTransactionTemplateName = "APPROVE_TICKET"
)

var AllCreditTransactionTemplateName = []CreditTransactionTemplateName{
	CreditTransactionTemplateNameInviteUser,
	CreditTransactionTemplateNameRegisterByInvitationCode,
	CreditTransactionTemplateNameApproveTicket,
}

func (e CreditTransactionTemplateName) IsValid() bool {
	switch e {
	case CreditTransactionTemplateNameInviteUser, CreditTransactionTemplateNameRegisterByInvitationCode, CreditTransactionTemplateNameApproveTicket:
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

type NotificationTypeName string

const (
	NotificationTypeNameWelcome       NotificationTypeName = "WELCOME"
	NotificationTypeNameLikePost      NotificationTypeName = "LIKE_POST"
	NotificationTypeNameReplyPost     NotificationTypeName = "REPLY_POST"
	NotificationTypeNameLikeReply     NotificationTypeName = "LIKE_REPLY"
	NotificationTypeNameFollowUser    NotificationTypeName = "FOLLOW_USER"
	NotificationTypeNameAcceptUser    NotificationTypeName = "ACCEPT_USER"
	NotificationTypeNameMentionInPost NotificationTypeName = "MENTION_IN_POST"
)

var AllNotificationTypeName = []NotificationTypeName{
	NotificationTypeNameWelcome,
	NotificationTypeNameLikePost,
	NotificationTypeNameReplyPost,
	NotificationTypeNameLikeReply,
	NotificationTypeNameFollowUser,
	NotificationTypeNameAcceptUser,
	NotificationTypeNameMentionInPost,
}

func (e NotificationTypeName) IsValid() bool {
	switch e {
	case NotificationTypeNameWelcome, NotificationTypeNameLikePost, NotificationTypeNameReplyPost, NotificationTypeNameLikeReply, NotificationTypeNameFollowUser, NotificationTypeNameAcceptUser, NotificationTypeNameMentionInPost:
		return true
	}
	return false
}

func (e NotificationTypeName) String() string {
	return string(e)
}

func (e *NotificationTypeName) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationTypeName(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationTypeName", str)
	}
	return nil
}

func (e NotificationTypeName) MarshalGQL(w io.Writer) {
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
	TicketPermissionAccept     TicketPermission = "ACCEPT"
	TicketPermissionApprove    TicketPermission = "APPROVE"
	TicketPermissionReject     TicketPermission = "REJECT"
	TicketPermissionSolve      TicketPermission = "SOLVE"
	TicketPermissionNewMessage TicketPermission = "NEW_MESSAGE"
)

var AllTicketPermission = []TicketPermission{
	TicketPermissionOpen,
	TicketPermissionClose,
	TicketPermissionAccept,
	TicketPermissionApprove,
	TicketPermissionReject,
	TicketPermissionSolve,
	TicketPermissionNewMessage,
}

func (e TicketPermission) IsValid() bool {
	switch e {
	case TicketPermissionOpen, TicketPermissionClose, TicketPermissionAccept, TicketPermissionApprove, TicketPermissionReject, TicketPermissionSolve, TicketPermissionNewMessage:
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
	TicketStatusAccepted TicketStatus = "ACCEPTED"
	TicketStatusApproved TicketStatus = "APPROVED"
	TicketStatusRejected TicketStatus = "REJECTED"
	TicketStatusSolved   TicketStatus = "SOLVED"
)

var AllTicketStatus = []TicketStatus{
	TicketStatusOpen,
	TicketStatusClosed,
	TicketStatusAccepted,
	TicketStatusApproved,
	TicketStatusRejected,
	TicketStatusSolved,
}

func (e TicketStatus) IsValid() bool {
	switch e {
	case TicketStatusOpen, TicketStatusClosed, TicketStatusAccepted, TicketStatusApproved, TicketStatusRejected, TicketStatusSolved:
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
