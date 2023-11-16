-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO universities(id, city_id, name) VALUES(1, 1, 'UNIVERSIDAD DE CARTAGENA');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM universities WHERE ID = 1;
-- +goose StatementEnd
