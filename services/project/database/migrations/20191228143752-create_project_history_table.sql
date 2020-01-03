-- +migrate Up
CREATE TABLE IF NOT EXISTS project_history (
    pid VARCHAR(50) REFERENCES project(pid) ON DELETE CASCADE,
    oid VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    manager VARCHAR(50) NOT NULL,
    client VARCHAR(50) NOT NULL,
    sector VARCHAR(50) NOT NULL,
    version INTEGER NOT NULL,
    created_at timestamp default current_timestamp
);

-- +migrate Down
DROP TABLE project_history;