-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO programmes_areas (programm_id, area_id, score) VALUES
(38, 1, 0),
(38, 2, 0),
(38, 3, 0),
(38, 4, 20),
(38, 5, 10),
(38, 6, 0),
(38, 7, 0),
(38, 8, 10),
(38, 9, 30),
(38, 10, 0),
(38, 11, 100),
(38, 12, 30),
(38, 13, 30),
(38, 14, 75),
(38, 15, 65);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM programmes_areas WHERE programm_id = 38;
-- +goose StatementEnd
