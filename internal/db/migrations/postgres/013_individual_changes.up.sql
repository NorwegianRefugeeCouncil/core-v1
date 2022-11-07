-- disable the soft-delete protection trigger
ALTER TABLE individual_registrations
    DISABLE TRIGGER individual_registrations_prevent_deleted_update;

-- add phone number columns
ALTER TABLE individual_registrations
    ADD COLUMN age                       INT          NULL,
    ADD COLUMN phone_number_1            VARCHAR(64)  NOT NULL DEFAULT '',
    ADD COLUMN phone_number_2            VARCHAR(64)  NOT NULL DEFAULT '',
    ADD COLUMN phone_number_3            VARCHAR(64)  NOT NULL DEFAULT '',
    ADD COLUMN normalized_phone_number_1 VARCHAR(64)  NOT NULL DEFAULT '',
    ADD COLUMN normalized_phone_number_2 VARCHAR(64)  NOT NULL DEFAULT '',
    ADD COLUMN normalized_phone_number_3 VARCHAR(64)  NOT NULL DEFAULT '',
    ADD COLUMN email_1                   VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN email_2                   VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN email_3                   VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN free_field_1              VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN free_field_2              VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN free_field_3              VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN free_field_4              VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN free_field_5              VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN comments                  TEXT         NOT NULL DEFAULT '';

-- copy over phone numbers and emails
UPDATE individual_registrations
SET phone_number_1            = phone_number,
    email_1                   = email,
    normalized_phone_number_1 = normalized_phone_number
WHERE phone_number <> ''
   OR email <> ''
   OR normalized_phone_number <> '';

UPDATE individual_registrations
SET displacement_status = 'returnee'
WHERE displacement_status = 'stateless';

UPDATE individual_registrations
SET age = EXTRACT(YEAR FROM AGE(NOW(), birth_date))
WHERE birth_date IS NOT NULL;

-- delete phone number column
ALTER TABLE individual_registrations
    DROP COLUMN phone_number,
    DROP COLUMN normalized_phone_number,
    DROP COLUMN email;

-- enable the soft-delete protection trigger
ALTER TABLE individual_registrations
    ENABLE TRIGGER individual_registrations_prevent_deleted_update;