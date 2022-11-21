ALTER TABLE individual_registrations

    -- add office column
    ADD COLUMN collection_office      VARCHAR(128)         NOT NULL DEFAULT '';