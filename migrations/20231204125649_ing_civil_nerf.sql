-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
DELETE from programmes_areas WHERE programm_id = 11;
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(11, 1, 20),
(11, 2, 0),
(11, 3, 0),
(11, 4, 20),
(11, 5, 0),
(11, 6, 5),
(11, 7, 0),
(11, 8, 5),
(11, 9, 15),
(11, 10, 0),
(11, 11, 100),
(11, 12, 70),
(11, 13, 80),
(11, 14, 70),
(11, 15, 50);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM programmes_areas WHERE programm_id = 11;
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(11, 1, 70),
(11, 2, 0),
(11, 3, 0),
(11, 4, 30),
(11, 5, 0),
(11, 6, 10),
(11, 7, 10),
(11, 8, 35),
(11, 9, 40),
(11, 10, 10),
(11, 11, 100),
(11, 12, 90),
(11, 13, 75),
(11, 14, 85),
(11, 15, 40);
-- +goose StatementEnd
