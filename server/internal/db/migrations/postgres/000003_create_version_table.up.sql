CREATE TABLE IF NOT EXISTS Version
(
    id          UUID        NOT NULL     DEFAULT uuid_generate_v4(),
    projectId   BIGINT      NOT NULL,
    orderId     BIGINT      NOT NULL,
    title       VARCHAR(40) NOT NULL,
    description VARCHAR(200),
    manager     VARCHAR(50) NOT NULL,
    client      VARCHAR(50) NOT NULL,
    sector      VARCHAR(50) NOT NULL,
    createdAt   TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updatedAt   TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    FOREIGN KEY (projectId) REFERENCES Project (id) ON DELETE CASCADE,
    PRIMARY KEY (id)
)