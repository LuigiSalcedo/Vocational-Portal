-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO academic_programmes(id, name) VALUES
(67, 'OPTOMETRIA'),
(68, 'FISIOTERAPIA'),
(69, 'MODA'),
(70, 'BIOLOGIA MARINA');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM academic_programmes WHERE id >= 67 AND id <= 70;
-- +goose StatementEnd
