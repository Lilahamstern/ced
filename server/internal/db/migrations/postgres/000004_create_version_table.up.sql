CREATE TABLE IF NOT EXISTS Version
(
    id            UUID   NOT NULL          DEFAULT uuid_generate_v4(),
    projectId     BIGINT NOT NULL,
    informationId UUID   NULL,
    createdAt     TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updatedAt     TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    FOREIGN KEY (projectId) REFERENCES Project (id) ON DELETE CASCADE,
    FOREIGN KEY (informationId) REFERENCES Information (id)
        ON DELETE SET NULL
        ON UPDATE RESTRICT,
    PRIMARY KEY (id)
)