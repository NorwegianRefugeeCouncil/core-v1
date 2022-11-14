ALTER TABLE individual_registrations

    -- rename identification_context column to engagement_context
    RENAME COLUMN identification_context TO engagement_context;
