package e

import "errors"

var ErrorTicketNotFound = errors.New("ticket not found")
var ErrorTicketIsNotOpen = errors.New("ticket is not open")
var ErrorTicketIsNotAccepted = errors.New("ticket is not accepted")
var ErrorTicketIsNotClosed = errors.New("ticket is not closed")
