package database

import (
	"fmt"
	"time"

	"github.com/favecode/note-core/graph/model"
	"github.com/go-pg/pg"
)

type Follower struct {
	DB *pg.DB
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
