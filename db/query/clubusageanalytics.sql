
-- name: GetHoursSpentInGymByDay :many
SELECT DATE_TRUNC('day', ci.checkin) AS day,
SUM(EXTRACT(EPOCH FROM (ci.checkout - ci.checkin))/3600) AS hours_spent
FROM checkinactivity ci
GROUP BY day
ORDER BY day;

-- name: GetBusiestTimeByHourAndDayOfWeek :many
SELECT
    EXTRACT(DOW FROM ci.checkin) AS day_of_week,
    EXTRACT(HOUR FROM ci.checkin) AS hour,
    COUNT(*) AS visits
FROM
    checkinactivity ci
GROUP BY
    day_of_week, hour
ORDER BY
    day_of_week, hour;

-- name: GetAverageVisitorsPerHourWeekdays :many
SELECT
    EXTRACT(HOUR FROM ci.checkin) AS hour,
    COUNT(*)/COUNT(DISTINCT DATE_TRUNC('day', ci.checkin))::float AS avg_visitors
FROM
    checkinactivity ci
WHERE
    EXTRACT(DOW FROM ci.checkin) BETWEEN 1 AND 5
GROUP BY
    hour
ORDER BY
    hour;

-- name: GetAverageVisitorsPerHourWeekends :many
SELECT
    EXTRACT(HOUR FROM ci.checkin) AS hour,
    COUNT(*)/COUNT(DISTINCT DATE_TRUNC('day', ci.checkin))::float AS avg_visitors
FROM
    checkinactivity ci
WHERE
    EXTRACT(DOW FROM ci.checkin) IN (0, 6)
GROUP BY
    hour
ORDER BY
    hour;