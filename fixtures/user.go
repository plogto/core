package fixtures

import (
	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

var UserID, _ = uuid.NewUUID()
var EmptyUser = &db.User{}
var UserWithID = &db.User{ID: UserID}
var UserWithSuperAdminRole = &db.User{ID: UserID, Role: db.UserRoleSuperAdmin}
var UserWithAdminRole = &db.User{ID: UserID, Role: db.UserRoleAdmin}
var UserWithUserRole = &db.User{ID: UserID, Role: db.UserRoleUser}
