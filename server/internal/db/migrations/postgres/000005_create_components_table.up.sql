CREATE TABLE IF NOT EXISTS Components
(
    id        UUID         NOT NULL    DEFAULT uuid_generate_v4(),
    versionId UUID         NOT NULL,
    compId    VARCHAR(100) NOT NULL,
    name      VARCHAR(50)  NOT NULL,
    profile   VARCHAR(30)  NOT NULL,
    material  VARCHAR(40)  NOT NULL,
    co        DECIMAL      NOT NULL,
    level     INT          NOT NULL,
    type      VARCHAR(50)  NOT NULL,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updatedAt TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    FOREIGN KEY (versionId) REFERENCES Version (id) ON DELETE CASCADE,
    PRIMARY KEY (id)
)