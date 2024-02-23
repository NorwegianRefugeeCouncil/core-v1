CREATE OR REPLACE FUNCTION no_empty_override()
    RETURNS TRIGGER
AS
$$BEGIN

    IF OLD.address != '' AND NEW.address = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='address', HINT = 'error_empty_override';
    END IF;

    IF OLD.birth_date != NULL AND NEW.birth_date = NULL THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='birth_date', HINT = 'error_empty_override';
    END IF;

    IF OLD.cognitive_disability_level != '' AND NEW.cognitive_disability_level = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='cognitive_disability_level', HINT = 'error_empty_override';
    END IF;

    IF OLD.collection_administrative_area_1 != '' AND NEW.collection_administrative_area_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='collection_administrative_area_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.collection_administrative_area_2 != '' AND NEW.collection_administrative_area_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='collection_administrative_area_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.collection_administrative_area_3 != '' AND NEW.collection_administrative_area_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='collection_administrative_area_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.collection_agent_name != '' AND NEW.collection_agent_name = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='collection_agent_name', HINT = 'error_empty_override';
    END IF;

    IF OLD.collection_agent_title != '' AND NEW.collection_agent_title = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='collection_agent_title', HINT = 'error_empty_override';
    END IF;

    IF OLD.collection_time != NULL AND NEW.collection_time = NULL THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='collection_time', HINT = 'error_empty_override';
    END IF;

    IF OLD.communication_disability_level != '' AND NEW.communication_disability_level = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='communication_disability_level', HINT = 'error_empty_override';
    END IF;

    IF OLD.community_id != '' AND NEW.community_id = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='community_id', HINT = 'error_empty_override';
    END IF;

    IF OLD.displacement_status != '' AND NEW.displacement_status = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='displacement_status', HINT = 'error_empty_override';
    END IF;

    IF OLD.full_name != '' AND NEW.full_name = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='full_name', HINT = 'error_empty_override';
    END IF;

    IF OLD.sex != '' AND NEW.sex = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='sex', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_cognitive_disability != '' AND NEW.has_cognitive_disability = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_cognitive_disability', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_communication_disability != '' AND NEW.has_communication_disability = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_communication_disability', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_consented_to_referral != '' AND NEW.has_consented_to_referral = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_consented_to_referral', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_consented_to_rgpd != '' AND NEW.has_consented_to_rgpd = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_consented_to_rgpd', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_hearing_disability != '' AND NEW.has_hearing_disability = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_hearing_disability', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_mobility_disability != '' AND NEW.has_mobility_disability = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_mobility_disability', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_selfcare_disability != '' AND NEW.has_selfcare_disability = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_selfcare_disability', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_vision_disability != '' AND NEW.has_vision_disability = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_vision_disability', HINT = 'error_empty_override';
    END IF;

    IF OLD.hearing_disability_level != '' AND NEW.hearing_disability_level = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='hearing_disability_level', HINT = 'error_empty_override';
    END IF;

    IF OLD.household_id != '' AND NEW.household_id = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='household_id', HINT = 'error_empty_override';
    END IF;

    IF OLD.identification_type_1 != '' AND NEW.identification_type_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='identification_type_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.identification_type_explanation_1 != '' AND NEW.identification_type_explanation_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='identification_type_explanation_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.identification_number_1 != '' AND NEW.identification_number_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='identification_number_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.identification_type_2 != '' AND NEW.identification_type_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='identification_type_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.identification_type_explanation_2 != '' AND NEW.identification_type_explanation_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='identification_type_explanation_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.identification_number_2 != '' AND NEW.identification_number_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='identification_number_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.identification_type_3 != '' AND NEW.identification_type_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='identification_type_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.identification_type_explanation_3 != '' AND NEW.identification_type_explanation_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='identification_type_explanation_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.identification_number_3 != '' AND NEW.identification_number_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='identification_number_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.engagement_context != '' AND NEW.engagement_context = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='engagement_context', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_head_of_community != '' AND NEW.is_head_of_community = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_head_of_community', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_head_of_household != '' AND NEW.is_head_of_household = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_head_of_household', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_minor != '' AND NEW.is_minor = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_minor', HINT = 'error_empty_override';
    END IF;

    IF OLD.mobility_disability_level != '' AND NEW.mobility_disability_level = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='mobility_disability_level', HINT = 'error_empty_override';
    END IF;

    IF OLD.nationality_1 != '' AND NEW.nationality_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='nationality_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.nationality_2 != '' AND NEW.nationality_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='nationality_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.preferred_contact_method != '' AND NEW.preferred_contact_method = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='preferred_contact_method', HINT = 'error_empty_override';
    END IF;

    IF OLD.preferred_contact_method_comments != '' AND NEW.preferred_contact_method_comments = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='preferred_contact_method_comments', HINT = 'error_empty_override';
    END IF;

    IF OLD.preferred_name != '' AND NEW.preferred_name = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='preferred_name', HINT = 'error_empty_override';
    END IF;

    IF OLD.preferred_communication_language != '' AND NEW.preferred_communication_language = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='preferred_communication_language', HINT = 'error_empty_override';
    END IF;

    IF OLD.prefers_to_remain_anonymous != '' AND NEW.prefers_to_remain_anonymous = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='prefers_to_remain_anonymous', HINT = 'error_empty_override';
    END IF;

    IF OLD.presents_protection_concerns != '' AND NEW.presents_protection_concerns = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='presents_protection_concerns', HINT = 'error_empty_override';
    END IF;

    IF OLD.selfcare_disability_level != '' AND NEW.selfcare_disability_level = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='selfcare_disability_level', HINT = 'error_empty_override';
    END IF;

    IF OLD.spoken_language_1 != '' AND NEW.spoken_language_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='spoken_language_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.spoken_language_2 != '' AND NEW.spoken_language_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='spoken_language_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.spoken_language_3 != '' AND NEW.spoken_language_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='spoken_language_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.vision_disability_level != '' AND NEW.vision_disability_level = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='vision_disability_level', HINT = 'error_empty_override';
    END IF;

    IF OLD.age != NULL AND NEW.age = NULL THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='age', HINT = 'error_empty_override';
    END IF;

    IF OLD.phone_number_1 != '' AND NEW.phone_number_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='phone_number_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.phone_number_2 != '' AND NEW.phone_number_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='phone_number_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.phone_number_3 != '' AND NEW.phone_number_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='phone_number_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.email_1 != '' AND NEW.email_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='email_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.email_2 != '' AND NEW.email_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='email_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.email_3 != '' AND NEW.email_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='email_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.free_field_1 != '' AND NEW.free_field_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='free_field_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.free_field_2 != '' AND NEW.free_field_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='free_field_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.free_field_3 != '' AND NEW.free_field_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='free_field_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.free_field_4 != '' AND NEW.free_field_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='free_field_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.free_field_5 != '' AND NEW.free_field_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='free_field_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.comments != '' AND NEW.comments = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='comments', HINT = 'error_empty_override';
    END IF;

    IF OLD.displacement_status_comment != '' AND NEW.displacement_status_comment = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='displacement_status_comment', HINT = 'error_empty_override';
    END IF;

    IF OLD.collection_office != '' AND NEW.collection_office = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='collection_office', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_female_headed_household != '' AND NEW.is_female_headed_household = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_female_headed_household', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_minor_headed_household != '' AND NEW.is_minor_headed_household = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_minor_headed_household', HINT = 'error_empty_override';
    END IF;

    IF OLD.first_name != '' AND NEW.first_name = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='first_name', HINT = 'error_empty_override';
    END IF;

    IF OLD.middle_name != '' AND NEW.middle_name = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='middle_name', HINT = 'error_empty_override';
    END IF;

    IF OLD.last_name != '' AND NEW.last_name = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='last_name', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_1 != '' AND NEW.service_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_agent_name_1 != '' AND NEW.service_agent_name_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_agent_name_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_cc_1 != '' AND NEW.service_cc_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_cc_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_requested_date_1 != '' AND NEW.service_requested_date_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_requested_date_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_delivered_date_1 != '' AND NEW.service_delivered_date_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_delivered_date_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_comments_1 != '' AND NEW.service_comments_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_comments_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_type_1 != '' AND NEW.service_type_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_type_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_sub_service_1 != '' AND NEW.service_sub_service_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_sub_service_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_location_1 != '' AND NEW.service_location_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_location_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_donor_1 != '' AND NEW.service_donor_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_donor_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_project_name_1 != '' AND NEW.service_project_name_1 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_project_name_1', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_2 != '' AND NEW.service_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_agent_name_2 != '' AND NEW.service_agent_name_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_agent_name_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_cc_2 != '' AND NEW.service_cc_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_cc_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_requested_date_2 != '' AND NEW.service_requested_date_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_requested_date_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_delivered_date_2 != '' AND NEW.service_delivered_date_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_delivered_date_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_comments_2 != '' AND NEW.service_comments_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_comments_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_type_2 != '' AND NEW.service_type_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_type_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_sub_service_2 != '' AND NEW.service_sub_service_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_sub_service_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_location_2 != '' AND NEW.service_location_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_location_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_donor_2 != '' AND NEW.service_donor_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_donor_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_project_name_2 != '' AND NEW.service_project_name_2 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_project_name_2', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_3 != '' AND NEW.service_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_agent_name_3 != '' AND NEW.service_agent_name_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_agent_name_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_cc_3 != '' AND NEW.service_cc_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_cc_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_requested_date_3 != '' AND NEW.service_requested_date_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_requested_date_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_delivered_date_3 != '' AND NEW.service_delivered_date_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_delivered_date_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_comments_3 != '' AND NEW.service_comments_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_comments_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_type_3 != '' AND NEW.service_type_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_type_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_sub_service_3 != '' AND NEW.service_sub_service_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_sub_service_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_location_3 != '' AND NEW.service_location_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_location_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_donor_3 != '' AND NEW.service_donor_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_donor_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_project_name_3 != '' AND NEW.service_project_name_3 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_project_name_3', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_4 != '' AND NEW.service_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_agent_name_4 != '' AND NEW.service_agent_name_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_agent_name_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_cc_4 != '' AND NEW.service_cc_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_cc_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_requested_date_4 != '' AND NEW.service_requested_date_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_requested_date_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_delivered_date_4 != '' AND NEW.service_delivered_date_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_delivered_date_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_comments_4 != '' AND NEW.service_comments_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_comments_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_type_4 != '' AND NEW.service_type_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_type_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_sub_service_4 != '' AND NEW.service_sub_service_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_sub_service_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_location_4 != '' AND NEW.service_location_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_location_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_donor_4 != '' AND NEW.service_donor_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_donor_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_project_name_4 != '' AND NEW.service_project_name_4 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_project_name_4', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_5 != '' AND NEW.service_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_agent_name_5 != '' AND NEW.service_agent_name_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_agent_name_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_cc_5 != '' AND NEW.service_cc_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_cc_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_requested_date_5 != '' AND NEW.service_requested_date_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_requested_date_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_delivered_date_5 != '' AND NEW.service_delivered_date_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_delivered_date_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_comments_5 != '' AND NEW.service_comments_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_comments_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_type_5 != '' AND NEW.service_type_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_type_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_sub_service_5 != '' AND NEW.service_sub_service_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_sub_service_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_location_5 != '' AND NEW.service_location_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_location_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_donor_5 != '' AND NEW.service_donor_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_donor_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_project_name_5 != '' AND NEW.service_project_name_5 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_project_name_5', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_6 != '' AND NEW.service_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_agent_name_6 != '' AND NEW.service_agent_name_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_agent_name_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_cc_6 != '' AND NEW.service_cc_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_cc_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_requested_date_6 != '' AND NEW.service_requested_date_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_requested_date_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_delivered_date_6 != '' AND NEW.service_delivered_date_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_delivered_date_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_comments_6 != '' AND NEW.service_comments_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_comments_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_type_6 != '' AND NEW.service_type_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_type_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_sub_service_6 != '' AND NEW.service_sub_service_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_sub_service_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_location_6 != '' AND NEW.service_location_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_location_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_donor_6 != '' AND NEW.service_donor_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_donor_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_project_name_6 != '' AND NEW.service_project_name_6 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_project_name_6', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_7 != '' AND NEW.service_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_agent_name_7 != '' AND NEW.service_agent_name_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_agent_name_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_cc_7 != '' AND NEW.service_cc_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_cc_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_requested_date_7 != '' AND NEW.service_requested_date_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_requested_date_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_delivered_date_7 != '' AND NEW.service_delivered_date_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_delivered_date_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_comments_7 != '' AND NEW.service_comments_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_comments_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_type_7 != '' AND NEW.service_type_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_type_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_sub_service_7 != '' AND NEW.service_sub_service_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_sub_service_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_location_7 != '' AND NEW.service_location_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_location_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_donor_7 != '' AND NEW.service_donor_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_donor_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.service_project_name_7 != '' AND NEW.service_project_name_7 = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='service_project_name_7', HINT = 'error_empty_override';
    END IF;

    IF OLD.mothers_name != '' AND NEW.mothers_name = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='mothers_name', HINT = 'error_empty_override';
    END IF;

    IF OLD.household_size != NULL AND NEW.household_size = NULL THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='household_size', HINT = 'error_empty_override';
    END IF;

    IF OLD.community_size != NULL AND NEW.community_size = NULL THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='community_size', HINT = 'error_empty_override';
    END IF;

    IF OLD.native_name != '' AND NEW.native_name = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='native_name', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_disability != '' AND NEW.has_disability = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_disability', HINT = 'error_empty_override';
    END IF;

    IF OLD.pwd_comments != '' AND NEW.pwd_comments = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='pwd_comments', HINT = 'error_empty_override';
    END IF;

    IF OLD.has_medical_condition != '' AND NEW.has_medical_condition = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='has_medical_condition', HINT = 'error_empty_override';
    END IF;

    IF OLD.needs_legal_and_physical_protection != '' AND NEW.needs_legal_and_physical_protection = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='needs_legal_and_physical_protection', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_child_at_risk != '' AND NEW.is_child_at_risk = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_child_at_risk', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_woman_at_risk != '' AND NEW.is_woman_at_risk = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_woman_at_risk', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_elder_at_risk != '' AND NEW.is_elder_at_risk = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_elder_at_risk', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_pregnant != '' AND NEW.is_pregnant = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_pregnant', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_lactating != '' AND NEW.is_lactating = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_lactating', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_single_parent != '' AND NEW.is_single_parent = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_single_parent', HINT = 'error_empty_override';
    END IF;

    IF OLD.is_separated_child != '' AND NEW.is_separated_child = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='is_separated_child', HINT = 'error_empty_override';
    END IF;

    IF OLD.vulnerability_comments != '' AND NEW.vulnerability_comments = '' THEN
        RAISE EXCEPTION 'OVR' USING DETAIL = OLD.id, COLUMN  ='vulnerability_comments', HINT = 'error_empty_override';
    END IF;

    RETURN NEW;
END;$$
LANGUAGE plpgsql
;

CREATE OR REPLACE TRIGGER individual_registrations_no_empty_override
    BEFORE UPDATE
    ON individual_registrations
    FOR EACH ROW
EXECUTE PROCEDURE no_empty_override();

