-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
DELETE FROM programmes_areas WHERE programm_id = 15;
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES 
(15, 1, 0),
(15, 2, 0),
(15, 3, 0),
(15, 4, 40),
(15, 5, 0),
(15, 6, 100),
(15, 7, 80),
(15, 8, 10),
(15, 9, 5),
(15, 10, 0),
(15, 11, 0),
(15, 12, 0),
(15, 13, 0),
(15, 14, 0),
(15, 15, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM programmes_areas WHERE programm_id = 15;
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(15, 10, 10),
(15, 11, 90),
(15, 12, 5),
(15, 13, 0),
(15, 14, 50),
(15, 15, 80);
-- +goose StatementEnd
