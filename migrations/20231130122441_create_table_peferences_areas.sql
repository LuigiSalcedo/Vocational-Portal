-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE preferences_areas(
  ID SERIAL PRIMARY KEY,
  PREFERENCE_ID INTEGER NOT NULL,
  AREA_ID INTEGER NOT NULL,
  FOREIGN KEY(PREFERENCE_ID) REFERENCES preferences(ID),
  FOREIGN KEY(AREA_ID) REFERENCES study_areas(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE preferences_areas;
-- +goose StatementEnd
