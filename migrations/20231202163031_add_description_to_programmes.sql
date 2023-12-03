-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE academic_programmes ADD COLUMN description VARCHAR(300);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE academic_programmes DROP COLUMN description;
-- +goose StatementEnd
