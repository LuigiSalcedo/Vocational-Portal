-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
DELETE FROM programmes_areas WHERE programm_id = 2;
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(2, 1, 0),
(2, 2, 0),
(2, 3, 0),
(2, 4, 0),
(2, 5, 0),
(2, 6, 0),
(2, 7, 0),
(2, 8, 40),
(2, 9, 100),
(2, 10, 10),
(2, 11, 30),
(2, 12, 0),
(2, 13, 0),
(2, 14, 5),
(2, 15, 5);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM programmes_areas WHERE programm_id = 2;
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(2, 1, 10),
(2, 2, 10),
(2, 2, 0),
(2, 3, 40),
(2, 4, 40),
(2, 5, 20),
(2, 6, 60),
(2, 7, 40),
(2, 8, 80),
(2, 9, 80),
(2, 10, 80),
(2, 11, 90),
(2, 12, 30),
(2, 13, 25),
(2, 14, 70),
(2, 15, 25);
-- +goose StatementEnd
