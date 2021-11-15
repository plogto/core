package model

import (
	"time"
)

type Notification struct {
	tableName          struct{}   `sql:"notification"`
	ID                 string     `json:"id"`
	NotificationTypeID string     `json:"notification_type_id"`
	SenderID           string     `json:"sender_id"`
	ReceiverID         string     `json:"receiver_id"`
	PostID             *string    `json:"post_id"`
	CommentID          *string    `json:"comment_id"`
	URL                string     `json:"url"`
	Read               *bool      `json:"read"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"-" sql:",soft_delete"`
}
