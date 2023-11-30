-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO preferences VALUES 
(1, 'CULTIVAR'),
(2, 'NATURALEZA'),
(3, 'FLORES'),
(4, 'COMPUTADORAS'),
(5, 'INVENTAR'),
(6, 'ESCRIBIR'),
(7, 'IMAGINAR'),
(8, 'HABLAR'),
(9, 'ENSEÑAR'),
(10, 'APRENDER'),
(11, 'LEER'),
(12, 'HISTORIA'),
(13, 'MUSICA'),
(14, 'BARCOS'),
(15, 'NUMEROS'),
(16, 'DINERO'),
(17, 'SEGURIDAD'),
(18, 'SALUD'),
(19, 'CULTURA'),
(20, 'AUTOMOVILES'),
(21, 'EDIFICACIONES'),
(22, 'CONSTRUCCIONES'),
(23, 'ANALISIS'),
(24, 'ANIMALES'),
(25, 'PENSAMIENTO'),
(26, 'CONTROL'),
(27, 'COMIDA'),
(28, 'SOCIALIZAR'),
(29, 'VENDER'),
(30, 'LIDERAZGO'),
(31, 'FUTURO');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM preferences WHERE ID >= 1 AND ID <= 31;
-- +goose StatementEnd
