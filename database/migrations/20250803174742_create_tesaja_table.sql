-- -- +goose Up
-- ALTER TABLE students
-- ADD COLUMN created_by INT,
-- ADD CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES teachers(id) ON DELETE SET NULL;

-- -- +goose Down
-- ALTER TABLE students
-- DROP FOREIGN KEY fk_created_by,
-- DROP COLUMN created_by;
