ALTER TABLE individual_registrations

    -- add comment column
    ADD COLUMN IF NOT EXISTS displacement_status_comment      TEXT         NOT NULL DEFAULT '';