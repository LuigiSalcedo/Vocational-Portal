-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO universities(id, city_id, name, url) VALUES 
(2, 1, 'UNIVERSIDAD TECNOLOGICA DE BOLIVAR', 'https://www.utb.edu.co'),
(3, 1, 'FUNDACION UNIVERSITARIA TECNOLOGICO CONFENALCO', 'https://tecnologicocomfenalco.edu.co'),
(4, 1, 'UNIVERSIDAD DEL SINU (CARTAGENA)', 'https://www.unisinu.edu.co'),
(5, 1, 'INSTITUCION UNIVERSITARIA MAYOR DE CARTAGENA', 'https://umayor.edu.co');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM universities WHERE ID >= 2 AND ID <= 5;
-- +goose StatementEnd
