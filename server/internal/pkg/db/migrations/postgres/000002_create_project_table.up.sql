CREATE TABLE IF NOT EXISTS Projects
(
    id        BIGINT NOT NULL,
    createdAt TIMESTAMP DEFAULT current_timestamp,
    updatedAt TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY (id)
)