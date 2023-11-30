-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE preferences (
  ID INTEGER NOT NULL,
  NAME VARCHAR(60) NOT NULL,
  PRIMARY KEY(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE prefereces;
-- +goose StatementEnd
