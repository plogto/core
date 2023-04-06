-- name: CreateCreditTransactionDescriptionVariable :one
INSERT INTO
	credit_transaction_description_variables (credit_transaction_info_id, content_id, KEY, TYPE)
VALUES
	($1, $2, $3, $4) RETURNING *;

-- name: GetCreditTransactionDescriptionVariableByContentID :one
SELECT
	*
FROM
	credit_transaction_description_variables
WHERE
	content_id = $1
	AND deleted_at IS NULL
ORDER BY
	created_at DESC
LIMIT
	1;

-- name: GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID :many
SELECT
	*
FROM
	credit_transaction_description_variables
WHERE
	credit_transaction_info_id = $1
	AND deleted_at IS NULL;