package convertor

import (
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

func DBTagsToModel(tags []*db.Tag) []*model.Tag {
	var result []*model.Tag

	for _, tag := range tags {
		result = append(result, &model.Tag{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	return result
}

func DBTagToModel(tag *db.Tag) *model.Tag {
	return &model.Tag{
		ID:   tag.ID,
		Name: tag.Name,
	}
}
