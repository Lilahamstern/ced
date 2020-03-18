CREATE TYPE project_return as (
	order_id integer,
	name varchar,
	description varchar,
	manager varchar,
	client varchar,
	sector varchar,
	project_id integer,
	created_at timestamp with time zone,
	updated_at timestamp with time zone
);

CREATE OR REPLACE FUNCTION get_projects(query varchar)
RETURNS SETOF project_return
AS
$BODY$

select order_id, name, description, manager, client, sector, project_id, created_at, updated_at
from project_information 
where 
CAST(order_id as TEXT) LIKE '%'|| query || '%'
OR name LIKE '%'|| query || '%'
OR manager LIKE '%'|| query || '%'
OR client LIKE '%'|| query || '%'
OR sector LIKE '%'|| query || '%'
ORDER BY updated_at DESC;

$BODY$ LANGUAGE sql;