-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO study_areas (id, name) VALUES
(1, 'AGRONOMIA'),
(2, 'VETERINARIA'),
(3, 'BELLAS ARTES'),
(4, 'CIENCIAS DE EDUCACION'),
(5, 'CIENCIAS DE LA SALUD'),
(6, 'CIENCIAS SOCIALES'),
(7, 'HUMANIDADES'),
(8, 'ECONOMIA'),
(9, 'ADMINISTRACION'),
(10, 'CONTADURIA'),
(11, 'INGENIERIA'),
(12, 'ARQUITECTURA'),
(13, 'URBANISMO'),
(14, 'MATEMATICAS'),
(15, 'CIENCIAS NATURALES');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM study_areas WHERE ID >= 1 AND ID <= 15;
-- +goose StatementEnd
