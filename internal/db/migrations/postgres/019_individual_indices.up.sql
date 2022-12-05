-- Add indices on the individual_registrations table
CREATE INDEX IF NOT EXISTS idx_individual_registrations__country_id
    ON individual_registrations
        (
         country_id,
         deleted_at
            );
CREATE INDEX IF NOT EXISTS idx_individual_registrations__search
    ON individual_registrations
        (
         country_id,
         deleted_at,
         full_name,
         sex,
         free_field_1,
         free_field_2,
         free_field_3,
         free_field_4,
         free_field_5,
         collection_office,
         normalized_phone_number_1,
         normalized_phone_number_2,
         normalized_phone_number_3,
         address,
         phone_number_1,
         phone_number_2,
         phone_number_3,
         identification_number_1,
         identification_number_2,
         identification_number_3,
         is_minor,
         household_id,
         community_id
            );