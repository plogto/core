package service

import (
	"context"
	"regexp"
	"strings"

	"github.com/plogto/core/config"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

func (s *Service) SearchTag(ctx context.Context, expression string) (*model.Tags, error) {
	limit := 10
	page := 1
	tags, _ := s.Tag.GetTagsByTagNameAndPagination(expression+"%", limit, page)

	return tags, nil
}

func (s *Service) CountTagByTagID(ctx context.Context, tagID string) (*int, error) {
	tagCount, _ := s.PostTag.CountPostTagsByTagID(tagID)

	return tagCount, nil
}

func (s *Service) GetTrends(ctx context.Context, input *model.PaginationInput) (*model.Tags, error) {
	var limit int = config.POSTS_PAGE_LIMIT
	var page int = 1

	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Page != nil && *input.Page > 0 {
			page = *input.Page
		}
	}
	tags, _ := s.PostTag.GetTagsOrderByCountTags(limit, page)

	return tags, nil
}

func (s *Service) GetTagByName(ctx context.Context, tagName string) (*model.Tag, error) {
	tag, _ := s.Tag.GetTagByName(tagName)

	if len(tag.ID) < 1 {
		return nil, nil
	}

	return tag, nil
}

func (s *Service) SaveTagsPost(postID, content string) {
	r := regexp.MustCompile("#(\\w|_)+")
	tags := r.FindAllString(content, -1)
	for i, tag := range tags {
		tags[i] = strings.TrimLeft(tag, "#")
	}
	for _, tagName := range util.UniqueSliceElement(tags) {
		tag := &model.Tag{
			Name: tagName,
		}
		s.Tag.CreateTag(tag)

		postTag := &model.PostTag{
			TagID:  tag.ID,
			PostID: postID,
		}
		s.PostTag.CreatePostTag(postTag)
	}
}
