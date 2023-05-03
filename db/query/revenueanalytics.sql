
-- name: GetClassRevenueGenerateByLocation :many
SELECT
    l.id AS location_id,
    SUM(c.cost) AS revenue
FROM
    class c
        JOIN schedule s ON c.scheduleid = s.id
        JOIN location l ON s.locationid = l.id
        JOIN classcatalogue cc ON c.id = cc.courseid
GROUP BY
    l.id;


-- name: GetRevenueGenerateByMemberships :many
SELECT
    SUM(mt.cost) AS total_revenue
FROM
    membership m
        JOIN membershiptypes mt ON m.membershipid = mt.id;