-- name: CreateUnRegUsers :one
INSERT INTO Unregisteredusers(created_at, updated_at,whatsapp_number,display_name)
VALUES (
$1, 
$2, 
$3, 
$4
)
RETURNING *;

-- name: GetUnRegUsers :one
SELECT * FROM Unregisteredusers
WHERE whatsapp_number = $1;
