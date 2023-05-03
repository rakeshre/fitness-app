-- name: CreateLocation :one
INSERT INTO location (
    city,
    state,
    zipcode
) VALUES (
             $1, $2, $3
         ) RETURNING *;

-- name: GetLocation :one
SELECT * FROM location
WHERE id = $1 LIMIT 1;

-- name: GetAllLocations :many
SELECT * FROM location;


-- -- name: UpdateUser :one
-- UPDATE users
-- SET
--     hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
--     password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
--     full_name = COALESCE(sqlc.narg(full_name), full_name),
--     email = COALESCE(sqlc.narg(email), email)
-- WHERE
--         username = sqlc.arg(username)
--     RETURNING *;