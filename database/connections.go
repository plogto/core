package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
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
	return c.Queries.CreateConnection(ctx, arg)
}

func (c *Connections) GetFollowersByUserIDAndPageInfo(ctx context.Context, followerID pgtype.UUID, filter ConnectionFilter) (*model.Connections, error) {
	var edges []*model.ConnectionsEdge
	var endCursor string

	connections, _ := c.Queries.GetFollowersByUserIDAndPageInfo(ctx, db.GetFollowersByUserIDAndPageInfoParams{
		Limit:       filter.Limit,
		FollowingID: followerID,
		Status:      2,
		CreatedAt:   filter.After,
	})

	totalCount, _ := c.Queries.CountFollowersByUserIDAndPageInfo(ctx, db.CountFollowersByUserIDAndPageInfoParams{
		FollowingID: followerID,
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
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (c *Connections) GetFollowingByUserIDAndPageInfo(ctx context.Context, followingID pgtype.UUID, filter ConnectionFilter) (*model.Connections, error) {
	var edges []*model.ConnectionsEdge
	var endCursor string

	connections, _ := c.Queries.GetFollowingByUserIDAndPageInfo(ctx, db.GetFollowingByUserIDAndPageInfoParams{
		Limit:      filter.Limit,
		FollowerID: followingID,
		Status:     2,
		CreatedAt:  filter.After,
	})

	totalCount, _ := c.Queries.CountFollowingByUserIDAndPageInfo(ctx, db.CountFollowingByUserIDAndPageInfoParams{
		FollowerID: followingID,
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
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (c *Connections) GetConnectionByID(ctx context.Context, id pgtype.UUID) (*db.Connection, error) {
	// TODO: use dataloader
	return c.Queries.GetConnectionByID(ctx, id)
}

func (c *Connections) GetConnection(ctx context.Context, followingID, followerID pgtype.UUID) (*db.Connection, error) {
	return c.Queries.GetConnection(ctx, db.GetConnectionParams{
		FollowerID:  followerID,
		FollowingID: followingID,
	})
}

func (c *Connections) CountFollowingConnectionByUserID(ctx context.Context, userID pgtype.UUID, status int32) (int64, error) {
	createdAt := time.Now()

	totalCount, _ := c.Queries.CountFollowingByUserIDAndPageInfo(ctx, db.CountFollowingByUserIDAndPageInfoParams{
		CreatedAt:  createdAt,
		FollowerID: userID,
		Status:     status,
	})

	return totalCount, nil
}

func (c *Connections) CountFollowersConnectionByUserID(ctx context.Context, userID pgtype.UUID, status int32) (int64, error) {
	createdAt := time.Now()

	totalCount, _ := c.Queries.CountFollowersByUserIDAndPageInfo(ctx, db.CountFollowersByUserIDAndPageInfoParams{
		CreatedAt:   createdAt,
		FollowingID: userID,
		Status:      status,
	})

	return totalCount, nil
}

func (c *Connections) UpdateConnection(ctx context.Context, arg db.UpdateConnectionParams) (*db.Connection, error) {
	return c.Queries.UpdateConnection(ctx, arg)
}

// TODO: fix this name or functionality
func (c *Connections) DeleteConnection(ctx context.Context, id pgtype.UUID) (*db.Connection, error) {
	DeletedAt := time.Now()

	return c.Queries.DeleteConnection(ctx, db.DeleteConnectionParams{
		ID:        id,
		DeletedAt: &DeletedAt,
	})
}
