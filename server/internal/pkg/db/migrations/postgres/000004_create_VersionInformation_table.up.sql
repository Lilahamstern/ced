CREATE TABLE IF NOT EXISTS VersionInformation
(
    versionId   UUID        NOT NULL,
    orderId     BIGINT      NOT NULL,
    title       VARCHAR(40) NOT NULL,
    description VARCHAR(200),
    manager     VARCHAR(50) NOT NULL,
    client      VARCHAR(50) NOT NULL,
    sector      VARCHAR(50) NOT NULL,
    createdAt   TIMESTAMP default current_timestamp,
    updatedAt   TIMESTAMP   NULL,
    FOREIGN KEY (versionId) REFERENCES Versions (id)
        ON DELETE CASCADE
        ON UPDATE RESTRICT,
    PRIMARY KEY (versionId)
)