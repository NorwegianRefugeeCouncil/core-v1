CREATE TABLE IF NOT EXISTS countries
(
    id   VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(64)  NOT NULL,
    UNIQUE (code)
);

ALTER TABLE individuals ADD COLUMN countries text[];

create index individuals_countries_index
    on individuals (countries);

