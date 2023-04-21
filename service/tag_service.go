package service

import (
	"context"
	"regexp"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/convertor"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

func (s *Service) SearchTag(ctx context.Context, expression string) (*model.Tags, error) {
	var limit = constants.TAGS_PAGE_LIMIT

	return s.Tags.GetTagsByTagNameAndPageInfo(ctx, expression+"%", limit)
}

func (s *Service) GetTrends(ctx context.Context, first *int) (*model.Tags, error) {
	var limit int32

	if first == nil {
		limit = constants.TRENDS_PAGE_LIMIT
	} else {
		limit = int32(*first)
	}

	return s.PostTags.GetTagsOrderByCountTags(ctx, limit)
}

func (s *Service) GetTagByID(ctx context.Context, id pgtype.UUID) (*model.Tag, error) {
	return graph.GetTagLoader(ctx).Load(convertor.UUIDToString(id))
}

func (s *Service) GetTagByName(ctx context.Context, tagName string) (*model.Tag, error) {
	return s.Tags.GetTagByName(ctx, tagName)
}

func (s *Service) SaveTagsPost(ctx context.Context, postID pgtype.UUID, content string) {
	r := regexp.MustCompile("#(\\w|_)+")
	tags := r.FindAllString(content, -1)
	for i, tag := range tags {
		tags[i] = strings.TrimLeft(tag, "#")
	}
	for _, tagName := range util.UniqueSliceElement(tags) {
		tag, _ := s.Tags.CreateTag(ctx, strings.ToLower(tagName))

		s.PostTags.CreatePostTag(ctx, tag.ID, postID)
	}
}
