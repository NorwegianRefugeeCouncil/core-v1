ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS service_cc_1               VARCHAR(64) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS service_requested_date_1   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_delivered_date_1   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_comments_1         TEXT NOT NULL DEFAULT '',

    ADD COLUMN IF NOT EXISTS service_cc_2               VARCHAR(64) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS service_requested_date_2   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_delivered_date_2   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_comments_2         TEXT NOT NULL DEFAULT '',

    ADD COLUMN IF NOT EXISTS service_cc_3               VARCHAR(64) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS service_requested_date_3   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_delivered_date_3   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_comments_3         TEXT NOT NULL DEFAULT '',

    ADD COLUMN IF NOT EXISTS service_cc_4               VARCHAR(64) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS service_requested_date_4   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_delivered_date_4   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_comments_4         TEXT NOT NULL DEFAULT '',

    ADD COLUMN IF NOT EXISTS service_cc_5               VARCHAR(64) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS service_requested_date_5   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_delivered_date_5   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_comments_5         TEXT NOT NULL DEFAULT '',

    ADD COLUMN IF NOT EXISTS service_cc_6               VARCHAR(64) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS service_requested_date_6   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_delivered_date_6   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_comments_6         TEXT NOT NULL DEFAULT '',

    ADD COLUMN IF NOT EXISTS service_cc_7               VARCHAR(64) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS service_requested_date_7   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_delivered_date_7   DATE DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS service_comments_7         TEXT NOT NULL DEFAULT '';

