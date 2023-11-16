-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE academic_programmes (
  ID INTEGER NOT NULL,
  NAME VARCHAR(50),
  PRIMARY KEY(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
