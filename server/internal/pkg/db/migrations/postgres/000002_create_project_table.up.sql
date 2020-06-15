CREATE TABLE IF NOT EXISTS Projects
(
    id        BIGINT    NOT NULL,
    createdAt TIMESTAMP      default current_timestamp,
    updatedAt TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (id)
)