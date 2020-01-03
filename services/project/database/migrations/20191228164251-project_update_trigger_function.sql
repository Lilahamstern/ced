
-- +migrate Up

-- CREATE OR REPLACE FUNCTION update_project() 
--     RETURNS trigger 
--     SET SCHEMA 'public'
--     LANGUAGE plpgsql
--     AS '
--     BEGIN
--     INSERT INTO
--         project_history(pid, oid, name, description, manager, client, sector, version)
--     VALUES
--         (OLD.pid, OLD.oid, OLD.name, OLD.description, OLD.manager, OLD.client, OLD.sector, 1);
--     RETURN OLD;
--     END;
--     ';
--DROP FUNCTION IF EXISTS update_project;


CREATE TRIGGER on_project_update
AFTER
UPDATE
    ON project FOR EACH ROW EXECUTE PROCEDURE update_project();


-- +migrate Down
DROP TRIGGER IF EXISTS on_project_update ON project;