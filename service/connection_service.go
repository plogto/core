package service

import (
	"context"
	"errors"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/database"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) FollowUser(ctx context.Context, userID string) (*model.Connection, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not follow yourself")
	}

	followingUser, _ := s.Users.GetUserByID(userID)
	connection, _ := s.Connections.GetConnection(userID, user.ID)
	if len(connection.ID) > 0 {
		return connection, nil
	}

	status := 2
	if followingUser.IsPrivate == bool(true) {
		status = 1
	}

	newConnection := &model.Connection{
		FollowerID:  user.ID,
		FollowingID: userID,
		Status:      &status,
	}

	s.Connections.CreateConnection(newConnection)
	if len(newConnection.ID) > 0 && status == 2 {
		s.CreateNotification(CreateNotificationArgs{
			Name:       constants.NOTIFICATION_FOLLOW_USER,
			SenderID:   user.ID,
			ReceiverID: userID,
			Url:        "/" + user.Username,
		})
	}

	return newConnection, nil
}

func (s *Service) UnfollowUser(ctx context.Context, userID string) (*model.Connection, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not unfollow yourself")
	}

	connection, _ := s.Connections.GetConnection(userID, user.ID)
	if len(connection.ID) < 1 {
		return nil, errors.New("connection not found")
	}

	deletedConnection, _ := s.Connections.DeleteConnection(connection.ID)

	if len(deletedConnection.ID) > 0 {
		s.RemoveNotification(CreateNotificationArgs{
			Name:       constants.NOTIFICATION_FOLLOW_USER,
			SenderID:   user.ID,
			ReceiverID: userID,
			Url:        "/" + user.Username,
		})
	}

	return &model.Connection{
		ID:          deletedConnection.ID,
		FollowerID:  deletedConnection.FollowerID,
		FollowingID: deletedConnection.FollowingID,
		Status:      nil,
		CreatedAt:   deletedConnection.CreatedAt,
		UpdatedAt:   deletedConnection.UpdatedAt,
	}, nil
}

func (s *Service) AcceptUser(ctx context.Context, userID string) (*model.Connection, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not accept yourself")
	}

	connection, _ := s.Connections.GetConnection(user.ID, userID)
	if len(connection.ID) < 1 {
		return nil, errors.New("connection not found")
	}

	if *connection.Status == 2 {
		return connection, nil
	}

	*connection.Status = 2
	updatedConnection, _ := s.Connections.UpdateConnection(connection)
	if len(updatedConnection.ID) > 0 {
		s.CreateNotification(CreateNotificationArgs{
			Name:       constants.NOTIFICATION_ACCEPT_USER,
			SenderID:   user.ID,
			ReceiverID: userID,
			Url:        "/" + user.Username,
		})
	}

	return updatedConnection, nil
}

func (s *Service) RejectUser(ctx context.Context, userID string) (*model.Connection, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not reject yourself")
	}

	connection, _ := s.Connections.GetConnection(user.ID, userID)
	if len(connection.ID) < 1 {
		return nil, errors.New("connection not found")
	}

	if *connection.Status == 2 {
		return nil, errors.New("can not reject accepted request")
	}

	return s.Connections.DeleteConnection(connection.ID)
}

func (s *Service) GetConnectionsByUsername(ctx context.Context, username string, input *model.PageInfoInput, resultType string) (*model.Connections, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	followingUser, _ := s.Users.GetUserByUsername(username)

	pageInfoInput := util.ExtractPageInfo(input)

	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	connectedStatus := 2

	switch resultType {
	case "followers":
		return s.Connections.GetFollowersByUserIDAndPagination(followingUser.ID, database.ConnectionFilter{
			Limit:  *pageInfoInput.First,
			After:  *pageInfoInput.After,
			Status: &connectedStatus,
		})
	case "following":
		return s.Connections.GetFollowingByUserIDAndPagination(followingUser.ID, database.ConnectionFilter{
			Limit:  *pageInfoInput.First,
			After:  *pageInfoInput.After,
			Status: &connectedStatus,
		})
	case "requests":
		status := 1
		return s.Connections.GetFollowRequestsByUserIDAndPagination(followingUser.ID, database.ConnectionFilter{
			Limit:  *pageInfoInput.First,
			After:  *pageInfoInput.After,
			Status: &status,
		})
	}

	return nil, nil
}

func (s *Service) GetFollowRequests(ctx context.Context, input *model.PageInfoInput) (*model.Connections, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return s.GetConnectionsByUsername(ctx, user.Username, input, "requests")
}

func (s *Service) GetConnectionStatus(ctx context.Context, userID string) (*int, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	if user == nil {
		return nil, nil
	}

	connection, err := s.Connections.GetConnection(userID, user.ID)
	if len(connection.ID) < 1 {
		return nil, nil
	}

	return connection.Status, err
}

func (s *Service) GetConnectionCount(ctx context.Context, userID string, resultType string) (int, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	switch resultType {
	case "followers":
		return s.Connections.CountConnectionByUserID("following_id", userID, 2)
	case "following":
		return s.Connections.CountConnectionByUserID("follower_id", userID, 2)
	case "requests":
		if user == nil || user.ID != userID {
			return 0, nil
		}
		return s.Connections.CountConnectionByUserID("following_id", userID, 1)
	}

	return 0, nil
}
