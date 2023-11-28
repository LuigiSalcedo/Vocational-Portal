-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO academic_programmes VALUES 
(38, 'INGENIERIA DE PROCESO');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM academic_programmes WHERE ID = 38;
-- +goose StatementEnd
