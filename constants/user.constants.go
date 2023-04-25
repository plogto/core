package constants

import "github.com/plogto/core/db"

var DEFAULT_USER_SETTINGS = db.UserSettings{
	RepliesVisible:               db.UserSettingValueOn,
	MediaVisible:                 db.UserSettingValueOn,
	LikesVisible:                 db.UserSettingValueOn,
	RepliesVisibleForCurrentUser: db.UserSettingValueOn,
	MediaVisibleForCurrentUser:   db.UserSettingValueOn,
	LikesVisibleForCurrentUser:   db.UserSettingValueOn,
}
