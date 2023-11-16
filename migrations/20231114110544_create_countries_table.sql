-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE countries (ID INTEGER NOT NULL, NAME VARCHAR(60) NOT NULL, PRIMARY KEY(ID));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
