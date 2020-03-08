CREATE TABLE `ced`.`project` (
  `Id` INT NOT NULL,
  `CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`Id`));

  

  CREATE TABLE `ced`.`projectInformation` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `OrderId` INT NOT NULL,
  `Name` VARCHAR(40) NOT NULL,
  `Description` VARCHAR(300) NULL,
  `Manager` VARCHAR(40) NOT NULL,
  `Client` VARCHAR(45) NOT NULL,
  `Sector` VARCHAR(45) NOT NULL,
  `ProjectId` INT NOT NULL,
  `CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`Id`));


 CREATE TABLE `ced`.`version` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Title` VARCHAR(45) NOT NULL,
  `Description` VARCHAR(300) NULL,
  `ProjectId` INT NOT NULL,
  `ProjectInformationId` INT NOT NULL,
  `CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`Id`));

  CREATE TABLE `ced`.`component` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `ComponentId` INT NOT NULL,
  `Name` VARCHAR(60) NOT NULL,
  `Profile` VARCHAR(30) NOT NULL,
  `Material` VARCHAR(50) NOT NULL,
  `Co` FLOAT NOT NULL,
  `Level` INT NOT NULL,
  `Type` VARCHAR(45) NOT NULL,
  `VersionId` INT NOT NULL,
  `CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`Id`));

  ALTER TABLE `ced`.`projectInformation` 
ADD INDEX `FK_Project_Id_ProjectInformation_id_idx` (`ProjectId` ASC) VISIBLE;
;
ALTER TABLE `ced`.`projectInformation` 
ADD CONSTRAINT `FK_Project_Id_ProjectInformation_id`
  FOREIGN KEY (`ProjectId`)
  REFERENCES `ced`.`project` (`Id`)
  ON DELETE CASCADE
  ON UPDATE RESTRICT;

  ALTER TABLE `ced`.`version` 
ADD INDEX `FK_Project_Id_Version_ProjectId_idx` (`ProjectId` ASC) VISIBLE,
ADD INDEX `FK_ProjectInformation_Id_Version_ProjectInfromationId_idx` (`ProjectInformationId` ASC) VISIBLE;
;
ALTER TABLE `ced`.`version` 
ADD CONSTRAINT `FK_Project_Id_Version_ProjectId`
  FOREIGN KEY (`ProjectId`)
  REFERENCES `ced`.`project` (`Id`)
  ON DELETE NO ACTION
  ON UPDATE RESTRICT,
ADD CONSTRAINT `FK_ProjectInformation_Id_Version_ProjectInfromationId`
  FOREIGN KEY (`ProjectInformationId`)
  REFERENCES `ced`.`projectInformation` (`Id`)
  ON DELETE CASCADE
  ON UPDATE RESTRICT;

  ALTER TABLE `ced`.`component` 
ADD INDEX `FK_version_Id_Component_VersionId_idx` (`VersionId` ASC) VISIBLE;
;
ALTER TABLE `ced`.`component` 
ADD CONSTRAINT `FK_version_Id_Component_VersionId`
  FOREIGN KEY (`VersionId`)
  REFERENCES `ced`.`version` (`Id`)
  ON DELETE CASCADE
  ON UPDATE RESTRICT;


