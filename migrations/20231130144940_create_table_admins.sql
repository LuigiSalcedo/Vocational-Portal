-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE admins (
  NAME VARCHAR(100) NOT NULL PRIMARY KEY,
  PASSWORD VARCHAR(250) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE admins;
-- +goose StatementEnd
