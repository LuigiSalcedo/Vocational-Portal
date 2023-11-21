-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE programmes_areas RENAME COLUMN programme_id TO programm_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
