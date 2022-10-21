package constants

const (
	FormParamIndividualAddress                    = "Address"
	FormParamIndividualBirthDate                  = "BirthDate"
	FormParamIndividualCountry                    = "Country"
	FormParamIndividualDisplacementStatus         = "DisplacementStatus"
	FormParamIndividualEmail                      = "Email"
	FormParamIndividualFullName                   = "FullName"
	FormParamIndividualGender                     = "Gender"
	FormParamIndividualIsMinor                    = "IsMinor"
	FormParamIndividualMentalImpairment           = "MentalImpairment"
	FormParamIndividualPhoneNumber                = "PhoneNumber"
	FormParamIndividualPhysicalImpairment         = "PhysicalImpairment"
	FormParamIndividualPreferredName              = "PreferredName"
	FormParamIndividualPresentsProtectionConcerns = "PresentsProtectionConcerns"
	FormParamIndividualSensoryImpairment          = "SensoryImpairment"

	FormParamGetIndividualsEmail               = "email"
	FormParamGetIndividualsName                = "name"
	FormParamGetIndividualsPhoneNumber         = "phone_number"
	FormParamsGetIndividualsAddress            = "address"
	FormParamsGetIndividualsGender             = "gender"
	FormParamsGetIndividualsIsMinor            = "is_minor"
	FormParamsGetIndividualsProtectionConcerns = "presents_protection_concerns"
	FormParamsGetIndividualsAgeFrom            = "age_from"
	FormParamsGetIndividualsAgeTo              = "age_to"
	FormParamGetIndividualsCountryID           = "country_id"
	FormParamsGetIndividualsDisplacementStatus = "displacement_status"

	DBColumnIndividualID                         = "id"
	DBColumnIndividualAddress                    = "address"
	DBColumnIndividualBirthDate                  = "birth_date"
	DBColumnIndividualCountryID                  = "country_id"
	DBColumnIndividualDisplacementStatus         = "displacement_status"
	DBColumnIndividualEmail                      = "email"
	DBColumnIndividualFullName                   = "full_name"
	DBColumnIndividualGender                     = "gender"
	DBColumnIndividualIsMinor                    = "is_minor"
	DBColumnIndividualMentalImpairment           = "mental_impairment"
	DBColumnIndividualPhoneNumber                = "phone_number"
	DBColumnIndividualPhysicalImpairment         = "physical_impairment"
	DBColumnIndividualPreferredName              = "preferred_name"
	DBColumnIndividualPresentsProtectionConcerns = "presents_protection_concerns"
	DBColumnIndividualSensoryImpairment          = "sensory_impairment"
	DBColumnIndividualNormalizedPhoneNumber      = "normalized_phone_number"

	FileColumnIndividualID                         = "id"
	FileColumnIndividualAddress                    = "address"
	FileColumnIndividualBirthDate                  = "birth_date"
	FileColumnIndividualDisplacementStatus         = "displacement_status"
	FileColumnIndividualEmail                      = "email"
	FileColumnIndividualFullName                   = "full_name"
	FileColumnIndividualGender                     = "gender"
	FileColumnIndividualIsMinor                    = "is_minor"
	FileColumnIndividualMentalImpairment           = "mental_impairment"
	FileColumnIndividualPhoneNumber                = "phone_number"
	FileColumnIndividualPhysicalImpairment         = "physical_impairment"
	FileColumnIndividualPreferredName              = "preferred_name"
	FileColumnIndividualPresentsProtectionConcerns = "presents_protection_concerns"
	FileColumnIndividualSensoryImpairment          = "sensory_impairment"
)

var IndividualDBColumns = []string{
	DBColumnIndividualID,
	DBColumnIndividualAddress,
	DBColumnIndividualBirthDate,
	DBColumnIndividualCountryID,
	DBColumnIndividualDisplacementStatus,
	DBColumnIndividualEmail,
	DBColumnIndividualFullName,
	DBColumnIndividualGender,
	DBColumnIndividualIsMinor,
	DBColumnIndividualMentalImpairment,
	DBColumnIndividualPhoneNumber,
	DBColumnIndividualPhysicalImpairment,
	DBColumnIndividualPreferredName,
	DBColumnIndividualPresentsProtectionConcerns,
	DBColumnIndividualSensoryImpairment,
}

var IndividualFileColumns = []string{
	FileColumnIndividualID,
	FileColumnIndividualAddress,
	FileColumnIndividualBirthDate,
	FileColumnIndividualDisplacementStatus,
	FileColumnIndividualEmail,
	FileColumnIndividualFullName,
	FileColumnIndividualGender,
	FileColumnIndividualIsMinor,
	FileColumnIndividualMentalImpairment,
	FileColumnIndividualPhoneNumber,
	FileColumnIndividualPhysicalImpairment,
	FileColumnIndividualPreferredName,
	FileColumnIndividualPresentsProtectionConcerns,
	FileColumnIndividualSensoryImpairment,
}

var IndividualDBToFileMap = map[string]string{
	DBColumnIndividualID:                         FileColumnIndividualID,
	DBColumnIndividualAddress:                    FileColumnIndividualAddress,
	DBColumnIndividualBirthDate:                  FileColumnIndividualBirthDate,
	DBColumnIndividualDisplacementStatus:         FileColumnIndividualDisplacementStatus,
	DBColumnIndividualEmail:                      FileColumnIndividualEmail,
	DBColumnIndividualFullName:                   FileColumnIndividualFullName,
	DBColumnIndividualGender:                     FileColumnIndividualGender,
	DBColumnIndividualIsMinor:                    FileColumnIndividualIsMinor,
	DBColumnIndividualMentalImpairment:           FileColumnIndividualMentalImpairment,
	DBColumnIndividualPhoneNumber:                FileColumnIndividualPhoneNumber,
	DBColumnIndividualPhysicalImpairment:         FileColumnIndividualPhysicalImpairment,
	DBColumnIndividualPreferredName:              FileColumnIndividualPreferredName,
	DBColumnIndividualPresentsProtectionConcerns: FileColumnIndividualPresentsProtectionConcerns,
	DBColumnIndividualSensoryImpairment:          FileColumnIndividualSensoryImpairment,
}

var IndividualFileToDBMap = map[string]string{
	FileColumnIndividualID:                         DBColumnIndividualID,
	FileColumnIndividualAddress:                    DBColumnIndividualAddress,
	FileColumnIndividualBirthDate:                  DBColumnIndividualBirthDate,
	FileColumnIndividualDisplacementStatus:         DBColumnIndividualDisplacementStatus,
	FileColumnIndividualEmail:                      DBColumnIndividualEmail,
	FileColumnIndividualFullName:                   DBColumnIndividualFullName,
	FileColumnIndividualGender:                     DBColumnIndividualGender,
	FileColumnIndividualIsMinor:                    DBColumnIndividualIsMinor,
	FileColumnIndividualMentalImpairment:           DBColumnIndividualMentalImpairment,
	FileColumnIndividualPhoneNumber:                DBColumnIndividualPhoneNumber,
	FileColumnIndividualPhysicalImpairment:         DBColumnIndividualPhysicalImpairment,
	FileColumnIndividualPreferredName:              DBColumnIndividualPreferredName,
	FileColumnIndividualPresentsProtectionConcerns: DBColumnIndividualPresentsProtectionConcerns,
	FileColumnIndividualSensoryImpairment:          DBColumnIndividualSensoryImpairment,
}
