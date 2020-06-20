CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Log queries that takes longer then 300ms to execute
SET log_min_duration_statement = '300';
