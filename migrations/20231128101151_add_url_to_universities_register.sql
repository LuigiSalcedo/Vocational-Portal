-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
UPDATE universities SET URL = 'https://www.unicartagena.edu.co' WHERE ID = 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
UPDATE universities SET URL = '' WHERE ID = 1;
-- +goose StatementEnd
