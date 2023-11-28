-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO academic_offers(UNIVERSITY_ID, AP_ID, NAME) VALUES
(2, 24, 'MARKETING Y TRANSFORMACION DIGITAL'),
(2, 25, 'CIENCIA DE DATOS'),
(2, 26, 'PSICOLOGIA'),
(2, 6, 'DERECHO'),
(2, 4, 'COMUNICACION SOCIAL'),
(2, 27, 'CIENCIA POLITICA Y RELACIONES INTERNACIONALES'),
(2, 28, 'FINANZAS Y NEGOCIOS INTERNACIONALES'),
(2, 5, 'CONTADURIA PUBLICA'),
(2, 5, 'CONTADURA PUBLICA - VIRTUAL'),
(2, 1, 'ADMINISTRACION DE EMPRESAS'),
(2, 7, 'ECONOMIA'),
(2, 13, 'INGENIERIA QUIMICA'),
(2, 29, 'INGENIERIA NAVAL'),
(2, 30, 'INGENIERIA MECATRONICA'),
(2, 31, 'INGENIERIA MECANICA'),
(2, 32, 'INGENIERIA INDUSTRIAL'),
(2, 33, 'INGENIERIA ELECTRONICA'),
(2, 34, 'INGENIERIA ELECTRICA'),
(2, 12, 'INGENIERIA DE SISTEMAS Y COMPUTACION'),
(2, 11, 'INGENIERIA CIVIL'),
(2, 35, 'INGENIERIA BIOMEDICA'),
(2, 36, 'INGENIERIA AMBIENTAL'),
(2, 37, 'ARQUITECTURA');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM academic_offers WHERE UNIVERSITY_ID = 2;
-- +goose StatementEnd