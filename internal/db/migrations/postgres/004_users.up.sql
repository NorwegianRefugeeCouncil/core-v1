CREATE TABLE IF NOT EXISTS users
(
    id      VARCHAR(32) PRIMARY KEY,
    subject VARCHAR(255) NOT NULL,
    email   VARCHAR(255) NOT NULL,
    UNIQUE (subject)
);

CREATE TABLE IF NOT EXISTS user_countries
(
    user_id    VARCHAR(32) NOT NULL,
    country_id VARCHAR(32) NOT NULL,
    permission VARCHAR(32) NOT NULL,
    primary key (user_id, country_id),
    foreign key (country_id) REFERENCES countries (id),
    foreign key (user_id) REFERENCES users (id)
)