CREATE OR REPLACE PROCEDURE createProjectWithVersion(BIGINT, BIGINT, VARCHAR(40), VARCHAR(200), VARCHAR(50),
                                                     VARCHAR(50), VARCHAR(50))
    LANGUAGE plpgsql
AS
$$
BEGIN
    INSERT INTO project(id)
    VALUES ($1);

    INSERT INTO version(projectid, orderid, title, description, manager, client, sector)
    VALUES ($1, $2, $3, $4, $5, $6, $7);

    COMMIT;
END;
$$;