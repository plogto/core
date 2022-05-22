package util

import (
	"math"

	"github.com/plogto/core/graph/model"
)

type GetPaginationParams struct {
	Limit     int
	Page      int
	TotalDocs int
}

func GetPagination(params *GetPaginationParams) (pagination *model.Pagination) {
	var totalPages = int(math.Ceil(float64(params.TotalDocs) / float64(params.Limit)))
	var nextPage *int

	if params.Page < totalPages {
		n := params.Page + 1
		nextPage = &n
	}

	return &model.Pagination{
		Page:       params.Page,
		Limit:      params.Limit,
		TotalDocs:  params.TotalDocs,
		TotalPages: totalPages,
		NextPage:   nextPage,
	}
}
