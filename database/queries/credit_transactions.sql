-- name: CreateCreditTransaction :one
INSERT INTO
	credit_transactions (
		user_id,
		recipient_id,
		amount,
		credit_transaction_info_id,
		TYPE,
		url
	)
VALUES
	($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetCreditTransactionByID :one
SELECT
	*
FROM
	credit_transactions
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetCreditTransactionByUrl :one
SELECT
	*
FROM
	credit_transactions
WHERE
	url = $1
	AND deleted_at IS NULL;

-- name: GetCreditsByUserID :one
SELECT
	sum(credit_transaction.amount) AS amount
FROM
	credit_transactions AS credit_transaction
	INNER JOIN credit_transaction_infos ON credit_transaction_infos.id = credit_transaction.credit_transaction_info_id
WHERE
	credit_transaction.user_id = $1
	AND credit_transaction_infos.status = 'approved'
	AND credit_transaction.deleted_at IS NULL;

-- name: GetCreditTransactionsByUserIDAndPageInfo :many
SELECT
	*
FROM
	credit_transactions AS credit_transaction
	INNER JOIN credit_transaction_infos ON credit_transaction_infos.id = credit_transaction.credit_transaction_info_id
WHERE
	credit_transaction.user_id = $1
	AND credit_transaction_infos.created_at < $2
	AND credit_transaction.deleted_at IS NULL
ORDER BY
	credit_transaction_infos.created_at DESC
LIMIT
	$3;

-- name: CountCreditTransactionsByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		credit_transactions AS credit_transaction
		INNER JOIN credit_transaction_infos ON credit_transaction_infos.id = credit_transaction.credit_transaction_info_id
	WHERE
		credit_transaction.user_id = $1
		AND credit_transaction_infos.created_at < $2
		AND credit_transaction.deleted_at IS NULL
	GROUP BY
		credit_transaction.id,
		credit_transaction_infos.id,
		credit_transaction_infos.created_at
	ORDER BY
		credit_transaction_infos.created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;