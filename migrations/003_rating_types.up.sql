CREATE TABLE rating_types (
    rating_id   BIGINT PRIMARY KEY DEFAULT unique_rowid(),
    code        VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(255) NULL
);