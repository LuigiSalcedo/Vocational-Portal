-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO countries(id, name) VALUES(1, 'COLOMBIA');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM countries WHERE ID = 1;
-- +goose StatementEnd
