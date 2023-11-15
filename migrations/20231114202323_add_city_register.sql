-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO cities(id, country_id, name) VALUES(1, 1, 'CARTAGENA DE INDIAS');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM cities WHERE ID = 1;
-- +goose StatementEnd
