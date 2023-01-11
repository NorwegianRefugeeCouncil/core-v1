ALTER TABLE individual_registrations

    -- add active column
    ADD COLUMN IF NOT EXISTS active      bool         NOT NULL DEFAULT true;