CREATE TABLE IF NOT EXISTS VersionInformation
(
    id          UUID        NOT NULL DEFAULT uuid_generate_v4(),
    orderId     BIGINT      NOT NULL,
    title       VARCHAR(40) NOT NULL,
    description VARCHAR(200),
    manager     VARCHAR(50) NOT NULL,
    client      VARCHAR(50) NOT NULL,
    sector      VARCHAR(50) NOT NULL,
    createdAt   TIMESTAMP            default current_timestamp,
    updatedAt   TIMESTAMP   NULL     DEFAULT NULL,
    PRIMARY KEY (id)
)