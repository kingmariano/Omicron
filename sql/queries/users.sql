-- name: CreateUser :one
INSERT INTO UnvalidatedUsers(id, created_at,whatsapp_number,display_name)
VALUES (
$1, 
$2, 
$3, 
$4
)
RETURNING *;
--

-- name: GetUserWhatsappNumber :one
SELECT * FROM UnvalidatedUsers 
WHERE whatsapp_number = $1;