-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE programmes_areas RENAME COLUMN PUNCTUATION TO SCORE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE programmes_areas RENAME COLUMN SCORE TO PUNCTUATION;
-- +goose StatementEnd
