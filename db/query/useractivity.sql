-- name: CreateUserActivity :one
INSERT INTO useractivity
("start" ,"end" ,userid,deviceid,locationid)
VALUES
    ($1,$2,$3,$4,$5) RETURNING *;


-- name: GetUserActivity :many
SELECT * FROM useractivity
WHERE userid = $1;

-- name: GetPastWorkoutData1 :many
SELECT d.description AS devicetype, SUM(EXTRACT(EPOCH FROM (ua.end - ua.start))) AS totaltimeseconds
FROM useractivity ua
         JOIN device d ON ua.deviceid = d.id
WHERE ua.userid = $1 AND ua.start >= NOW() - INTERVAL '1 days'
GROUP BY  d.description
ORDER BY d.description;

-- name: GetPastWorkoutData7 :many
SELECT d.description AS devicetype, SUM(EXTRACT(EPOCH FROM (ua.end - ua.start))) AS totaltimeseconds
FROM useractivity ua
         JOIN device d ON ua.deviceid = d.id
WHERE ua.userid = $1 AND ua.start >= NOW() - INTERVAL '7 days'
GROUP BY  d.description
ORDER BY d.description;

-- name: GetPastWorkoutData30 :many
SELECT d.description AS devicetype, SUM(EXTRACT(EPOCH FROM (ua.end - ua.start))) AS totaltimeseconds
FROM useractivity ua
         JOIN device d ON ua.deviceid = d.id
WHERE ua.userid = $1 AND ua.start >= NOW() - INTERVAL '30 days'
GROUP BY  d.description
ORDER BY d.description;

-- name: GetPastWorkoutData60 :many
SELECT d.description AS devicetype, SUM(EXTRACT(EPOCH FROM (ua.end - ua.start))) AS totaltimeseconds
FROM useractivity ua
         JOIN device d ON ua.deviceid = d.id
WHERE ua.userid = $1 AND ua.start >= NOW() - INTERVAL '60 days'
GROUP BY  d.description
ORDER BY d.description;

-- name: GetPastWorkoutData90 :many
SELECT d.description AS devicetype, SUM(EXTRACT(EPOCH FROM (ua.end - ua.start))) AS totaltimeseconds
FROM useractivity ua
         JOIN device d ON ua.deviceid = d.id
WHERE ua.userid = $1 AND ua.start >= NOW() - INTERVAL '90 days'
GROUP BY  d.description
ORDER BY d.description;

-- name: GetDayWiseActivity :many
SELECT
    DATE(ua.start) AS date,
    SUM(EXTRACT(EPOCH FROM (ua.end - ua.start))) AS total_time_seconds
FROM
    useractivity ua
WHERE
    ua.userid = $1
  AND EXTRACT(YEAR FROM ua.start) = EXTRACT(YEAR FROM NOW())
GROUP BY
    DATE(ua.start)
ORDER BY
    DATE(ua.start);


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