-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE universities(
  ID SERIAL, 
  CITY_ID INTEGER NOT NULL,
  NAME VARCHAR(60),
  PRIMARY KEY(ID),
  FOREIGN KEY(CITY_ID) REFERENCES cities(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE universities;
-- +goose StatementEnd
