ALTER TABLE countries
    -- drop jwt_group column
    DROP COLUMN IF EXISTS jwt_group;

ALTER TABLE countries
    -- add nrc_organisation column
    ADD COLUMN nrc_organisation      VARCHAR(255)         NOT NULL DEFAULT '';
