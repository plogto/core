package database

import (
	"fmt"

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
