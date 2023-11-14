-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE academic_offers(
  ID SERIAL,
  UNIVERSITY_ID INTEGER NOT NULL,
  AP_ID INTEGER NOT NULL, -- AP = academic_programm
  NAME VARCHAR(50) NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(UNIVERSITY_ID) REFERENCES universities(ID),
  FOREIGN KEY(AP_ID) REFERENCES academic_programmes(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
