CREATE EXTENSION IF NOT EXISTS vector;

CREATE TYPE claim_verdict AS ENUM ('PROCESSING', 'REFUTED', 'INCONCLUSIVE', 'SUPPORTED');
CREATE TYPE run_status as ENUM('PENDING', 'RUNNING', 'SUCCESSFUL', 'FAILED');

CREATE TABLE IF NOT EXISTS claim (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    content TEXT NOT NULL,
    verdict claim_verdict NOT NULL DEFAULT 'PROCESSING'::claim_verdict,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS source (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    url TEXT UNIQUE NOT NULL,
    title TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    published_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS chunk (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    content TEXT NOT NULL,
    embedding vector(1024),
    source_id BIGINT NOT NULL REFERENCES source(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS run (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    claim_id BIGINT NOT NULL REFERENCES claim(id),
    status run_status NOT NULL DEFAULT 'PENDING'::run_status,
    model_id TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    finished_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS claim_chunk (
    claim_id BIGINT NOT NULL REFERENCES claim(id),
    chunk_id BIGINT NOT NULL REFERENCES chunk(id),

    PRIMARY KEY (claim_id, chunk_id)
);

DROP INDEX IF EXISTS idx_claim_verdict_id;
DROP INDEX IF EXISTS idx_run_status_id;
DROP INDEX IF EXISTS idx_claim_chunk_chunk_id;

CREATE INDEX idx_claim_verdict_id ON claim(verdict, id);
CREATE INDEX idx_run_status_id ON run(status, id);
CREATE INDEX idx_claim_chunk_chunk_id ON claim_chunk(chunk_id);