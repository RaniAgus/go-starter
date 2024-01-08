-- +goose Up
-- +goose StatementBegin
CREATE TABLE films (
  id       uuid         PRIMARY KEY DEFAULT gen_random_uuid(),
  title    varchar(255) NOT NULL,
  director varchar(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE films;
-- +goose StatementEnd
