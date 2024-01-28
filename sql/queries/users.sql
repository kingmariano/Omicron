-- name: CreateUnRegUser :one
INSERT INTO Unregisteredusers(created_at, updated_at,whatsapp_number,display_name)
VALUES (
$1, 
$2, 
$3, 
$4
)
RETURNING *;

-- name: GetUnRegUser :one
SELECT * FROM Unregisteredusers
WHERE whatsapp_number = $1;

-- name: GetUnRegUsers :many
SELECT * FROM Unregisteredusers;
