package constants

import "github.com/plogto/core/db"

var DEFAULT_USER_SETTINGS = db.UserSettings{
	IsRepliesVisible: db.UserSettingValueOn,
	IsMediaVisible:   db.UserSettingValueOn,
	IsLikesVisible:   db.UserSettingValueOn,
}
