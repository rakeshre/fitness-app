-- name: GetMembershipCountsByType :many
SELECT
    mt.member_type,
    COUNT(m.membershipid) AS total_members
FROM
    membership m
        JOIN membershiptypes mt ON m.membershipid = mt.id
GROUP BY
    mt.member_type;

-- name: GetDailyNewMemberEnrolments :many
SELECT
    DATE_TRUNC('day', m.startdate) AS day,
    COUNT(m.userid) AS new_members
FROM
    membership m
GROUP BY
    day
ORDER BY
    day;

-- name: GetKMostFrequentMembers :many
SELECT
    ua.userid,
    u.name,
    COUNT(ua.id) AS total_visits
FROM
    useractivity ua
        JOIN users u ON ua.userid = u.id
GROUP BY
    ua.userid, u.name
ORDER BY
    total_visits DESC
    limit $1;