package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/database"
	"github.com/plogto/core/db"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

func (s *Service) FollowUser(ctx context.Context, userID uuid.UUID) (*db.Connection, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not follow yourself")
	}

	followingUser, _ := graph.GetUserLoader(ctx).Load(userID.String())
	connection, _ := s.Connections.GetConnection(ctx, userID, user.ID)
	if validation.IsConnectionExists(connection) {
		return connection, nil
	}

	status := 2
	if followingUser.IsPrivate == bool(true) {
		status = 1
	}

	newConnection, _ := s.Connections.CreateConnection(ctx, db.CreateConnectionParams{
		FollowerID:  user.ID,
		FollowingID: userID,
		Status:      int32(status),
	})

	if validation.IsConnectionStatusAccepted(newConnection) {
		s.CreateNotification(ctx, CreateNotificationArgs{
			Name:       db.NotificationTypeNameFollowUser,
			SenderID:   user.ID,
			ReceiverID: userID,
			Url:        "/" + user.Username,
		})
	}

	return newConnection, nil
}

func (s *Service) UnfollowUser(ctx context.Context, userID uuid.UUID) (*db.Connection, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not unfollow yourself")
	}

	connection, _ := s.Connections.GetConnection(ctx, userID, user.ID)
	if !validation.IsConnectionExists(connection) {
		return nil, errors.New("connection not found")
	}

	deletedConnection, _ := s.Connections.DeleteConnection(ctx, connection.ID)

	if validation.IsConnectionExists(deletedConnection) {
		s.RemoveNotification(ctx, CreateNotificationArgs{
			Name:       db.NotificationTypeNameFollowUser,
			SenderID:   user.ID,
			ReceiverID: userID,
			Url:        "/" + user.Username,
		})
	}

	return &db.Connection{
		ID:          deletedConnection.ID,
		FollowerID:  deletedConnection.FollowerID,
		FollowingID: deletedConnection.FollowingID,
		Status:      0,
		CreatedAt:   deletedConnection.CreatedAt,
		UpdatedAt:   deletedConnection.UpdatedAt,
	}, nil
}

func (s *Service) AcceptUser(ctx context.Context, userID uuid.UUID) (*db.Connection, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not accept yourself")
	}

	connection, _ := s.Connections.GetConnection(ctx, user.ID, userID)

	if !validation.IsConnectionExists(connection) {
		return nil, errors.New("connection not found")
	}

	if connection.Status == 2 {
		return connection, nil
	}

	connection.Status = 2
	updatedConnection, _ := s.Connections.UpdateConnection(ctx, db.UpdateConnectionParams{
		FollowerID:  connection.FollowerID,
		FollowingID: connection.FollowingID,
		Status:      connection.Status,
	})
	if validation.IsConnectionExists(connection) {
		s.CreateNotification(ctx, CreateNotificationArgs{
			Name:       db.NotificationTypeNameAcceptUser,
			SenderID:   user.ID,
			ReceiverID: userID,
			Url:        "/" + user.Username,
		})
	}

	return updatedConnection, nil
}

func (s *Service) RejectUser(ctx context.Context, userID uuid.UUID) (*db.Connection, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not reject yourself")
	}

	connection, _ := s.Connections.GetConnection(ctx, user.ID, userID)
	if !validation.IsConnectionExists(connection) {
		return nil, errors.New("connection not found")
	}

	if validation.IsConnectionStatusAccepted(connection) {
		return nil, errors.New("can not reject accepted request")
	}

	return s.Connections.DeleteConnection(ctx, connection.ID)
}

func (s *Service) GetConnectionsByUsername(ctx context.Context, username string, input *model.PageInfoInput, resultType constants.ConnectionResult) (*model.Connections, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	followingUser, _ := s.Users.GetUserByUsername(ctx, username)

	pageInfo := util.ExtractPageInfo(input)

	if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	limit := int32(pageInfo.First)

	switch resultType {
	case constants.Followers:
		return s.Connections.GetFollowersByUserIDAndPageInfo(ctx, followingUser.ID, database.ConnectionFilter{
			Limit:  limit,
			After:  pageInfo.After,
			Status: 2,
		})
	case constants.Following:
		return s.Connections.GetFollowingByUserIDAndPageInfo(ctx, followingUser.ID, database.ConnectionFilter{
			Limit:  limit,
			After:  pageInfo.After,
			Status: 2,
		})
	case constants.Requests:
		return s.Connections.GetFollowersByUserIDAndPageInfo(ctx, followingUser.ID, database.ConnectionFilter{
			Limit:  limit,
			After:  pageInfo.After,
			Status: 1,
		})
	}

	return nil, nil
}

func (s *Service) GetFollowRequests(ctx context.Context, input *model.PageInfoInput) (*model.Connections, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return s.GetConnectionsByUsername(ctx, user.Username, input, constants.Requests)
}

func (s *Service) GetConnectionStatus(ctx context.Context, userID uuid.UUID) (*int, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	if user == nil {
		return nil, nil
	}

	connection, err := s.Connections.GetConnection(ctx, userID, user.ID)

	zeroStatus := 0
	if !validation.IsConnectionExists(connection) {
		return &zeroStatus, nil
	}

	status := int(connection.Status)

	return &status, err
}

func (s *Service) GetConnectionCount(ctx context.Context, userID uuid.UUID, resultType constants.ConnectionResult) (int64, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	switch resultType {
	case constants.Followers:
		return s.Connections.CountFollowersConnectionByUserID(ctx, userID, 2)
	case constants.Following:
		return s.Connections.CountFollowingConnectionByUserID(ctx, userID, 2)
	case constants.Requests:
		if user == nil || user.ID != userID {
			return 0, nil
		}
		return s.Connections.CountFollowersConnectionByUserID(ctx, userID, 1)
	}

	return 0, nil
}
