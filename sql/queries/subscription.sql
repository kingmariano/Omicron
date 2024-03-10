-- name: CreateUserSubscription :exec
INSERT INTO Subscription (subscription_id, userid, expiry_date)
VALUES($1, $2, $3);
--

-- name: GetUserSubscription :one
SELECT * FROM Subscription
WHERE userid = $1;
--