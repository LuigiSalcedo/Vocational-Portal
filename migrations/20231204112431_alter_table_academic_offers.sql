-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE academic_offers ALTER COLUMN price SET NOT NULL;
ALTER TABLE academic_offers ALTER COLUMN url SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE academic_offers ALTER COLUMN price DROP NOT NULL;
ALTER TABLE academic_offers ALTER COLUMN url DROP NOT NULL;
-- +goose StatementEnd
