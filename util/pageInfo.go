package util

import (
	"encoding/base64"
	"time"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/graph/model"
)

func ExtractPageInfo(params *model.PageInfoInput) (pageInfo *model.PageInfoInput) {
	var first int = constants.POSTS_PAGE_FIRST
	var after string

	if params != nil {
		if params.First != nil {
			first = *params.First
		}

		if params.After != nil {
			after = ConvertCursorToDateTime(*params.After)
		}
	}

	return &model.PageInfoInput{
		First: &first,
		After: &after,
	}
}

func ConvertCreateAtToCursor(createdAt time.Time) string {
	dateTime := createdAt.Format(time.RFC3339Nano)
	cursor := base64.StdEncoding.EncodeToString([]byte(dateTime))

	return cursor
}

func ConvertCursorToDateTime(cursor string) string {
	decodedAfter, _ := base64.StdEncoding.DecodeString(cursor)

	return string(decodedAfter)
}
