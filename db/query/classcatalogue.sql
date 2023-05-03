-- name: CreateClassCatalogue :one
INSERT INTO classcatalogue (courseid,userid)
VALUES ($1,$2) RETURNING *;


-- name: GetUserClass :many
SELECT * FROM classcatalogue
WHERE userid = $1;

-- name: GetClassEnrolment :many
SELECT userid FROM classcatalogue
WHERE courseid = $1;

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