

CREATE TABLE companies (
    company_id   BIGINT  PRIMARY KEY DEFAULT unique_rowid(),
    ticker       VARCHAR(16) NOT NULL UNIQUE,
    name         VARCHAR(255) NOT NULL
);
