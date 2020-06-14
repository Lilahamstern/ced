CREATE TABLE IF NOT EXISTS Versions
(
    id        UUID      NOT NULL DEFAULT uuid_generate_v4(),
    projectId BIGINT    NOT NULL,
    createdAt TIMESTAMP          default current_timestamp,
    updatedAt TIMESTAMP NULL,
    FOREIGN KEY (projectId) REFERENCES Projects (id) ON DELETE CASCADE,
    PRIMARY KEY (id)
)