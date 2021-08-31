package database

import (
	"fmt"
	"time"

	"github.com/favecode/poster-core/graph/model"
	"github.com/favecode/poster-core/util"
	"github.com/go-pg/pg"
)

type Connection struct {
	DB *pg.DB
}

type ConnectionFilter struct {
	Limit  int
	Page   int
	Status *int
}

func (f *Connection) GetConnectionsByFieldAndPagination(field string, value string, filter ConnectionFilter) (*model.Connections, error) {
	var connections []*model.Connection
	var offset = (filter.Page - 1) * filter.Limit

	query := f.DB.Model(&connections).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil)

	if filter.Status != nil {
		query.Where("status = ?", *filter.Status)
	}

	query.Offset(offset).Limit(filter.Limit)

	totalDocs, err := query.Order("created_at DESC").Returning("*").SelectAndCount()

	return &model.Connections{
		Pagination: util.GetPatination(&util.GetPaginationParams{
			Limit:     filter.Limit,
			Page:      filter.Page,
			TotalDocs: totalDocs,
		}),
		Connections: connections,
	}, err
}

func (f *Connection) GetFollowersByUserIdAndPagination(followerId string, filter ConnectionFilter) (*model.Connections, error) {
	return f.GetConnectionsByFieldAndPagination("following_id", followerId, filter)
}

func (f *Connection) GetFollowingByUserIdAndPagination(followingId string, filter ConnectionFilter) (*model.Connections, error) {
	return f.GetConnectionsByFieldAndPagination("follower_id", followingId, filter)
}

func (f *Connection) GetFollowRequestsByUserIdAndPagination(followingId string, filter ConnectionFilter) (*model.Connections, error) {
	return f.GetConnectionsByFieldAndPagination("following_id", followingId, filter)
}

func (f *Connection) CreateConnection(connection *model.Connection) (*model.Connection, error) {
	_, err := f.DB.Model(connection).Returning("*").Insert()
	return connection, err
}

func (f *Connection) GetConnectionByField(field, value string) (*model.Connection, error) {
	var connection model.Connection
	err := f.DB.Model(&connection).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &connection, err
}

func (f *Connection) GetConnection(followingId string, followerId string) (*model.Connection, error) {
	var connection model.Connection
	err := f.DB.Model(&connection).Where("following_id = ?", followingId).Where("follower_id = ?", followerId).Where("deleted_at is ?", nil).First()
	return &connection, err
}

func (f *Connection) UpdateConnection(connection *model.Connection) (*model.Connection, error) {
	_, err := f.DB.Model(connection).Where("id = ?", connection.ID).Where("deleted_at is ?", nil).Returning("*").Update()
	return connection, err
}

func (f *Connection) DeleteConnection(id string) (*model.Connection, error) {
	DeletedAt := time.Now()
	var connection = &model.Connection{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := f.DB.Model(connection).Set("deleted_at = ?deleted_at").Where("id = ?id").Where("deleted_at is ?", nil).Returning("*").Update()
	return connection, err
}
