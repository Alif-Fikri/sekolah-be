-- +goose Up
-- +goose StatementBegin
CREATE TABLE teachers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    nik VARCHAR(50) UNIQUE NOT NULL,
    nuptk VARCHAR(50),
    school_name VARCHAR(100),
    school_level ENUM('SD', 'SMP', 'SMA'),
    is_slb BOOLEAN DEFAULT FALSE,
    role ENUM('guru', 'siswa') DEFAULT 'guru',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE teachers;
-- +goose StatementEnd
