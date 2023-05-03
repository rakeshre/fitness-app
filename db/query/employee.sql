-- name: CreateEmployee :one
INSERT INTO employee (
    name,
    email,
    hashedpassword,
    locationid
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: GetEmployee :one
SELECT * FROM employee
WHERE name = $1 LIMIT 1;

-- name: GetEmployeeFromEmail :one
SELECT * FROM employee
WHERE email = $1 LIMIT 1;

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