ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS is_female_headed_household   BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS is_minor_headed_household    BOOLEAN NOT NULL DEFAULT TRUE;
