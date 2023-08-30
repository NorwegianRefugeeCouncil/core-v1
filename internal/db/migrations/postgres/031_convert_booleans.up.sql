ALTER TABLE individual_registrations ALTER COLUMN has_vision_disability DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN has_vision_disability SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN has_mobility_disability DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN has_mobility_disability SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN has_selfcare_disability DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN has_selfcare_disability SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN has_cognitive_disability DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN has_cognitive_disability SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN has_hearing_disability DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN has_hearing_disability SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN has_disability DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN has_disability SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN has_communication_disability DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN has_communication_disability SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN has_consented_to_rgpd DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN has_consented_to_rgpd SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN has_consented_to_referral DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN has_consented_to_referral SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN is_minor_headed_household DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN is_minor_headed_household SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN is_minor DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN is_minor SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN is_head_of_household DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN is_head_of_household SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN is_female_headed_household DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN is_female_headed_household SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN is_head_of_community DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN is_head_of_community SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN presents_protection_concerns DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN presents_protection_concerns SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN prefers_to_remain_anonymous DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN prefers_to_remain_anonymous SET DEFAULT null;

ALTER TABLE individual_registrations ALTER COLUMN inactive DROP NOT NULL;
ALTER TABLE individual_registrations ALTER COLUMN inactive SET DEFAULT null;
