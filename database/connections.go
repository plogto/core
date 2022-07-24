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
	Page   int
	Status *int
}

func (c *Connections) GetConnectionsByFieldAndPagination(field, value string, filter ConnectionFilter) (*model.Connections, error) {
	var connections []*model.Connection
	var offset = (filter.Page - 1) * filter.Limit

	query := c.DB.Model(&connections).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil)

	if filter.Status != nil {
		query.Where("status = ?", *filter.Status)
	}

	query.Offset(offset).Limit(filter.Limit)

	totalDocs, err := query.Order("created_at DESC").Returning("*").SelectAndCount()

	return &model.Connections{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     filter.Limit,
			Page:      filter.Page,
			TotalDocs: totalDocs,
		}),
		Connections: connections,
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
