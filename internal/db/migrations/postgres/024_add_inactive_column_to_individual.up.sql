ALTER TABLE individual_registrations
    DROP COLUMN IF EXISTS active;

ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS inactive      bool         NOT NULL DEFAULT false;