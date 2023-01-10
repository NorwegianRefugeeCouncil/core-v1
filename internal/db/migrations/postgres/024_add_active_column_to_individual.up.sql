ALTER TABLE individual_registrations

    -- add active column
    ADD COLUMN active      bool         NOT NULL DEFAULT true;