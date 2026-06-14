CREATE TABLE urls (
    id BIGSERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(50) NOT NULL UNIQUE,
    click_count INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
/* TIMESTAMPTZ stores timezone-aware timestamps and is generally preferred in PostgreSQL. */
