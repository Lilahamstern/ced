CREATE TABLE IF NOT EXISTS Project
(
    id        BIGINT NOT NULL,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updatedAt TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    PRIMARY KEY (id)
)