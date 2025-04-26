-- +goose Up
-- +goose StatementBegin
CREATE TYPE school_level_enum AS ENUM ('SD', 'SMP', 'SMA');
CREATE TYPE role_enum AS ENUM ('guru', 'siswa');

CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    nidn VARCHAR(50),
    school_name VARCHAR(100),
    school_level school_level_enum,
    is_slb BOOLEAN DEFAULT FALSE,
    role role_enum DEFAULT 'guru',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE teachers;
DROP TYPE IF EXISTS school_level_enum;
DROP TYPE IF EXISTS role_enum;
-- +goose StatementEnd
