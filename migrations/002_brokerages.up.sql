CREATE TABLE brokerages (
    brokerage_id BIGINT PRIMARY KEY DEFAULT unique_rowid(),
    name         VARCHAR(255) NOT NULL UNIQUE
);