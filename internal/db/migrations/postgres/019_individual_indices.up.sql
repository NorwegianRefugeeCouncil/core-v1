-- Add indices on the individual_registrations table

CREATE INDEX IF NOT EXISTS idx_individual_registrations__address ON individual_registrations (address);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__age ON individual_registrations (age);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__birth_date ON individual_registrations (birth_date);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__community_id ON individual_registrations (community_id);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__country_id ON individual_registrations (country_id);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__displacement_status ON individual_registrations (displacement_status);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__email_1 ON individual_registrations (email_1);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__email_2 ON individual_registrations (email_2);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__email_3 ON individual_registrations (email_3);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__free_field_1 ON individual_registrations (free_field_1);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__free_field_2 ON individual_registrations (free_field_2);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__free_field_3 ON individual_registrations (free_field_3);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__free_field_4 ON individual_registrations (free_field_4);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__free_field_5 ON individual_registrations (free_field_5);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__full_name ON individual_registrations (full_name);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__household_id ON individual_registrations (household_id);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__is_head_of_community ON individual_registrations (is_head_of_community);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__is_head_of_household ON individual_registrations (is_head_of_household);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__is_minor ON individual_registrations (is_minor);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__phone_number_1 ON individual_registrations (phone_number_1);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__phone_number_2 ON individual_registrations (phone_number_2);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__phone_number_3 ON individual_registrations (phone_number_3);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__preferred_communication_language ON individual_registrations (preferred_communication_language);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__preferred_name ON individual_registrations (preferred_name);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__sex ON individual_registrations (sex);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__spoken_language_1 ON individual_registrations (spoken_language_1);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__spoken_language_2 ON individual_registrations (spoken_language_2);
CREATE INDEX IF NOT EXISTS idx_individual_registrations__spoken_language_3 ON individual_registrations (spoken_language_3);





