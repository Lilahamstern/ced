-- +migrate Up
CREATE TABLE IF NOT EXISTS project (
    pid VARCHAR(50) NOT NULL,
    oid VARCHAR(50) NOT NULL unique,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    manager VARCHAR(50) NOT NULL,
    client VARCHAR(50) NOT NULL,
    sector VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (pid)
);

-- +migrate Down
DROP TABLE project;