ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS has_disability     bool         NOT NULL DEFAULT false;
ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS pwd_comments      TEXT         NOT NULL DEFAULT '';