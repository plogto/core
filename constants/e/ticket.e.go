package e

import "errors"

var ErrorTicketNotFound = errors.New("ticket not found")
var ErrorTicketIsNotOpen = errors.New("ticket is not open")
var ErrorTicketIsNotApproved = errors.New("ticket is not approved")
var ErrorTicketIsNotClosed = errors.New("ticket is not closed")
