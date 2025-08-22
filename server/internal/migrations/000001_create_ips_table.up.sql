CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE ips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    origin_ip VARCHAR NOT NULL,
    origin_reversed_ip VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP   
);