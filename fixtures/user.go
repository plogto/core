package fixtures

import "github.com/plogto/core/graph/model"

var EmptyUser = &model.User{}
var UserWithID = &model.User{ID: "id"}
var UserWithSuperAdminRole = &model.User{ID: "id", Role: model.UserRoleSuperAdmin}
var UserWithAdminRole = &model.User{ID: "id", Role: model.UserRoleAdmin}
var UserWithUserRole = &model.User{ID: "id", Role: model.UserRoleUser}
