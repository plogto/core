package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/plogto/core/database"
	"github.com/plogto/core/graph/model"
)

type Service struct {
	Connections                           database.Connections
	CreditTransactionDescriptionVariables database.CreditTransactionDescriptionVariables
	CreditTransactionTemplates            database.CreditTransactionTemplates
	CreditTransactionInfos                database.CreditTransactionInfos
	CreditTransactions                    database.CreditTransactions
	Files                                 database.Files
	InvitedUsers                          database.InvitedUsers
	LikedPosts                            database.LikedPosts
	NotificationTypes                     database.NotificationTypes
	Users                                 database.Users
	Passwords                             database.Passwords
	Posts                                 database.Posts
	Tickets                               database.Tickets
	TicketMessages                        database.TicketMessages
	Tags                                  database.Tags
	TicketMessageAttachments              database.TicketMessageAttachments
	PostAttachments                       database.PostAttachments
	PostTags                              database.PostTags
	PostMentions                          database.PostMentions
	SavedPosts                            database.SavedPosts
	OnlineUsers                           database.OnlineUsers
	Notifications                         database.Notifications
	OnlineNotifications                   map[string]chan *model.NotificationsEdge
	mu                                    sync.Mutex
}

func New(service Service) *Service {
	return &Service{
		Connections:                           service.Connections,
		CreditTransactionDescriptionVariables: service.CreditTransactionDescriptionVariables,
		CreditTransactionTemplates:            service.CreditTransactionTemplates,
		CreditTransactionInfos:                service.CreditTransactionInfos,
		CreditTransactions:                    service.CreditTransactions,
		Files:                                 service.Files,
		InvitedUsers:                          service.InvitedUsers,
		LikedPosts:                            service.LikedPosts,
		NotificationTypes:                     service.NotificationTypes,
		Users:                                 service.Users,
		Passwords:                             service.Passwords,
		Posts:                                 service.Posts,
		Tickets:                               service.Tickets,
		TicketMessages:                        service.TicketMessages,
		Tags:                                  service.Tags,
		PostTags:                              service.PostTags,
		PostMentions:                          service.PostMentions,
		TicketMessageAttachments:              service.TicketMessageAttachments,
		PostAttachments:                       service.PostAttachments,
		SavedPosts:                            service.SavedPosts,
		OnlineUsers:                           service.OnlineUsers,
		Notifications:                         service.Notifications,
		OnlineNotifications:                   map[string]chan *model.NotificationsEdge{},
	}
}

func (s *Service) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}
