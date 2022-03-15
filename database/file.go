package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type File struct {
	DB *pg.DB
}

func (f *File) GetFileByField(field string, value string) (*model.File, error) {
	var file model.File
	err := f.DB.Model(&file).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(file.ID) < 1 {
		return nil, nil
	}
	return &file, err
}

func (f *File) GetFileByHash(hash string) (*model.File, error) {
	return f.GetFileByField("hash", hash)
}

func (f *File) GetFileByName(name string) (*model.File, error) {
	return f.GetFileByField("name", name)
}

func (f *File) CreateFile(file *model.File) (*model.File, error) {
	_, err := f.DB.Model(file).Returning("*").Insert()

	return file, err
}
