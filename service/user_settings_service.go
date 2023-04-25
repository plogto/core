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

	if input.RepliesVisible != nil {
		user.Settings.RepliesVisible = db.UserSettingValue(*input.RepliesVisible)
		didUpdate = true
	}

	if input.MediaVisible != nil {
		user.Settings.MediaVisible = db.UserSettingValue(*input.MediaVisible)
		didUpdate = true
	}

	if input.LikesVisible != nil {
		user.Settings.LikesVisible = db.UserSettingValue(*input.LikesVisible)
		didUpdate = true
	}

	if input.RepliesVisibleForCurrentUser != nil {
		user.Settings.RepliesVisibleForCurrentUser = db.UserSettingValue(*input.RepliesVisibleForCurrentUser)
		didUpdate = true
	}

	if input.MediaVisibleForCurrentUser != nil {
		user.Settings.MediaVisibleForCurrentUser = db.UserSettingValue(*input.MediaVisibleForCurrentUser)
		didUpdate = true
	}

	if input.LikesVisibleForCurrentUser != nil {
		user.Settings.LikesVisibleForCurrentUser = db.UserSettingValue(*input.LikesVisibleForCurrentUser)
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
