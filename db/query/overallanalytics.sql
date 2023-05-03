

-- name: GetPreviousWeekCheckinCount :one
SELECT
    COUNT(*) AS checkin_count
FROM
    checkinactivity c
WHERE
        c.checkin  >= DATE_TRUNC('week', NOW()) - INTERVAL '1 week'
  AND c.checkin  < DATE_TRUNC('week', NOW()) ;



-- name: GetCurrentWeekCheckinCount :one
SELECT
    COUNT(*) AS checkin_count
FROM
    checkinactivity c
WHERE
        c.checkin  >= DATE_TRUNC('week', NOW());



-- name: GetPreviousWeekMembershipsCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    membership m
WHERE
        m.startdate  >= DATE_TRUNC('week', NOW()) - INTERVAL '1 week'
  AND m.startdate  < DATE_TRUNC('week', NOW());




-- name: GetCurrentWeekMembershipsCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    membership m
WHERE
        m.startdate  >= DATE_TRUNC('week', NOW())
  AND m.startdate  < DATE_TRUNC('week', NOW()) + INTERVAL '1 week';



-- name: GetPreviousWeekClassEnrolmentCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    classcatalogue cc
        JOIN class c ON cc.courseid = c.id
WHERE
        cc.enrolmentdate >= DATE_TRUNC('week', NOW()) - INTERVAL '1 week'
  AND cc.enrolmentdate < DATE_TRUNC('week', NOW());


-- name: GetCurrentWeekClassEnrolmentCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    classcatalogue cc
        JOIN class c ON cc.courseid = c.id
WHERE
        cc.enrolmentdate >= DATE_TRUNC('week', NOW())
  AND cc.enrolmentdate < DATE_TRUNC('week', NOW()) + INTERVAL '1 week';





-- name: GetPreviousMonthCheckinCount :one
SELECT
    COUNT(*) AS checkin_count
FROM
    checkinactivity c
WHERE
        c.checkin  >= DATE_TRUNC('month', NOW()) - INTERVAL '1 month'
  AND c.checkin  < DATE_TRUNC('month', NOW()) ;



-- name: GetCurrentMonthCheckinCount :one
SELECT
    COUNT(*) AS checkin_count
FROM
    checkinactivity c
WHERE
        c.checkin  >= DATE_TRUNC('month', NOW());



-- name: GetPreviousMonthMembershipsCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    membership m
WHERE
        m.startdate  >= DATE_TRUNC('month', NOW()) - INTERVAL '1 month'
  AND m.startdate  < DATE_TRUNC('month', NOW());




-- name: GetCurrentMonthMembershipsCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    membership m
WHERE
        m.startdate  >= DATE_TRUNC('month', NOW())
  AND m.startdate  < DATE_TRUNC('month', NOW()) + INTERVAL '1 month';



-- name: GetPreviousMonthClassEnrolmentCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    classcatalogue cc
        JOIN class c ON cc.courseid = c.id
WHERE
        cc.enrolmentdate >= DATE_TRUNC('month', NOW()) - INTERVAL '1 month'
  AND cc.enrolmentdate < DATE_TRUNC('month', NOW());


-- name: GetCurrentMonthClassEnrolmentCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    classcatalogue cc
        JOIN class c ON cc.courseid = c.id
WHERE
        cc.enrolmentdate >= DATE_TRUNC('month', NOW())
  AND cc.enrolmentdate < DATE_TRUNC('month', NOW()) + INTERVAL '1 month';









-- name: GetPreviousDayCheckinCount :one
SELECT
    COUNT(*) AS checkin_count
FROM
    checkinactivity c
WHERE
        c.checkin  >= DATE_TRUNC('day', NOW()) - INTERVAL '1 day'
  AND c.checkin  < DATE_TRUNC('day', NOW()) ;



-- name: GetCurrentDayCheckinCount :one
SELECT
    COUNT(*) AS checkin_count
FROM
    checkinactivity c
WHERE
        c.checkin  >= DATE_TRUNC('day', NOW());



-- name: GetPreviousDayMembershipsCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    membership m
WHERE
        m.startdate  >= DATE_TRUNC('day', NOW()) - INTERVAL '1 day'
  AND m.startdate  < DATE_TRUNC('day', NOW());




-- name: GetCurrentDayMembershipsCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    membership m
WHERE
        m.startdate  >= DATE_TRUNC('day', NOW())
  AND m.startdate  < DATE_TRUNC('day', NOW()) + INTERVAL '1 day';



-- name: GetPreviousDayClassEnrolmentCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    classcatalogue cc
        JOIN class c ON cc.courseid = c.id
WHERE
        cc.enrolmentdate >= DATE_TRUNC('day', NOW()) - INTERVAL '1 day'
  AND cc.enrolmentdate < DATE_TRUNC('day', NOW());


-- name: GetCurrentDayClassEnrolmentCount :one
SELECT
    COUNT(*) AS new_member_count
FROM
    classcatalogue cc
        JOIN class c ON cc.courseid = c.id
WHERE
        cc.enrolmentdate >= DATE_TRUNC('day', NOW())
  AND cc.enrolmentdate < DATE_TRUNC('day', NOW()) + INTERVAL '1 day';



