package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Connections struct {
	Queries *db.Queries
}

type ConnectionFilter struct {
	Limit  int32
	After  time.Time
	Status int32
}

func (c *Connections) CreateConnection(ctx context.Context, arg db.CreateConnectionParams) (*db.Connection, error) {
	connection, err := c.Queries.CreateConnection(ctx, arg)

	if err != nil {
		return nil, err
	}

	return connection, nil
}

func (c *Connections) GetFollowersByUserIDAndPageInfo(ctx context.Context, followerID string, filter ConnectionFilter) (*model.Connections, error) {
	var edges []*model.ConnectionsEdge
	var endCursor string

	// FIXME
	FollowerID, _ := uuid.Parse(followerID)

	connections, err := c.Queries.GetFollowersByUserIDAndPageInfo(ctx, db.GetFollowersByUserIDAndPageInfoParams{
		Limit:       filter.Limit,
		FollowingID: FollowerID,
		Status:      2,
		CreatedAt:   filter.After,
	})

	totalCount, _ := c.Queries.CountFollowersByUserIDAndPageInfo(ctx, db.CountFollowersByUserIDAndPageInfoParams{
		Limit:       filter.Limit,
		FollowingID: FollowerID,
		Status:      2,
		CreatedAt:   filter.After,
	})

	for _, value := range connections {
		edges = append(edges, &model.ConnectionsEdge{Node: &db.Connection{
			ID:          value.ID,
			FollowerID:  value.FollowerID,
			FollowingID: value.FollowingID,
			Status:      value.Status,
			CreatedAt:   value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > int64(filter.Limit) {
		hasNextPage = true
	}

	return &model.Connections{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (c *Connections) GetFollowingByUserIDAndPageInfo(ctx context.Context, followingID string, filter ConnectionFilter) (*model.Connections, error) {

	var edges []*model.ConnectionsEdge
	var endCursor string

	// FIXME
	FollowingID, _ := uuid.Parse(followingID)

	connections, err := c.Queries.GetFollowingByUserIDAndPageInfo(ctx, db.GetFollowingByUserIDAndPageInfoParams{
		Limit:      filter.Limit,
		FollowerID: FollowingID,
		Status:     2,
		CreatedAt:  filter.After,
	})

	totalCount, _ := c.Queries.CountFollowingByUserIDAndPageInfo(ctx, db.CountFollowingByUserIDAndPageInfoParams{
		Limit:      filter.Limit,
		FollowerID: FollowingID,
		Status:     2,
		CreatedAt:  filter.After,
	})

	for _, value := range connections {
		edges = append(edges, &model.ConnectionsEdge{Node: &db.Connection{
			ID:          value.ID,
			FollowerID:  value.FollowerID,
			FollowingID: value.FollowingID,
			Status:      value.Status,
			CreatedAt:   value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > int64(filter.Limit) {
		hasNextPage = true
	}

	return &model.Connections{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (c *Connections) GetConnectionByID(ctx context.Context, id uuid.UUID) (*db.Connection, error) {
	// FIXME: use dataloader
	connection, err := c.Queries.GetConnectionByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return connection, nil
}

func (c *Connections) GetConnection(ctx context.Context, followingID, followerID string) (*db.Connection, error) { // FIXME
	FollowerID, _ := uuid.Parse(followerID)
	FollowingID, _ := uuid.Parse(followingID)

	connection, err := c.Queries.GetConnection(ctx, db.GetConnectionParams{
		FollowerID:  FollowerID,
		FollowingID: FollowingID,
	})

	if err != nil {
		return nil, err
	}

	return connection, nil
}

func (c *Connections) CountFollowingConnectionByUserID(ctx context.Context, userID uuid.UUID, status int32) (int64, error) {
	totalCount, err := c.Queries.CountFollowingByUserIDAndPageInfo(ctx, db.CountFollowingByUserIDAndPageInfoParams{
		FollowerID: userID,
		Status:     status,
	})

	if err != nil {
		return 0, err
	}

	return totalCount, nil
}

func (c *Connections) CountFollowersConnectionByUserID(ctx context.Context, userID uuid.UUID, status int32) (int64, error) {
	totalCount, err := c.Queries.CountFollowersByUserIDAndPageInfo(ctx, db.CountFollowersByUserIDAndPageInfoParams{
		FollowingID: userID,
		Status:      status,
	})

	if err != nil {
		return 0, err
	}

	return totalCount, nil
}

func (c *Connections) UpdateConnection(ctx context.Context, arg db.UpdateConnectionParams) (*db.Connection, error) {
	connection, err := c.Queries.UpdateConnection(ctx, arg)

	if err != nil {
		return nil, err
	}

	return connection, nil
}

// TODO: fix this name or functionality
func (c *Connections) DeleteConnection(ctx context.Context, id uuid.UUID) (*db.Connection, error) {
	// FIXME
	DeletedAt := sql.NullTime{time.Now(), true}

	connection, err := c.Queries.DeleteConnection(ctx, db.DeleteConnectionParams{
		ID:        id,
		DeletedAt: DeletedAt,
	})

	if err != nil {
		return nil, err
	}
	return connection, nil
}
