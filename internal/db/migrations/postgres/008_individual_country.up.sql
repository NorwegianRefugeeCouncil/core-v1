ALTER TABLE individuals
    DROP COLUMN countries;

ALTER TABLE individuals
    ADD COLUMN country_id VARCHAR(255);

