ALTER TABLE countries
    ADD COLUMN jwt_group VARCHAR(255);

UPDATE countries
SET jwt_group = id
where jwt_group is null;

ALTER TABLE countries
    ALTER COLUMN jwt_group SET NOT NULL;
