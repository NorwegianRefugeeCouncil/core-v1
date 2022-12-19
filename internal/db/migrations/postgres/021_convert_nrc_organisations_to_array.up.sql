ALTER TABLE countries

    -- rename nrc_organisation column to nrc_organisations
    ALTER COLUMN nrc_organisations DROP DEFAULT,
    ALTER COLUMN nrc_organisations TYPE VARCHAR(255)[] USING ARRAY[nrc_organisations],
    ALTER COLUMN nrc_organisations SET DEFAULT '{}'::VARCHAR(255)[];
