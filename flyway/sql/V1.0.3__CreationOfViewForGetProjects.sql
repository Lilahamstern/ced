CREATE TYPE project_return as (
	project_id integer,
	order_id integer,
	name varchar,
	description varchar,
	manager varchar,
	client varchar,
	sector varchar,
	created_at timestamp with time zone,
	updated_at timestamp with time zone
);

CREATE OR REPLACE FUNCTION get_projects(query varchar)
RETURNS SETOF project_return
AS
$BODY$

select DISTINCT ON (project_id) t1.project_id, t1.order_id, t1.name, t1.description, t1.manager, t1.client, t1.sector, t1.created_at
from project_information t1
INNER JOIN
(SELECT project_id, MAX(created_at) AS MaxDateTime
FROM project_information
GROUP by project_id
ORDER BY project_id DESC) as t2
ON t1.project_id = t2.project_id
and t1.created_at = t2.MaxDateTime
where 
CAST(order_id as TEXT) LIKE '%'|| query || '%'
OR name LIKE '%'|| query || '%'
OR manager LIKE '%'|| query || '%'
OR client LIKE '%'|| query || '%'
OR sector LIKE '%'|| query || '%';

$BODY$ LANGUAGE sql;