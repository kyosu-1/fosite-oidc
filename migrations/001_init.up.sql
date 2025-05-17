CREATE TABLE IF NOT EXISTS oauth2_clients (
    id TEXT PRIMARY KEY,
    secret TEXT NOT NULL,
    redirect_uris TEXT[] NOT NULL,
    scopes TEXT[] NOT NULL
);

CREATE TABLE IF NOT EXISTS oauth2_access_tokens (
    signature TEXT PRIMARY KEY,
    requester BYTEA NOT NULL,
    requested_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS oauth2_refresh_tokens (
    signature TEXT PRIMARY KEY,
    requester BYTEA NOT NULL,
    requested_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

