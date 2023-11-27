ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS has_medical_condition bool DEFAULT null;
ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS needs_legal_and_physical_protection bool DEFAULT null;
ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS is_child_at_risk bool DEFAULT null;
ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS is_woman_at_risk bool DEFAULT null;
ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS is_elder_at_risk bool DEFAULT null;
ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS is_pregnant bool DEFAULT null;
ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS is_lactating bool DEFAULT null;
ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS is_single_parent bool DEFAULT null;
ALTER TABLE individual_registrations
    ADD COLUMN IF NOT EXISTS is_separated_child bool DEFAULT null;