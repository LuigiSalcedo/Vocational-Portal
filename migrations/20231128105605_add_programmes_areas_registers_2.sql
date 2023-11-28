-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(24, 1, 0),
(24, 2, 0),
(24, 3, 50),
(24, 4, 10),
(24, 5, 0),
(24, 6, 50),
(24, 7, 80),
(24, 8, 55),
(24, 9, 65),
(24, 10, 5),
(24, 11, 10),
(24, 12, 10),
(24, 13, 40),
(24, 14, 0),
(24, 15, 0),
(25, 1, 0),
(25, 2, 0),
(25, 3, 0),
(25, 4, 20),
(25, 5, 5),
(25, 6, 5),
(25, 7, 0),
(25, 8, 50),
(25, 9, 30),
(25, 10, 30),
(25, 11, 85),
(25, 12, 5),
(25, 13, 0),
(25, 14, 85),
(25, 15, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM programmes_areas WHERE programm_id >= 24 AND programm_id <= 37;
-- +goose StatementEnd
