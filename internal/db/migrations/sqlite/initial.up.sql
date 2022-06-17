create table if not exists individuals
(
    id                      VARCHAR(32)
        primary key,
    full_name               VARCHAR(255),
    phone_number            VARCHAR(255),
    normalized_phone_number VARCHAR(255),
    email                   VARCHAR(255),
    address                 VARCHAR(255),
    birth_date              DATE,
    gender                  VARCHAR(255) default ''
);

create index if not exists individuals_email_index
    on individuals (email);

create index if not exists individuals_gender_index
    on individuals (gender);

create index if not exists individuals_search_idx
    on individuals (id, full_name, phone_number, normalized_phone_number, email, address, birth_date, gender);

CREATE VIRTUAL TABLE IF NOT EXISTS individuals_fts USING fts5
(
    full_name,
    phone_number,
    normalized_phone_number,
    email,
    address,
    birth_date
);

CREATE TRIGGER if not exists individuals_fts_insert
    AFTER INSERT
    ON individuals
BEGIN
    INSERT INTO individuals_fts(rowid, full_name, phone_number, normalized_phone_number, email, address, birth_date)
    VALUES (new.rowid, new.full_name, new.phone_number, new.normalized_phone_number, new.email, new.address, new.birth_date);
END;


CREATE TRIGGER if not exists individuals_fts_update
    AFTER UPDATE
    ON individuals
BEGIN
    UPDATE individuals_fts
    SET full_name=new.full_name,
        phone_number=new.phone_number,
        normalized_phone_number=new.normalized_phone_number,
        email=new.email,
        address=new.address,
        birth_date=new.birth_date
    WHERE rowid = new.rowid;
END;


CREATE TRIGGER if not exists individuals_fts_delete
    AFTER DELETE
    ON individuals
BEGIN
    DELETE FROM individuals_fts WHERE rowid = old.rowid;
END;

