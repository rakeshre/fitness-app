-- name: CreateSchedule :one
INSERT INTO schedule (
starttime,
endtime,
startdate,
enddate,
locationid,
day
)
VALUES
($1, $2, $3,$4, $5, $6) RETURNING *;


-- name: GetSchedule :one
SELECT * FROM schedule
WHERE id = $1 LIMIT 1;

-- name: GetAllSchedules :many
SELECT * FROM schedule;