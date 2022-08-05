package database

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Connections struct {
	DB *pg.DB
}

type ConnectionFilter struct {
	Limit  int
	After  string
	Status *int
}

func (c *Connections) GetConnectionsByFieldAndPagination(field, value string, filter ConnectionFilter) (*model.Connections, error) {
	var connections []*model.Connection
	var edges []*model.ConnectionsEdge
	var endCursor string

	query := c.DB.Model(&connections).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil)

	if len(filter.After) > 0 {
		query.Where("created_at < ?", filter.After)
	}

	if filter.Status != nil {
		query.Where("status = ?", *filter.Status)
	}

	totalCount, err :=
		query.Limit(filter.Limit).Order("created_at DESC").SelectAndCount()

	for _, value := range connections {
		edges = append(edges, &model.ConnectionsEdge{Node: &model.Connection{
			ID:          value.ID,
			FollowerID:  value.FollowerID,
			FollowingID: value.FollowingID,
			CreatedAt:   value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > filter.Limit {
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

func (c *Connections) GetFollowersByUserIDAndPagination(followerID string, filter ConnectionFilter) (*model.Connections, error) {
	return c.GetConnectionsByFieldAndPagination("following_id", followerID, filter)
}

func (c *Connections) GetFollowingByUserIDAndPagination(followingID string, filter ConnectionFilter) (*model.Connections, error) {
	return c.GetConnectionsByFieldAndPagination("follower_id", followingID, filter)
}

func (c *Connections) GetFollowRequestsByUserIDAndPagination(followingID string, filter ConnectionFilter) (*model.Connections, error) {
	return c.GetConnectionsByFieldAndPagination("following_id", followingID, filter)
}

func (c *Connections) CreateConnection(connection *model.Connection) (*model.Connection, error) {
	_, err := c.DB.Model(connection).Returning("*").Insert()
	return connection, err
}

func (c *Connections) GetConnectionByField(field, value string) (*model.Connection, error) {
	var connection model.Connection
	err := c.DB.Model(&connection).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &connection, err
}

func (c *Connections) GetConnection(followingID, followerID string) (*model.Connection, error) {
	var connection model.Connection
	err := c.DB.Model(&connection).Where("following_id = ?", followingID).Where("follower_id = ?", followerID).Where("deleted_at is ?", nil).First()
	return &connection, err
}

func (c *Connections) UpdateConnection(connection *model.Connection) (*model.Connection, error) {
	_, err := c.DB.Model(connection).Where("id = ?", connection.ID).Where("deleted_at is ?", nil).Returning("*").Update()
	return connection, err
}

// TODO: fix this name or functionality
func (c *Connections) DeleteConnection(id string) (*model.Connection, error) {
	DeletedAt := time.Now()
	var connection = &model.Connection{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := c.DB.Model(connection).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return connection, err
}

func (c *Connections) CountConnectionByUserID(field, userID string, status int) (int, error) {
	return c.DB.Model((*model.Connection)(nil)).Where(fmt.Sprintf("%v = ?", field), userID).Where("status = ?", status).Where("deleted_at is ?", nil).Count()
}
