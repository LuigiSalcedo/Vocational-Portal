-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE programmes_areas(
  PROGRAMME_ID INTEGER NOT NULL,
  AREA_ID INTEGER NOT NULL,
  PUNTUATION INTEGER NOT NULL,
  FOREIGN KEY(PROGRAMME_ID) REFERENCES academic_programmes(ID),
  FOREIGN KEY(AREA_ID) REFERENCES study_areas(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
