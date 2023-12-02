-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
DELETE FROM universities WHERE
ID = 6 or
ID = 8 or 
ID = 12 or
ID = 13 or
ID = 14;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
INSERT INTO universites(city_id, id, name) VALUES
(2, 6, 'UNIVERSIDAD NACIONAL (AMAZONIA)'),
(5, 8, 'UNIVERSIDAD NACIONAL (CARIBE)'),
(8, 12, 'UNIVERSIDAD NACIONAL (PLAMIRA)'),
(9, 13, 'UNIVERSIDAD NACIONAL (ORINOQUIA)'),
(10, 14, 'UNIVERSIDAD NACIONAL (TUMACO)');
-- +goose StatementEnd
