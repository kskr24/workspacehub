-- +goose Up
-- +goose StatementBegin
-- sessions table
CREATE TABLE IF NOT EXISTS sessions
(
    id      BIGSERIAL PRIMARY KEY,
    ua      VARCHAR(255) NOT NULL DEFAULT '', -- User agent information
    ip      VARCHAR(255) NOT NULL DEFAULT '', -- IP address
    user_id BIGINT       NOT NULL,            -- Foreign key to user table
    token   VARCHAR(255) NOT NULL DEFAULT '', -- Session token
    created BIGINT       NOT NULL DEFAULT 0,  -- Creation timestamp
    expires BIGINT       NOT NULL DEFAULT 0,  -- Expiration timestamp
    updated BIGINT       NOT NULL DEFAULT 0   -- Last updated timestamp
);

CREATE INDEX IF NOT EXISTS idx_sessions_token
    ON sessions (token);

CREATE INDEX IF NOT EXISTS idx_sessions_user
    ON sessions (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
