-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO academic_programmes(id, name) VALUES
(39, 'TURISMO');

INSERT INTO programmes_areas(programm_id, area_id, score) VALUES
(39, 1, 0),
(39, 2, 0),
(39, 3, 0),
(39, 4, 15),
(39, 5, 0),
(39, 6, 35),
(39, 7, 65),
(39, 8, 90),
(39, 9, 90),
(39, 10, 80),
(39, 11, 0),
(39, 12, 0),
(39, 13, 0),
(39, 14, 20),
(39, 15, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM programmes_areas WHERE programm_id = 39;
DELETE FROM academic_programmes WHERE id = 39;
-- +goose StatementEnd
