-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE universities ADD COLUMN URL VARCHAR(100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE universities DROP COLUMN URL;
-- +goose StatementEnd
