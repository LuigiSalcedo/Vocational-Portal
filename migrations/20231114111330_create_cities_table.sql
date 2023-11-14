-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE cities (
  ID INTEGER NOT NULL, 
  COUNTRY_ID INTEGER NOT NULL,
  NAME VARCHAR(50) NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(COUNTRY_ID) REFERENCES countries(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE cities;
-- +goose StatementEnd
