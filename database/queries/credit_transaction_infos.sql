-- name: CreateCreditTransactionInfo :one
INSERT INTO
	credit_transaction_infos (
		description,
		credit_transaction_template_id,
		status
	)
VALUES
	($1, $2, $3) RETURNING *;

-- name: GetCreditTransactionInfoByID :one
SELECT
	*
FROM
	credit_transaction_infos
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: UpdateCreditTransactionInfoStatus :one
UPDATE
	credit_transaction_infos
SET
	status = $1
WHERE
	id = $2
	AND deleted_at IS NULL RETURNING *;