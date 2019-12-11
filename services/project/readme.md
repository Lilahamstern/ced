docker exec -it container-id ./cockroach sql --insecure

CREATE DATABASE project;
CREATE DATABASE component;

CREATE USER project_admin WITH PASSWORD 'supersecret';
CREATE USER component_admin WITH PASSWORD 'supersecret';

GRANT ALL ON DATABASE project TO project_admin;
GRANT ALL ON DATABASE component TO component_admin;
