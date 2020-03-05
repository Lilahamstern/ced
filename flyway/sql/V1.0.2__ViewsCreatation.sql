CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `root`@`%` 
    SQL SECURITY DEFINER
VIEW `advancedProjectInformation` AS
    SELECT 
        `v`.`Id` AS `VersionId`,
        `pi`.`ProjectId` AS `ProjectId`,
        `pi`.`OrderId` AS `OrderId`,
        `pi`.`Name` AS `Name`,
        `pi`.`Description` AS `Description`,
        `pi`.`Manager` AS `Manager`,
        `pi`.`Client` AS `Client`,
        `pi`.`Sector` AS `Sector`,
        `pi`.`CreatedAt` AS `CreatedAt`,
        ROUND(SUM(`c`.`Co`), 2) AS `Co`
    FROM
        ((`version` `v`
        LEFT JOIN `component` `c` ON ((`v`.`Id` = `c`.`VersionId`)))
        LEFT JOIN `projectInformation` `pi` ON ((`v`.`ProjectInformationId` = `pi`.`Id`)))
    GROUP BY `v`.`Id`