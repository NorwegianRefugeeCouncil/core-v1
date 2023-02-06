ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS mothers_name                  VARCHAR(255)  NOT NULL DEFAULT ''
