-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE offer_searches(
  ID SERIAL,
  OFFER_ID INTEGER NOT NULL,
  SEARCH_DATE DATE NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(OFFER_ID) REFERENCES academic_offers(Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
