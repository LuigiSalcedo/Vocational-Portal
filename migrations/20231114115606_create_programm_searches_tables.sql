-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE programm_searches(
  ID SERIAL,
  PROGRAMM_ID INTEGER NOT NULL,
  SEARCH_DATE DATE NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(PROGRAMM_ID) REFERENCES academic_programmes(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
