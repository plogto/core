package service

import (
	"context"
	"errors"

	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

func (s *Service) EditUserSettings(ctx context.Context, input model.EditUserSettingsInput) (*db.User, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	didUpdate := false

	if input.IsRepliesVisible != nil {
		user.Settings.IsRepliesVisible = db.UserSettingValue(*input.IsRepliesVisible)
		didUpdate = true
	}

	if input.IsMediaVisible != nil {
		user.Settings.IsMediaVisible = db.UserSettingValue(*input.IsMediaVisible)
		didUpdate = true
	}

	if input.IsLikesVisible != nil {
		user.Settings.IsLikesVisible = db.UserSettingValue(*input.IsLikesVisible)
		didUpdate = true
	}

	if !didUpdate {
		return nil, nil
	}

	updatedUser, _ := s.Users.UpdateUserSettings(ctx, user.ID, user.Settings)

	return updatedUser, nil
}

func (s *Service) IsSettingValueOff(settingValue db.UserSettingValue) bool {
	return settingValue == db.UserSettingValueOff
}
