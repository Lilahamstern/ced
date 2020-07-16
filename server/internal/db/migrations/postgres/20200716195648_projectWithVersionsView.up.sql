CREATE VIEW GET_PROJECTS_WITH_LATEST_VERSION_VIEW AS
SELECT p.id          AS project_id,
       p.createdat   AS project_createdat,
       p.updatedat   AS project_updatedat,
       v.id          AS version_id,
       v.orderid     AS version_order_id,
       v.title       AS version_title,
       v.description AS version_description,
       v.manager     AS version_manager,
       v.client      AS version_client,
       v.sector      AS version_sector,
       v.createdat   AS version_createdat,
       v.updatedat   AS version_updatedat
FROM (SELECT v.projectid, MAX(v.updatedat) AS updatedat
      FROM version v
      GROUP BY v.projectid
     ) v1
         INNER JOIN version v ON v1.projectid = v.projectid AND v1.updatedat = v.updatedat
         RIGHT JOIN project p ON v.projectid = p.id
ORDER BY v.updatedat DESC;