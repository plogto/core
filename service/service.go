package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/plogto/core/database"
	"github.com/plogto/core/graph/model"
)

type Service struct {
	Users                                 database.Users
	Passwords                             database.Passwords
	Posts                                 database.Posts
	Files                                 database.Files
	Connections                           database.Connections
	CreditTransactions                    database.CreditTransactions
	CreditTransactionTemplates            database.CreditTransactionTemplates
	CreditTransactionInfos                database.CreditTransactionInfos
	CreditTransactionDescriptionVariables database.CreditTransactionDescriptionVariables
	Tickets                               database.Tickets
	TicketMessages                        database.TicketMessages
	Tags                                  database.Tags
	TicketMessageAttachments              database.TicketMessageAttachments
	PostAttachments                       database.PostAttachments
	PostTags                              database.PostTags
	PostMentions                          database.PostMentions
	LikedPosts                            database.LikedPosts
	SavedPosts                            database.SavedPosts
	InvitedUsers                          database.InvitedUsers
	OnlineUsers                           database.OnlineUsers
	Notifications                         database.Notifications
	NotificationTypes                     database.NotificationTypes
	OnlineNotifications                   map[string]chan *model.NotificationsEdge
	mu                                    sync.Mutex
}

func New(service Service) *Service {
	return &Service{
		Users:                                 service.Users,
		Passwords:                             service.Passwords,
		Posts:                                 service.Posts,
		Files:                                 service.Files,
		Connections:                           service.Connections,
		CreditTransactions:                    service.CreditTransactions,
		CreditTransactionTemplates:            service.CreditTransactionTemplates,
		CreditTransactionInfos:                service.CreditTransactionInfos,
		CreditTransactionDescriptionVariables: service.CreditTransactionDescriptionVariables,
		Tickets:                               service.Tickets,
		TicketMessages:                        service.TicketMessages,
		Tags:                                  service.Tags,
		PostTags:                              service.PostTags,
		PostMentions:                          service.PostMentions,
		TicketMessageAttachments:              service.TicketMessageAttachments,
		PostAttachments:                       service.PostAttachments,
		LikedPosts:                            service.LikedPosts,
		SavedPosts:                            service.SavedPosts,
		InvitedUsers:                          service.InvitedUsers,
		OnlineUsers:                           service.OnlineUsers,
		Notifications:                         service.Notifications,
		NotificationTypes:                     service.NotificationTypes,
		OnlineNotifications:                   map[string]chan *model.NotificationsEdge{},
	}
}

func (s *Service) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}
