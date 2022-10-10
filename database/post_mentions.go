package database

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type PostMentions struct {
	DB *pg.DB
}

func (p *PostMentions) CreatePostMention(postMention *model.PostMention) (*model.PostMention, error) {
	_, err := p.DB.Model(postMention).Returning("*").Insert()
	return postMention, err
}

func (p *PostMentions) GetPostMentionByField(field string, value string) (*model.PostMention, error) {
	var postMention model.PostMention
	err := p.DB.Model(&postMention).
		Where(fmt.Sprintf("%v = ?", field), value).
		Where("deleted_at is ?", nil).
		First()

	return &postMention, err
}

func (p *PostMentions) GetPostMentionsByPostID(postID string) ([]*model.PostMention, error) {
	var postMentions []*model.PostMention
	err := p.DB.Model(&postMentions).
		Where("post_id = ?", postID).
		Where("deleted_at is ?", nil).
		First()

	return postMentions, err
}

func (p *PostMentions) DeletePostMention(postMention *model.PostMention) (*model.PostMention, error) {
	DeletedAt := time.Now()
	postMention.DeletedAt = &DeletedAt

	_, err := p.DB.Model(postMention).
		Set("deleted_at = ?deleted_at").
		Where("post_id =?post_id").
		Where("user_id =?user_id").
		Where("deleted_at is ?", nil).
		Update()

	return postMention, err
}

func (p *PostMentions) DeletePostMentionsByPostID(postMention *model.PostMention) ([]*model.PostMention, error) {
	var postMentions []*model.PostMention
	query := p.DB.Model(&postMentions).
		Where("post_id = ?", postMention.PostID)

	_, err := query.Set("deleted_at = ?", postMention.DeletedAt).Update()
	return postMentions, err
}
