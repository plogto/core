package validation

import "github.com/plogto/core/graph/model"

func IsSuperAdmin(user *model.User) bool {
	return len(user.ID) > 0 && user.Role == model.UserRoleSuperAdmin
}

func IsAdmin(user *model.User) bool {
	return len(user.ID) > 0 && user.Role == model.UserRoleAdmin
}

func IsUser(user *model.User) bool {
	return len(user.ID) > 0 && user.Role == model.UserRoleUser
}

func IsUserExists(user *model.User) bool {
	return len(user.ID) > 0
}
