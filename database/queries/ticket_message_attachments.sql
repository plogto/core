-- name: CreateTicketMessageAttachment :one
INSERT INTO
	ticket_message_attachments (ticket_message_id, file_id)
VALUES
	($1, $2) RETURNING *;