-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE academic_offers ADD COLUMN price VARCHAR(20);
ALTER TABLE academic_offers ADD COLUMN url VARCHAR(200);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE academic_offers DROP COLUMN(url);
ALTER TABLE academic_offers DROP COLUMN(price);
-- +goose StatementEnd
