ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS is_female_headed_household   BOOLEAN       NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS is_minor_headed_household    BOOLEAN       NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS first_name                   VARCHAR(255)  NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS middle_name                  VARCHAR(255)  NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS last_name                    VARCHAR(255)  NOT NULL DEFAULT '';

