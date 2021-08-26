package database

import (
	"fmt"
	"time"

	"github.com/favecode/note-core/graph/model"
	"github.com/favecode/note-core/util"
	"github.com/go-pg/pg"
)

type Follower struct {
	DB *pg.DB
}

func (f *Follower) GetFollowersByFieldAndPagination(field string, value string, limit int, page int) (*model.Followers, error) {
	var followers []*model.Follower
	var offset = (page - 1) * limit

	query := f.DB.Model(&followers).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.Followers{
		Pagination: util.GetPatination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		Followers: followers,
	}, err
}

func (f *Follower) GetFollowersByUserIdAndPagination(userId string, limit int, page int) (*model.Followers, error) {
	return f.GetFollowersByFieldAndPagination("user_id", userId, limit, page)
}

func (f *Follower) GetFollowingByUserIdAndPagination(userId string, limit int, page int) (*model.Followers, error) {
	return f.GetFollowersByFieldAndPagination("follower_id", userId, limit, page)
}

func (f *Follower) CreateFollower(follower *model.Follower) (*model.Follower, error) {
	_, err := f.DB.Model(follower).Returning("*").Insert()
	return follower, err
}

func (f *Follower) GetFollowerByField(field, value string) (*model.Follower, error) {
	var follower model.Follower
	err := f.DB.Model(&follower).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &follower, err
}

func (f *Follower) GetFollowerByUserIdAndFollowerId(userId string, followerId string) (*model.Follower, error) {
	var follower model.Follower
	err := f.DB.Model(&follower).Where("user_id = ?", userId).Where("follower_id = ?", followerId).Where("deleted_at is ?", nil).First()
	return &follower, err
}

func (f *Follower) UpdateFollower(follower *model.Follower) (*model.Follower, error) {
	_, err := f.DB.Model(follower).Where("id = ?", follower.ID).Where("deleted_at is ?", nil).Returning("*").Update()
	return follower, err
}

func (f *Follower) DeleteFollower(id string) (*model.Follower, error) {
	DeletedAt := time.Now()
	var follower = &model.Follower{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := f.DB.Model(follower).Set("deleted_at = ?deleted_at").Where("id = ?id").Where("deleted_at is ?", nil).Returning("*").Update()
	return follower, err
}
