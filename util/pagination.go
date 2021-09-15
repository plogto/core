package util

import (
	"github.com/favecode/plog-core/graph/model"
)

type GetPaginationParams struct {
	Limit     int
	Page      int
	TotalDocs int
}

func GetPatination(params *GetPaginationParams) (pagination *model.Pagination) {
	var totalPages int = params.TotalDocs / params.Limit
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
