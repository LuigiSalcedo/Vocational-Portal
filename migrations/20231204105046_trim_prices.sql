-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
UPDATE academic_offers SET price = TRIM(price);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
