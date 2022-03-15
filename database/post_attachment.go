package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type PostAttachment struct {
	DB *pg.DB
}

func (p *PostAttachment) CreatePostAttachment(postAttachment *model.PostAttachment) (*model.PostAttachment, error) {
	_, err := p.DB.Model(postAttachment).Returning("*").Insert()
	return postAttachment, err
}

func (p *PostAttachment) GetPostAttachmentsByField(field string, value string) ([]*model.PostAttachment, error) {
	var postAttachments []*model.PostAttachment
	err := p.DB.Model(&postAttachments).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).Select()
	if len(postAttachments) < 1 {
		return nil, nil
	}
	return postAttachments, err
}

func (p *PostAttachment) GetPostAttachmentsByPostID(postID string) ([]*model.PostAttachment, error) {
	return p.GetPostAttachmentsByField("post_id", postID)
}
