-- +goose Up

-- +goose StatementBegin
CREATE TABLE subjects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    mata_pelajaran VARCHAR(100) NOT NULL,
    school_level ENUM('SD', 'SMP', 'SMA') NOT NULL,
    created_by_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by_id) REFERENCES teachers(id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE classes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name_kelas VARCHAR(50) NOT NULL,
    class_level VARCHAR(50) NOT NULL,
    school_level ENUM('SD', 'SMP', 'SMA') NOT NULL,
    guru_pengampu_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (guru_pengampu_id) REFERENCES teachers(id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE IF EXISTS classes;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS subjects;
-- +goose StatementEnd
