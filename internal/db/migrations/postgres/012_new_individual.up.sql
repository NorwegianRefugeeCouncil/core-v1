DROP table IF EXISTS individuals;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

ALTER TABLE countries
    ADD COLUMN IF NOT EXISTS id_temp UUID DEFAULT gen_random_uuid();

ALTER TABLE countries
    RENAME COLUMN id TO id_old;

ALTER TABLE countries
    RENAME COLUMN id_temp TO id;

ALTER TABLE countries
    DROP COLUMN id_old;

ALTER TABLE countries
    ADD PRIMARY KEY (id);

DROP TABLE IF EXISTS individual_registrations;
CREATE TABLE individual_registrations
(
    address                           varchar(512)             NOT NULL,
    birth_date                        date                     NULL,
    cognitive_disability_level        varchar(32)              NOT NULL,
    collection_administrative_area_1  varchar(128)             NOT NULL,
    collection_administrative_area_2  varchar(128)             NOT NULL,
    collection_administrative_area_3  varchar(128)             NOT NULL,
    collection_agent_name             varchar(128)             NOT NULL,
    collection_agent_title            varchar(64)              NOT NULL,
    collection_time                   timestamp with time zone NOT NULL,
    communication_disability_level    varchar(32)              NOT NULL,
    community_id                      varchar(64)              NOT NULL,
    country_id                        uuid                     NOT NULL,
    created_at                        timestamp with time zone NOT NULL,
    deleted_at                        timestamp with time zone NULL,
    displacement_status               varchar(64)              NOT NULL,
    email                             varchar(255)             NOT NULL,
    full_name                         varchar(255)             NOT NULL,
    gender                            varchar(32)              NOT NULL,
    has_cognitive_disability          boolean                  NOT NULL,
    has_communication_disability      boolean                  NOT NULL,
    has_consented_to_rgpd             boolean                  NOT NULL,
    has_consented_to_referral         boolean                  NOT NULL,
    has_hearing_disability            boolean                  NOT NULL,
    has_mobility_disability           boolean                  NOT NULL,
    has_selfcare_disability           boolean                  NOT NULL,
    has_vision_disability             boolean                  NOT NULL,
    hearing_disability_level          varchar(32)              NOT NULL,
    household_id                      varchar(64)              NOT NULL,
    id                                uuid                     NOT NULL,
    identification_type_1             varchar(64)              NOT NULL,
    identification_type_explanation_1 text                     NOT NULL,
    identification_number_1           varchar(64)              NOT NULL,
    identification_type_2             varchar(64)              NOT NULL,
    identification_type_explanation_2 text                     NOT NULL,
    identification_number_2           varchar(64)              NOT NULL,
    identification_type_3             varchar(64)              NOT NULL,
    identification_type_explanation_3 text                     NOT NULL,
    identification_number_3           varchar(64)              NOT NULL,
    identification_context            varchar(64)              NOT NULL,
    internal_id                       varchar(64)              NOT NULL,
    is_head_of_community              boolean                  NOT NULL,
    is_head_of_household              boolean                  NOT NULL,
    is_minor                          boolean                  NOT NULL,
    mobility_disability_level         varchar(32)              NOT NULL,
    nationality_1                     varchar(64)              NOT NULL,
    nationality_2                     varchar(64)              NOT NULL,
    normalized_phone_number           varchar(64)              NOT NULL,
    phone_number                      varchar(64)              NOT NULL,
    preferred_contact_method          varchar(64)              NOT NULL,
    preferred_contact_method_comments text                     NOT NULL,
    preferred_name                    varchar(255)             NOT NULL,
    preferred_communication_language  varchar(64)              NOT NULL,
    prefers_to_remain_anonymous       boolean                  NOT NULL,
    presents_protection_concerns      boolean                  NOT NULL,
    selfcare_disability_level         varchar(32)              NOT NULL,
    spoken_language_1                 varchar(64)              NOT NULL,
    spoken_language_2                 varchar(64)              NOT NULL,
    spoken_language_3                 varchar(64)              NOT NULL,
    updated_at                        timestamp with time zone NOT NULL,
    vision_disability_level           varchar(32)              NOT NULL,
    CONSTRAINT individual_registration_pkey PRIMARY KEY (id)
);


/*
this procedure is used to prevent updating soft-deleted
records.

this procedure will be triggered before any update
on the individuals table.

Internally, it compares the old and the new record.
If any of the values are being changed, other than the
deleted_at column, it will throw an error.

It relies on the fact that for each update, the application
always sets the updated_at column. So for each update, there
will always be a change detected by this procedure,
and it will successfully throw an error if a user tries
to update a soft-deleted record.
 */
CREATE OR REPLACE FUNCTION prevent_deleted_individual_update()
    RETURNS TRIGGER
AS
$$
BEGIN
    IF (
                TG_OP = 'UPDATE'
            AND OLD.deleted_at IS NOT NULL
            AND EXISTS(
                        SELECT 1
                        FROM jsonb_each(to_jsonb(NEW)) AS post,
                             jsonb_each(to_jsonb(OLD)) AS pre
                        WHERE post.key = pre.key
                          AND (
                            post.key != 'deleted_at' AND post.value IS DISTINCT FROM pre.value
                            )
                          AND post.key != 'deleted_at'
                    )
        ) THEN
        RAISE EXCEPTION 'Cannot update a soft deleted record';

    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- trigger to update updated_at column
CREATE OR REPLACE TRIGGER individual_registrations_prevent_deleted_update
    BEFORE UPDATE
    ON individual_registrations
    FOR EACH ROW
EXECUTE PROCEDURE prevent_deleted_individual_update();

