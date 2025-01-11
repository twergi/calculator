-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS results (
    id      BIGINT  PRIMARY KEY NOT NULL,
    result  BIGINT  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS results;
-- +goose StatementEnd
