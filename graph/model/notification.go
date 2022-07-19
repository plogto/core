package model

import (
	"time"
)

type Notification struct {
	tableName          struct{} `pg:"notifications"`
	ID                 string
	NotificationTypeID string
	SenderID           string
	ReceiverID         string
	PostID             *string
	ReplyID            *string
	URL                string
	Read               *bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time `pg:"-,soft_delete"`
}
