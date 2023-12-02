-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO universities(city_id, id, name, url) VALUES
(2, 6, 'UNIVERSIDAD NACIONAL (AMAZONIA)', 'https://amazonia.unal.edu.co'),
(3, 7, 'UNIVERSIDAD NACIONAL', 'https://registro.bogota.unal.edu.co'),
(5, 8, 'UNIVERSIDAD NACIONAL (CARIBE)', 'https://caribe.unal.edu.co'),
(4, 9, 'UNIVERSIDAD NACIONAL (LA PAZ)', 'https://delapaz.unal.edu.co'),
(6, 10, 'UNIVERSIDAD NACIONAL (MANIZALES)', 'https://registro.manizales.unal.edu.co'),
(7, 11, 'UNIVERSIDAD NACIONAL (MEDELLIN)', 'https://registroymatricula.medellin.unal.edu.co'),
(8, 12, 'UNIVERSIDAD NACIONAL (PALMIRA)', 'https://palmira.unal.edu.co'),
(9, 13, 'UNIVERSIDAD NACIONAL (ORINOQUIA)', 'http://orinoquia.unal.edu.co'),
(10, 14, 'UNIVERSIDAD NACIONAL (TUMACO)', 'https://tumaco-pacifico.unal.edu.co');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM universities WHERE id >= 6 AND id <= 14;
-- +goose StatementEnd
