package validation

import (
	"github.com/plogto/core/graph/model"
	"github.com/samber/lo"
)

func IsUserExists(user *model.User) bool {
	return user != nil && lo.IsNotEmpty(user.ID)
}

func IsSuperAdmin(user *model.User) bool {
	return IsUserExists(user) && user.Role == model.UserRoleSuperAdmin
}

func IsAdmin(user *model.User) bool {
	return IsUserExists(user) && user.Role == model.UserRoleAdmin
}

func IsUser(user *model.User) bool {
	return IsUserExists(user) && user.Role == model.UserRoleUser
}
