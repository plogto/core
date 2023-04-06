-- name: CreateInvitedUser :one
INSERT INTO
	invited_users (inviter_id, invitee_id)
VALUES
	($1, $2) RETURNING *;