package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type Files struct {
	DB *pg.DB
}

func (f *Files) GetFileByField(field string, value string) (*model.File, error) {
	var file model.File
	err := f.DB.Model(&file).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(file.ID) < 1 {
		return nil, nil
	}
	return &file, err
}

func (f *Files) GetFilesByPostId(postID string) ([]*model.File, error) {
	var files []*model.File
	err := f.DB.Model(&files).
		ColumnExpr("file.*").
		ColumnExpr("post_attachments.file_id").
		Join("INNER JOIN post_attachments ON post_attachments.file_id = file.id").
		GroupExpr("post_attachments.file_id, file.id").
		Where("post_attachments.post_id = ?", postID).
		Where("file.deleted_at is ?", nil).
		Where("post_attachments.deleted_at is ?", nil).
		Returning("*").Select()

	return files, err
}

func (f *Files) GetFileByHash(hash string) (*model.File, error) {
	return f.GetFileByField("hash", hash)
}

func (f *Files) GetFileByName(name string) (*model.File, error) {
	return f.GetFileByField("name", name)
}

func (f *Files) GetFileByID(id string) (*model.File, error) {
	return f.GetFileByField("id", id)
}

func (f *Files) CreateFile(file *model.File) (*model.File, error) {
	_, err := f.DB.Model(file).Returning("*").Insert()
	if len(file.ID) < 1 {
		return nil, err
	}
	return file, err
}
