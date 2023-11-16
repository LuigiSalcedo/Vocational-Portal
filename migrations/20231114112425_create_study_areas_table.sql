-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE study_areas (
  ID INTEGER NOT NULL,
  NAME VARCHAR(30) NOT NULL,
  PRIMARY KEY(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
