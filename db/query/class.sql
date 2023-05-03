-- name: CreateClass :one
INSERT INTO class (
    instructorname,
    name,
    cost,
    scheduleid
)
VALUES
    ($1, $2, $3,$4) RETURNING *;


-- name: GetClass :one
SELECT * FROM class
WHERE id = $1 LIMIT 1;

-- name: GetAllClasses :one
SELECT * FROM class;


-- name: GetClasses :many
SELECT
    c.id AS class_id,
    c.name AS class_name,
    c.instructorname,
    c.cost,
    s.startdate,
    s.enddate,
    s.starttime,
    s.endtime,
    CASE
        WHEN cc.id IS NULL THEN 'Not Enrolled'
        ELSE 'Enrolled'
        END AS enrollment_status
FROM
    class c
        JOIN schedule s ON c.scheduleid = s.id
        LEFT JOIN classcatalogue cc ON c.id = cc.courseid AND cc.userid = $1
WHERE
        s.locationid = $2
  AND s.day = $3
ORDER BY
    s.starttime;


-- name: GetClassesForEmployee :many
SELECT
    c.id AS class_id,
    c.name AS class_name,
    c.instructorname,
    c.cost,
    s.startdate,
    s.enddate,
    s.starttime,
    s.endtime
from class c
         JOIN schedule s ON c.scheduleid = s.id
WHERE
        s.locationid = $2 AND s.day = $1
ORDER BY
    s.starttime;



-- name: GetUpcomingClasses :many
WITH class_dates AS (
    SELECT
        c.id AS class_id,
        c.name AS class_name,
        s.starttime,
        s.endtime,
        s.day,
        GENERATE_SERIES(s.startdate, LEAST(s.enddate, NOW() + INTERVAL '7 days'), INTERVAL '1 day') AS class_date
    FROM
        classcatalogue cc
            JOIN class c ON cc.courseid = c.id
            JOIN schedule s ON c.scheduleid = s.id
    WHERE
            cc.userid = $1
)
SELECT
    class_id,
    class_name,
    class_date,
    starttime,
    endtime
FROM
    class_dates c
where
        TRIM(BOTH FROM TO_CHAR(class_date, 'Day')) = c.day and
    class_date BETWEEN NOW() AND NOW() + INTERVAL '7 days'
ORDER BY
    class_date, class_id;





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