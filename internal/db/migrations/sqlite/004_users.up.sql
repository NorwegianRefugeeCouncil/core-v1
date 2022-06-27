CREATE TABLE IF NOT EXISTS users
(
    id      VARCHAR(32) PRIMARY KEY,
    subject VARCHAR(255) NOT NULL,
    email   VARCHAR(255) NOT NULL,
    UNIQUE (subject)
);

CREATE TABLE IF NOT EXISTS user_countries
(
    user_id    VARCHAR(32),
    country_id VARCHAR(32),
    primary key (user_id, country_id)
)