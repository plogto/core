package service

import (
	"context"
	"errors"

	"github.com/favecode/poster-core/config"
	"github.com/favecode/poster-core/database"
	"github.com/favecode/poster-core/graph/model"
	"github.com/favecode/poster-core/middleware"
)

func (s *Service) FollowUser(ctx context.Context, userID string) (*model.Connection, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not follow yourself")
	}

	followingUser, _ := s.User.GetUserByID(userID)

	connection, _ := s.Connection.GetConnection(userID, user.ID)

	if len(connection.ID) > 0 {
		return connection, nil
	}

	status := 2
	if followingUser.Private == bool(true) {
		status = 1
	}

	newConnection := &model.Connection{
		FollowerID:  user.ID,
		FollowingID: userID,
		Status:      &status,
	}

	s.Connection.CreateConnection(newConnection)

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

	connection, _ := s.Connection.GetConnection(userID, user.ID)

	if len(connection.ID) < 1 {
		return nil, errors.New("connection not found")
	}

	deletedConnection, _ := s.Connection.DeleteConnection(connection.ID)

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

	connection, _ := s.Connection.GetConnection(user.ID, userID)

	if len(connection.ID) < 1 {
		return nil, errors.New("connection not found")
	}

	if *connection.Status == 2 {
		return connection, nil
	}

	*connection.Status = 2

	updatedConnection, _ := s.Connection.UpdateConnection(connection)

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

	connection, _ := s.Connection.GetConnection(user.ID, userID)

	if len(connection.ID) < 1 {
		return nil, errors.New("connection not found")
	}

	if *connection.Status == 2 {
		return nil, errors.New("can not reject accepted request")
	}

	deletedConnection, _ := s.Connection.DeleteConnection(connection.ID)

	return deletedConnection, nil
}

func (s *Service) GetUserConnectionsByUsername(ctx context.Context, username string, input *model.GetUserConnectionsByUserIDInput, resultType string) (*model.Connections, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	followingUser, _ := s.User.GetUserByUsername(username)

	var limit int = config.POSTS_PAGE_LIMIT
	var page int = 1

	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Page != nil && *input.Page > 0 {
			page = *input.Page
		}
	}

	connection, _ := s.Connection.GetConnection(followingUser.ID, user.ID)
	if followingUser.ID != user.ID {
		if followingUser.Private == bool(true) {
			if len(connection.ID) < 1 || *connection.Status == 2 {
				return nil, errors.New("you need to follow this user")
			}
		}
	}

	switch resultType {
	case "follower":
		connections, _ := s.Connection.GetFollowersByUserIdAndPagination(followingUser.ID, database.ConnectionFilter{
			Limit: limit,
			Page:  page,
		})
		return connections, nil
	case "following":
		connections, _ := s.Connection.GetFollowingByUserIdAndPagination(followingUser.ID, database.ConnectionFilter{
			Limit: limit,
			Page:  page,
		})
		return connections, nil
	case "requested":
		status := 1
		connections, _ := s.Connection.GetRequestedByUserIdAndPagination(followingUser.ID, database.ConnectionFilter{
			Limit:  limit,
			Page:   page,
			Status: &status,
		})
		return connections, nil
	}

	return nil, nil
}

func (s *Service) GetConnectionStatus(ctx context.Context, userId string) (*int, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	connection, _ := s.Connection.GetConnection(userId, user.ID)

	return connection.Status, nil
}
