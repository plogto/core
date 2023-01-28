-- name: GetCreditTransactionTemplateByID :one
SELECT
	*
FROM
	credit_transaction_templates
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetCreditTransactionTemplateByName :one
SELECT
	*
FROM
	credit_transaction_templates
WHERE
	NAME = $1
	AND deleted_at IS NULL;