-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SELECT 'up SQL query';
INSERT INTO academic_offers(university_id, ap_id, name) VALUES
(5, 1, 'TECNICO PROFESIONAL EN OPERACION DE PROCESOS ADUANEROS'),
(5, 1, 'TECNOLOGIA EN GESTION LOGISTICA Y PORTUARIA Y ADUANERA'),
(5, 1, 'ADMINISTRACION EN COMERCIO EXTERIOR'),
(5, 39, 'TECNOLOGIA EN TURISMO E IDIOMAS'),
(5, 39, 'TECNOLOGIA EN OPERACION TURISTICA'),
(5, 39, 'TECNOLOGIA EN GESTION TURISTICA'),
(5, 39, 'ADMINISTRACION TURISTICA'),
(5, 1, 'TECNOLOGIA EN GESTION EMPRESARIAL'),
(5, 1, 'TECNICO PROFESIONAL EN PROCESOS ADMINISTRATIVOS'),
(5, 1, 'PROFESIONAL EN ADMINISTRACION DE EMPRESAS'),
(5, 37, 'TECNOLOGIA EN DELINEANTE DE ARQUITECTURA E INGENIERIA'),
(5, 33, 'TECNICO PROFESIONAL EN MATENIMIENTO ELETROMECANICO');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM academic_offers WHERE university_id = 5;
-- +goose StatementEnd
