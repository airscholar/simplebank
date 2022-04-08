-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1
LIMIT 1;

-- name: GetTransfersByFromId :many
SELECT * FROM transfers
WHERE from_account_id = $1
LIMIT $2 OFFSET $3;

-- name: GetTransfersByToId :many
SELECT * FROM transfers
WHERE to_account_id = $1
LIMIT $2 OFFSET $3;

-- name: GetTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: CreateTransfer :one
INSERT INTO transfers(
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateTransferByFromId :one
UPDATE transfers
set to_account_id = $2,
    amount = $3
WHERE from_account_id = $1
RETURNING *;

-- name: UpdateTransferByToId :one
UPDATE transfers
set from_account_id = $2,
    amount = $3
WHERE to_account_id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;

