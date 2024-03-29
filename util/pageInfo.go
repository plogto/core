package util

import (
	"encoding/base64"
	"time"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/graph/model"
)

type PageInfoResult struct {
	First int32
	After time.Time
}

func ExtractPageInfo(params *model.PageInfoInput) (pageInfo *PageInfoResult) {
	after := time.Now()

	result := &PageInfoResult{
		First: constants.POSTS_PAGE_LIMIT,
		After: after,
	}

	if params != nil {
		if params.First != nil {
			result.First = int32(*params.First)
		}

		if params.After != nil {
			date := ConvertCursorToDateTime(*params.After)

			if time, err := time.Parse(time.RFC3339Nano, date); err == nil {
				result.After = time
			}
		}
	}

	return result
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
