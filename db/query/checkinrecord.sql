-- name: CreateCheckinRecords :one
insert INTO checkinrecords
(type,userid,employeeid,locationid )
VALUES
    ($1,$2,$3,$4) RETURNING *;


-- name: GetLatestCheckinRecord :one
SELECT * FROM checkinrecords
WHERE userid = $1
ORDER BY time desc
LIMIT 1;



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