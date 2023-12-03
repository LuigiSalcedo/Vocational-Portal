-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO academic_programmes(id, name) VALUES
(71, 'MEDICINA');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM academic_programmes WHERE id = 71;
-- +goose StatementEnd
