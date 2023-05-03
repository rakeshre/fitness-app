


### Average enrollment by class type and day of the week:
 SELECT    day, classtype, AVG(enrollment) as avg_enrollmen FROM    class
    JOIN (
        SELECT
            classcatalogue.courseid,
            COUNT(*) as enrollment
        FROM
            classcatalogue
            JOIN class ON classcatalogue.courseid = class.id
        WHERE
            class.regstatus = 'Open'
        GROUP BY
            classcatalogue.courseid
    ) AS enrollment ON enrollment.courseid = class.id
GROUP BY
    day,
    classtype
ORDER BY
    day ASC; 

    
    
This query calculates the average enrollment for each class type and day of the week.


### Enrollment trend by week over the past 12 weeks:
SELECT
    date_trunc('week', startdate) AS week_start,
    COUNT(*) AS enrollment
FROM
    class
    JOIN classcatalogue ON class.id = classcatalogue.courseid
WHERE
    class.regstatus = 'Open' AND
    startdate >= date_trunc('week', now() - interval '12 weeks')
GROUP BY
    week_start
ORDER BY
    week_start ASC;
This query shows the enrollment trend for the past 12 weeks, grouping classes by the start of their week.

### Enrollment comparison by location and day of the week:
SELECT
    day,
    location.city,
    COUNT(*) AS enrollment
FROM
    class
    JOIN classcatalogue ON class.id = classcatalogue.courseid
    JOIN location ON class.locationid = location.id
WHERE
    class.regstatus = 'Open'
GROUP BY
    day,
    location.city
ORDER BY
    day ASC;

### Show the total revenue generated from class enrollments for the current month:

SELECT SUM(cost) FROM class
INNER JOIN classcatalogue ON class.id = classcatalogue.courseid
WHERE date_trunc('month', startdate) = date_trunc('month', now());


### Show the most popular class based on enrollment:

SELECT name, COUNT(*) as enrollments FROM classcatalogue
INNER JOIN class ON classcatalogue.courseid = class.id
GROUP BY class.name
ORDER BY enrollments DESC
LIMIT 1;



# By Hour GYM 

## Show the number of visitors by the hour for the current day:

SELECT extract(hour from checkin) as hour, COUNT(*) FROM checkinactivity
WHERE date_trunc('day', checkin) = date_trunc('day', now())
GROUP BY hour;

## Show the number of visitors by the hour for the current week, grouped by weekday:

SELECT extract(dow from checkin) as weekday, extract(hour from checkin) as hour, COUNT(*) FROM checkinactivity
WHERE checkin >= date_trunc('week', now()) AND checkin <= now()
GROUP BY weekday, hour;




