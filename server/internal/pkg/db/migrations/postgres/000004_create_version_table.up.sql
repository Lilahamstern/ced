CREATE TABLE IF NOT EXISTS Versions
(
    id            UUID   NOT NULL DEFAULT uuid_generate_v4(),
    projectId     BIGINT NOT NULL,
    informationId UUID   NULL,
    createdAt     TIMESTAMP       default current_timestamp,
    updatedAt     TIMESTAMP       DEFAULT current_timestamp,
    FOREIGN KEY (projectId) REFERENCES Projects (id) ON DELETE CASCADE,
    FOREIGN KEY (informationId) REFERENCES VersionInformation (id)
        ON DELETE SET NULL
        ON UPDATE RESTRICT,
    PRIMARY KEY (id)
)