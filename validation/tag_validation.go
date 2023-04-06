package validation

import (
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

type TagType interface {
	db.Tag | model.Tag
}

// TODO: fix validation
func IsTagExists[T TagType](tag *T) bool {
	return tag != nil
}
