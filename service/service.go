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
	Notifications                         database.Notifications
	Passwords                             database.Passwords
	PostAttachments                       database.PostAttachments
	PostMentions                          database.PostMentions
	PostTags                              database.PostTags
	Posts                                 database.Posts
	SavedPosts                            database.SavedPosts
	Tags                                  database.Tags
	TicketMessageAttachments              database.TicketMessageAttachments
	TicketMessages                        database.TicketMessages
	Users                                 database.Users
	Tickets                               database.Tickets
	OnlineUsers                           database.OnlineUsers
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
		Notifications:                         service.Notifications,
		Passwords:                             service.Passwords,
		PostAttachments:                       service.PostAttachments,
		PostMentions:                          service.PostMentions,
		PostTags:                              service.PostTags,
		Posts:                                 service.Posts,
		SavedPosts:                            service.SavedPosts,
		Tags:                                  service.Tags,
		TicketMessageAttachments:              service.TicketMessageAttachments,
		TicketMessages:                        service.TicketMessages,
		Users:                                 service.Users,
		Tickets:                               service.Tickets,
		OnlineUsers:                           service.OnlineUsers,
		OnlineNotifications:                   map[string]chan *model.NotificationsEdge{},
	}
}

func (s *Service) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}
