package service

import (
	"context"
	"regexp"
	"strings"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

func (s *Service) SearchTag(ctx context.Context, expression string) (*model.Tags, error) {
	var limit = constants.TAGS_PAGE_LIMIT

	return s.Tags.GetTagsByTagNameAndPagination(expression+"%", limit)
}

func (s *Service) GetTrends(ctx context.Context, first *int) (*model.Tags, error) {
	var limit int

	if first == nil {
		limit = constants.TRENDS_PAGE_LIMIT
	} else {
		limit = *first
	}

	return s.PostTags.GetTagsOrderByCountTags(limit)
}

func (s *Service) GetTagByID(ctx context.Context, id string) (*model.Tag, error) {
	return s.Tags.GetTagByID(id)
}

func (s *Service) GetTagByName(ctx context.Context, tagName string) (*model.Tag, error) {
	return s.Tags.GetTagByName(tagName)
}

func (s *Service) SaveTagsPost(postID, content string) {
	r := regexp.MustCompile("#(\\w|_)+")
	tags := r.FindAllString(content, -1)
	for i, tag := range tags {
		tags[i] = strings.TrimLeft(tag, "#")
	}
	for _, tagName := range util.UniqueSliceElement(tags) {
		tag := &model.Tag{
			Name: strings.ToLower(tagName),
		}
		s.Tags.CreateTag(tag)

		postTag := &model.PostTag{
			TagID:  tag.ID,
			PostID: postID,
		}
		s.PostTags.CreatePostTag(postTag)
	}
}
