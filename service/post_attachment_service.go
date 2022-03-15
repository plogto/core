package service

import (
	"context"
)

func (s *Service) GetPostAttachmentsByPostID(ctx context.Context, postID string) ([]string, error) {
	postAttachments, _ := s.PostAttachment.GetPostAttachmentsByPostID(postID)

	var attachments []string

	for _, v := range postAttachments {
		attachments = append(attachments, v.Name)
	}
	return attachments, nil
}
