-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(71, 1, 0),
(71, 2, 10),
(71, 3, 0),
(71, 4, 5),
(71, 5, 100),
(71, 6, 0),
(71, 7, 5),
(71, 8, 0),
(71, 9, 0),
(71, 10, 0),
(71, 11, 0),
(71, 12, 0),
(71, 13, 0),
(71, 14, 0),
(71, 15, 90);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM programmes_areas WHERE programm_id = 71;
-- +goose StatementEnd
