ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS inactive      bool         NOT NULL DEFAULT false;