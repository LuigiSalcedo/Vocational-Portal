-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO cities(country_id, id, name) VALUES
(1, 2, 'LETICIA'),
(1, 3, 'BOGOTÁ'),
(1, 4, 'LA PAZ'),
(1, 5, 'SAN ANDRÉS'),
(1, 6, 'MANIZALES'),
(1, 7, 'MEDELLÍN'),
(1, 8, 'PALMIRA'),
(1, 9, 'ARAUCA'),
(1, 10, 'TUMACO');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM cities WHERE id >= 2 AND i <= 10;
-- +goose StatementEnd
