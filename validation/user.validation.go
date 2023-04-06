package validation

import (
	"github.com/plogto/core/db"
	"github.com/samber/lo"
)

func IsUserExists(user *db.User) bool {
	return user != nil && lo.IsNotEmpty(user.ID)
}

func IsSuperAdmin(user *db.User) bool {
	return IsUserExists(user) && user.Role == db.UserRoleSuperAdmin
}

func IsAdmin(user *db.User) bool {
	return IsUserExists(user) && user.Role == db.UserRoleAdmin
}

func IsUser(user *db.User) bool {
	return IsUserExists(user) && user.Role == db.UserRoleUser
}
