package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type InvitedUsers struct {
	DB *pg.DB
}

func (i *InvitedUsers) CreateInvitedUser(invitedUser *model.InvitedUser) (*model.InvitedUser, error) {
	_, err := i.DB.Model(invitedUser).
		Where("inviter_id = ?inviter_id").
		Where("invitee_id = ?invitee_id").
		Insert()

	return invitedUser, err
}

func (i *InvitedUsers) UpdateInvitedUser(invitedUser *model.InvitedUser) (*model.InvitedUser, error) {
	_, err := i.DB.Model(invitedUser).Where("id = ?", invitedUser.ID).Where("deleted_at is ?", nil).Returning("*").Update()

	return invitedUser, err
}
