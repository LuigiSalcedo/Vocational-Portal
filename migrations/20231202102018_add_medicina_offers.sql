-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO academic_offers(university_id, ap_id, name) VALUES
(1, 71, 'MEDICINA'),
(4, 71, 'MEDICINA');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM academic_offers WHERE ap_id = 71;
-- +goose StatementEnd
