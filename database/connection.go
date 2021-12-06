package database

import (
	"fmt"
	"time"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/util"
	"github.com/go-pg/pg/v10"
)

type Connection struct {
	DB *pg.DB
}

type ConnectionFilter struct {
	Limit  int
	Page   int
	Status *int
}

func (c *Connection) GetConnectionsByFieldAndPagination(field, value string, filter ConnectionFilter) (*model.Connections, error) {
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

func (c *Connection) GetFollowersByUserIdAndPagination(followerId string, filter ConnectionFilter) (*model.Connections, error) {
	return c.GetConnectionsByFieldAndPagination("following_id", followerId, filter)
}

func (c *Connection) GetFollowingByUserIdAndPagination(followingId string, filter ConnectionFilter) (*model.Connections, error) {
	return c.GetConnectionsByFieldAndPagination("follower_id", followingId, filter)
}

func (c *Connection) GetFollowRequestsByUserIdAndPagination(followingId string, filter ConnectionFilter) (*model.Connections, error) {
	return c.GetConnectionsByFieldAndPagination("following_id", followingId, filter)
}

func (c *Connection) CreateConnection(connection *model.Connection) (*model.Connection, error) {
	_, err := c.DB.Model(connection).Returning("*").Insert()
	return connection, err
}

func (c *Connection) GetConnectionByField(field, value string) (*model.Connection, error) {
	var connection model.Connection
	err := c.DB.Model(&connection).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &connection, err
}

func (c *Connection) GetConnection(followingId, followerId string) (*model.Connection, error) {
	var connection model.Connection
	err := c.DB.Model(&connection).Where("following_id = ?", followingId).Where("follower_id = ?", followerId).Where("deleted_at is ?", nil).First()
	return &connection, err
}

func (c *Connection) UpdateConnection(connection *model.Connection) (*model.Connection, error) {
	_, err := c.DB.Model(connection).Where("id = ?", connection.ID).Where("deleted_at is ?", nil).Returning("*").Update()
	return connection, err
}

// TODO: fix this name or functionality
func (c *Connection) DeleteConnection(id string) (*model.Connection, error) {
	DeletedAt := time.Now()
	var connection = &model.Connection{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := c.DB.Model(connection).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return connection, err
}

func (c *Connection) CountConnectionByUserId(field, userId string, status int) (*int, error) {
	count, err := c.DB.Model((*model.Connection)(nil)).Where(fmt.Sprintf("%v = ?", field), userId).Where("status = ?", status).Where("deleted_at is ?", nil).Count()
	return &count, err
}
