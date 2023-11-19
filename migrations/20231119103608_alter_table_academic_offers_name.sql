-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE academic_offers ALTER COLUMN NAME TYPE VARCHAR(100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
