-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(67, 1, 0),
(67, 2, 0),
(67, 3, 0),
(67, 4, 0),
(67, 5, 100),
(67, 6, 0),
(67, 7, 0),
(67, 8, 0),
(67, 9, 0),
(67, 10, 0),
(67, 11, 50),
(67, 12, 0),
(67, 13, 0),
(67, 14, 10),
(67, 15, 50),
(68, 1, 0),
(68, 2, 0),
(68, 3, 0),
(68, 4, 0),
(68, 5, 100),
(68, 6, 0),
(68, 7, 0),
(68, 8, 0),
(68, 9, 0),
(68, 10, 0),
(68, 11, 0),
(68, 12, 0),
(68, 13, 0),
(68, 14, 0),
(68, 15, 70),
(69, 1, 0),
(69, 2, 0),
(69, 3, 90),
(69, 4, 20),
(69, 5, 0),
(69, 6, 10),
(69, 7, 40),
(69, 8, 0),
(69, 9, 0),
(69, 10, 0),
(69, 11, 0),
(69, 12, 0),
(69, 13, 0),
(69, 14, 0),
(69, 15, 0),
(70, 1, 0),
(70, 2, 10),
(70, 2, 0),
(70, 3, 0),
(70, 4, 40),
(70, 5, 0),
(70, 6, 0),
(70, 7, 0),
(70, 8, 0),
(70, 9, 0),
(70, 10, 0),
(70, 11, 10),
(70, 12, 0),
(70, 13, 0),
(70, 14, 30),
(70, 15, 100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE  FROM programmes_areas WHERE programm_id >= 67 AND programm_id <= 70;
-- +goose StatementEnd
