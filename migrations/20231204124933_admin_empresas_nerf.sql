-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
DELETE FROM programmes_areas WHERE programm_id = 1;
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(1, 1, 0),
(1, 2, 0),
(1, 3, 0),
(1, 4, 5),
(1, 5, 0),
(1, 6, 0),
(1, 7, 0),
(1, 8, 80),
(1, 9, 100),
(1, 10, 80),
(1, 11, 0),
(1, 12, 0),
(1, 13, 0),
(1, 14, 10),
(1, 15, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM programmes_areas WHERE programm_id = 1;
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(1, 1, 10),
(1, 2, 10),
(1, 3, 0),
(1, 4, 40),
(1, 5, 30),
(1, 6, 70),
(1, 7, 50),
(1, 8, 90),
(1, 9, 100),
(1, 10, 95),
(1, 11, 30),
(1, 12, 15),
(1, 13, 15),
(1, 14, 68),
(1, 15, 10);
-- +goose StatementEnd
