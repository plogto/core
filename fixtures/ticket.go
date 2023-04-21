package fixtures

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

var TicketID = pgtype.UUID{}
var EmptyTicket = &db.Ticket{}
var TicketWithID = &db.Ticket{ID: TicketID}

var TicketWithUserID = &db.Ticket{ID: TicketID, UserID: UserWithID.ID}
var OpenTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeOpen}
var ClosedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeClosed}
var AcceptedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeAccepted}
var RejectedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeRejected}
var ApprovedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeApproved}
var SolvedTicket = &db.Ticket{ID: TicketID, Status: db.TicketStatusTypeSolved}
