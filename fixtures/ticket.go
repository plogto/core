package fixtures

import "github.com/plogto/core/graph/model"

var EmptyTicket = &model.Ticket{}
var TicketWithID = &model.Ticket{ID: "id"}
var TicketWithUserID = &model.Ticket{ID: "id", UserID: UserWithID.ID}
var OpenTicket = &model.Ticket{ID: "id", Status: model.TicketStatusOpen}
var ClosedTicket = &model.Ticket{ID: "id", Status: model.TicketStatusClosed}
var AcceptedTicket = &model.Ticket{ID: "id", Status: model.TicketStatusAccepted}
var RejectedTicket = &model.Ticket{ID: "id", Status: model.TicketStatusRejected}
var ApprovedTicket = &model.Ticket{ID: "id", Status: model.TicketStatusApproved}
var SolvedTicket = &model.Ticket{ID: "id", Status: model.TicketStatusSolved}
