-- +goose Up
-- +goose StatementBegin
CREATE TABLE versions (
    id BIGSERIAL PRIMARY KEY,
    version VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE versions;
-- +goose StatementEnd
