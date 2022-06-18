create table if not exists individuals
(
    id                      VARCHAR(32) primary key,
    full_name               VARCHAR(255),
    phone_number            VARCHAR(255),
    normalized_phone_number VARCHAR(255),
    email                   VARCHAR(255),
    address                 VARCHAR(255),
    birth_date              DATE,
    gender                  VARCHAR(255)
);

create index if not exists individuals_email_index
    on individuals (email);

create index if not exists individuals_gender_index
    on individuals (gender);

create index if not exists individuals_search_idx
    on individuals (id, full_name, phone_number, normalized_phone_number, email, address, birth_date, gender);
