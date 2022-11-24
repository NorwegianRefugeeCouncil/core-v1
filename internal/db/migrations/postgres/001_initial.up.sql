CREATE TABLE IF NOT EXISTS countries
(
    id        UUID          PRIMARY KEY,
    name      VARCHAR(255)  NOT NULL,
    code      VARCHAR(64)   NOT NULL,
    jwt_group VARCHAR(255)  NOT NULL
    UNIQUE (code)
);

CREATE TABLE IF NOT EXISTS individual_registrations
(
    -- system columns
    id                                UUID                     PRIMARY KEY,
    country_id                        UUID                     NOT NULL,
    created_at                        timestamp with time zone NOT NULL,
    updated_at                        timestamp with time zone NOT NULL,
    deleted_at                        timestamp with time zone NULL,

    -- GRF columns
    address                           VARCHAR(512)             NOT NULL,
    age                               INT                      NULL,
    birth_date                        DATE                     NULL,
    cognitive_disability_level        VARCHAR(32)              NOT NULL,
    collection_administrative_area_1  VARCHAR(128)             NOT NULL,
    collection_administrative_area_2  VARCHAR(128)             NOT NULL,
    collection_administrative_area_3  VARCHAR(128)             NOT NULL,
    collection_agent_name             VARCHAR(128)             NOT NULL,
    collection_agent_title            VARCHAR(64)              NOT NULL,
    collection_time                   timestamp with time zone NOT NULL,
    communication_disability_level    VARCHAR(32)              NOT NULL,
    community_id                      VARCHAR(64)              NOT NULL,
    displacement_status               VARCHAR(64)              NOT NULL,
    displacement_status_comment       TEXT                     NOT NULL DEFAULT '';
    email_1                           VARCHAR(255)             NOT NULL DEFAULT '',
    email_2                           VARCHAR(255)             NOT NULL DEFAULT '',
    email_3                           VARCHAR(255)             NOT NULL DEFAULT '',
    free_field_1                      VARCHAR(255)             NOT NULL DEFAULT '',
    free_field_2                      VARCHAR(255)             NOT NULL DEFAULT '',
    free_field_3                      VARCHAR(255)             NOT NULL DEFAULT '',
    free_field_4                      VARCHAR(255)             NOT NULL DEFAULT '',
    free_field_5                      VARCHAR(255)             NOT NULL DEFAULT '',
    full_name                         VARCHAR(255)             NOT NULL,
    sex                               VARCHAR(32)              NOT NULL,
    has_cognitive_disability          BOOLEAN                  NOT NULL,
    has_communication_disability      BOOLEAN                  NOT NULL,
    has_consented_to_rgpd             BOOLEAN                  NOT NULL,
    has_consented_to_referral         BOOLEAN                  NOT NULL,
    has_hearing_disability            BOOLEAN                  NOT NULL,
    has_mobility_disability           BOOLEAN                  NOT NULL,
    has_selfcare_disability           BOOLEAN                  NOT NULL,
    has_vision_disability             BOOLEAN                  NOT NULL,
    hearing_disability_level          VARCHAR(32)              NOT NULL,
    household_id                      VARCHAR(64)              NOT NULL,
    identification_type_1             VARCHAR(64)              NOT NULL,
    identification_type_explanation_1 TEXT                     NOT NULL,
    identification_number_1           VARCHAR(64)              NOT NULL,
    identification_type_2             VARCHAR(64)              NOT NULL,
    identification_type_explanation_2 TEXT                     NOT NULL,
    identification_number_2           VARCHAR(64)              NOT NULL,
    identification_type_3             VARCHAR(64)              NOT NULL,
    identification_type_explanation_3 TEXT                     NOT NULL,
    identification_number_3           VARCHAR(64)              NOT NULL,
    engagement_context                VARCHAR(64)              NOT NULL,
    internal_id                       VARCHAR(64)              NOT NULL,
    is_head_of_community              BOOLEAN                  NOT NULL,
    is_head_of_household              BOOLEAN                  NOT NULL,
    is_minor                          BOOLEAN                  NOT NULL,
    mobility_disability_level         VARCHAR(32)              NOT NULL,
    nationality_1                     VARCHAR(64)              NOT NULL,
    nationality_2                     VARCHAR(64)              NOT NULL,
    normalized_phone_number_1         VARCHAR(64)              NOT NULL DEFAULT '',
    normalized_phone_number_2         VARCHAR(64)              NOT NULL DEFAULT '',
    normalized_phone_number_3         VARCHAR(64)              NOT NULL DEFAULT '',
    phone_number_1                    VARCHAR(64)              NOT NULL DEFAULT '',
    phone_number_2                    VARCHAR(64)              NOT NULL DEFAULT '',
    phone_number_3                    VARCHAR(64)              NOT NULL DEFAULT '',
    preferred_contact_method          VARCHAR(64)              NOT NULL,
    preferred_contact_method_comments TEXT                     NOT NULL,
    preferred_name                    VARCHAR(255)             NOT NULL,
    preferred_communication_language  VARCHAR(64)              NOT NULL,
    prefers_to_remain_anonymous       BOOLEAN                  NOT NULL,
    presents_protection_concerns      BOOLEAN                  NOT NULL,
    selfcare_disability_level         VARCHAR(32)              NOT NULL,
    spoken_language_1                 VARCHAR(64)              NOT NULL,
    spoken_language_2                 VARCHAR(64)              NOT NULL,
    spoken_language_3                 VARCHAR(64)              NOT NULL,
    vision_disability_level           VARCHAR(32)              NOT NULL,
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
