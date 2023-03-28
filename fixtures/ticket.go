package fixtures

import (
	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

var TicketID, _ = uuid.NewUUID()
var EmptyTicket = &db.Ticket{}
var TicketWithID = &db.Ticket{ID: TicketID}

// TODO: remove uuid.MustParse
var TicketWithUserID = &db.Ticket{ID: TicketID, UserID: uuid.MustParse(UserWithID.ID)}
var OpenTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeOpen}
var ClosedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeClosed}
var AcceptedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeAccepted}
var RejectedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeRejected}
var ApprovedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeApproved}
var SolvedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeSolved}
