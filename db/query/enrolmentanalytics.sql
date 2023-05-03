-- name: GetClassesOfferedAndAttendees :many
SELECT
    s."day",
    COUNT(c.id) AS classes_offered,
    COUNT(cc.id) AS attendees
FROM
    class c
        JOIN schedule s ON c.scheduleid = s.id
        LEFT JOIN classcatalogue cc ON c.id = cc.courseid
where s.locationid =$1
GROUP BY
    s."day"
ORDER BY
    s."day";

-- name: GetAllClassesOfferedAndAttendees :many
SELECT
    s."day",
    COUNT(c.id) AS classes_offered,
    COUNT(cc.id) AS attendees
FROM
    class c
        JOIN schedule s ON c.scheduleid = s.id
        LEFT JOIN classcatalogue cc ON c.id = cc.courseid
GROUP BY
    s."day"
ORDER BY
    s."day";

-- name: GetAllClassesOfferedAndAttendeesPerWeek :many
SELECT
    DATE_TRUNC('week', s.startdate),
    COUNT(c.id) AS classes_offered,
    COUNT(cc.id) AS attendees
FROM
    class c
        JOIN schedule s ON c.scheduleid = s.id
        LEFT JOIN classcatalogue cc ON c.id = cc.courseid
GROUP BY
    DATE_TRUNC('week', s.startdate)
ORDER BY
    DATE_TRUNC('week', s.startdate);

-- name: GetClassesOfferedAndAttendeesPerWeek :many
SELECT
    DATE_TRUNC('week', s.startdate) ,
    COUNT(c.id) AS classes_offered,
    COUNT(cc.id) AS attendees
FROM
    class c
        JOIN schedule s ON c.scheduleid = s.id
        LEFT JOIN classcatalogue cc ON c.id = cc.courseid
where s.locationid =$1
GROUP BY
    DATE_TRUNC('week', s.startdate)
ORDER BY
    DATE_TRUNC('week', s.startdate);

-- name: GetAllTopAttendedClass :many
SELECT
    c.id AS class_id,
    c.name AS class_name,
    COUNT(cc.id) AS attendees
FROM
    class c
        JOIN classcatalogue cc ON c.id = cc.courseid
GROUP BY
    c.id
ORDER BY
    attendees DESC;

-- name: GetMostPopularHourForClassesOnWeekdays :many
SELECT
    EXTRACT(HOUR FROM s.starttime),
    COUNT(cc.id) AS attendees
FROM
    class c
    JOIN schedule s ON c.scheduleid = s.id
    JOIN classcatalogue cc ON c.id = cc.courseid
WHERE
    EXTRACT(DOW FROM s.startdate) BETWEEN 1 AND 5
GROUP BY
    EXTRACT(HOUR FROM s.starttime)
ORDER BY
    attendees DESC;

-- name: GetMostPopularHourForClassesOnWeekends :many
SELECT
    EXTRACT(HOUR FROM s.starttime) ,
    COUNT(cc.id) AS attendees
FROM
    class c
    JOIN schedule s ON c.scheduleid = s.id
    JOIN classcatalogue cc ON c.id = cc.courseid
WHERE
    EXTRACT(DOW FROM s.startdate) IN (0, 6)
GROUP BY
    EXTRACT(HOUR FROM s.starttime)
ORDER BY
    attendees DESC;