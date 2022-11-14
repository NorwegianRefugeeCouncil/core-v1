package validation

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/stretchr/testify/assert"
)

type IndividualBuilder struct {
	individual *api.Individual
}

func NewIndividualBuilder() *IndividualBuilder {
	return &IndividualBuilder{
		individual: &api.Individual{},
	}
}

func (i *IndividualBuilder) Build() *api.Individual {
	return i.individual
}

func (i *IndividualBuilder) WithAddress(address string) *IndividualBuilder {
	i.individual.Address = address
	return i
}

func (i *IndividualBuilder) WithAge(age *int) *IndividualBuilder {
	i.individual.Age = age
	return i
}

func (i *IndividualBuilder) WithBirthDate(birthDate *time.Time) *IndividualBuilder {
	i.individual.BirthDate = birthDate
	return i
}

func (i *IndividualBuilder) WithCognitiveDisabilityLevel(level api.DisabilityLevel) *IndividualBuilder {
	i.individual.CognitiveDisabilityLevel = level
	return i
}

func (i *IndividualBuilder) WithCollectionAdministrativeArea1(area string) *IndividualBuilder {
	i.individual.CollectionAdministrativeArea1 = area
	return i
}

func (i *IndividualBuilder) WithCollectionAdministrativeArea2(area string) *IndividualBuilder {
	i.individual.CollectionAdministrativeArea2 = area
	return i
}

func (i *IndividualBuilder) WithCollectionAdministrativeArea3(area string) *IndividualBuilder {
	i.individual.CollectionAdministrativeArea3 = area
	return i
}

func (i *IndividualBuilder) WithCollectionAgentName(collectionAgentName string) *IndividualBuilder {
	i.individual.CollectionAgentName = collectionAgentName
	return i
}

func (i *IndividualBuilder) WithCollectionAgentTitle(collectionAgentTitle string) *IndividualBuilder {
	i.individual.CollectionAgentTitle = collectionAgentTitle
	return i
}

func (i *IndividualBuilder) WithCollectionTime(t time.Time) *IndividualBuilder {
	i.individual.CollectionTime = t
	return i
}

func (i *IndividualBuilder) WithComments(comments string) *IndividualBuilder {
	i.individual.Comments = comments
	return i
}

func (i *IndividualBuilder) WithCommunicationDisabilityLevel(level api.DisabilityLevel) *IndividualBuilder {
	i.individual.CommunicationDisabilityLevel = level
	return i
}

func (i *IndividualBuilder) WithCommunityID(communityID string) *IndividualBuilder {
	i.individual.CommunityID = communityID
	return i
}

func (i *IndividualBuilder) WithCountryID(countryID string) *IndividualBuilder {
	i.individual.CountryID = countryID
	return i
}

func (i *IndividualBuilder) WithCreatedAt(createdAt time.Time) *IndividualBuilder {
	i.individual.CreatedAt = createdAt
	return i
}

func (i *IndividualBuilder) WithDeletedAt(deletedAt *time.Time) *IndividualBuilder {
	i.individual.DeletedAt = deletedAt
	return i
}

func (i *IndividualBuilder) WithDisplacementStatus(displacementStatus api.DisplacementStatus) *IndividualBuilder {
	i.individual.DisplacementStatus = displacementStatus
	return i
}

func (i *IndividualBuilder) WithEmail1(email string) *IndividualBuilder {
	i.individual.Email1 = email
	return i
}

func (i *IndividualBuilder) WithEmail2(email string) *IndividualBuilder {
	i.individual.Email2 = email
	return i
}

func (i *IndividualBuilder) WithEmail3(email string) *IndividualBuilder {
	i.individual.Email3 = email
	return i
}

func (i *IndividualBuilder) WithFullName(fullName string) *IndividualBuilder {
	i.individual.FullName = fullName
	return i
}

func (i *IndividualBuilder) WithFreeField1(freeField string) *IndividualBuilder {
	i.individual.FreeField1 = freeField
	return i
}

func (i *IndividualBuilder) WithFreeField2(freeField string) *IndividualBuilder {
	i.individual.FreeField2 = freeField
	return i
}

func (i *IndividualBuilder) WithFreeField3(freeField string) *IndividualBuilder {
	i.individual.FreeField3 = freeField
	return i
}

func (i *IndividualBuilder) WithFreeField4(freeField string) *IndividualBuilder {
	i.individual.FreeField4 = freeField
	return i
}

func (i *IndividualBuilder) WithFreeField5(freeField string) *IndividualBuilder {
	i.individual.FreeField5 = freeField
	return i
}

func (i *IndividualBuilder) WithSex(sex api.Sex) *IndividualBuilder {
	i.individual.Sex = sex
	return i
}

func (i *IndividualBuilder) WithHasCognitiveDisability(hasCognitiveDisability bool) *IndividualBuilder {
	i.individual.HasCognitiveDisability = hasCognitiveDisability
	return i
}

func (i *IndividualBuilder) WithHasCommunicationDisability(hasCommunicationDisability bool) *IndividualBuilder {
	i.individual.HasCommunicationDisability = hasCommunicationDisability
	return i
}

func (i *IndividualBuilder) WithHasConsentedToRGPD(hasConsentedToRGPD bool) *IndividualBuilder {
	i.individual.HasConsentedToRGPD = hasConsentedToRGPD
	return i
}

func (i *IndividualBuilder) WithHasConsentedToReferral(hasConsentedToReferral bool) *IndividualBuilder {
	i.individual.HasConsentedToReferral = hasConsentedToReferral
	return i
}

func (i *IndividualBuilder) WithHasHearingDisability(hasHearingDisability bool) *IndividualBuilder {
	i.individual.HasHearingDisability = hasHearingDisability
	return i
}

func (i *IndividualBuilder) WithHasMobilityDisability(hasMobilityDisability bool) *IndividualBuilder {
	i.individual.HasMobilityDisability = hasMobilityDisability
	return i
}

func (i *IndividualBuilder) WithHasSelfCareDisability(hasSelfCareDisability bool) *IndividualBuilder {
	i.individual.HasSelfCareDisability = hasSelfCareDisability
	return i
}

func (i *IndividualBuilder) WithHasVisionDisability(hasVisionDisability bool) *IndividualBuilder {
	i.individual.HasVisionDisability = hasVisionDisability
	return i
}

func (i *IndividualBuilder) WithHearingDisabilityLevel(level api.DisabilityLevel) *IndividualBuilder {
	i.individual.HearingDisabilityLevel = level
	return i
}

func (i *IndividualBuilder) WithHouseholdID(householdID string) *IndividualBuilder {
	i.individual.HouseholdID = householdID
	return i
}

func (i *IndividualBuilder) WithID(id string) *IndividualBuilder {
	i.individual.ID = id
	return i
}

func (i *IndividualBuilder) WithIdentificationType1(identificationType string) *IndividualBuilder {
	i.individual.IdentificationType1 = identificationType
	return i
}

func (i *IndividualBuilder) WithIdentificationTypeExplanation1(identificationTypeExplanation string) *IndividualBuilder {
	i.individual.IdentificationTypeExplanation1 = identificationTypeExplanation
	return i
}

func (i *IndividualBuilder) WithIdentificationNumber1(identificationNumber string) *IndividualBuilder {
	i.individual.IdentificationNumber1 = identificationNumber
	return i
}

func (i *IndividualBuilder) WithIdentificationType2(identificationType string) *IndividualBuilder {
	i.individual.IdentificationType2 = identificationType
	return i
}

func (i *IndividualBuilder) WithIdentificationTypeExplanation2(identificationTypeExplanation string) *IndividualBuilder {
	i.individual.IdentificationTypeExplanation2 = identificationTypeExplanation
	return i
}

func (i *IndividualBuilder) WithIdentificationNumber2(identificationNumber string) *IndividualBuilder {
	i.individual.IdentificationNumber2 = identificationNumber
	return i
}

func (i *IndividualBuilder) WithIdentificationType3(identificationType string) *IndividualBuilder {
	i.individual.IdentificationType3 = identificationType
	return i
}

func (i *IndividualBuilder) WithIdentificationTypeExplanation3(identificationTypeExplanation string) *IndividualBuilder {
	i.individual.IdentificationTypeExplanation3 = identificationTypeExplanation
	return i
}

func (i *IndividualBuilder) WithIdentificationNumber3(identificationNumber string) *IndividualBuilder {
	i.individual.IdentificationNumber3 = identificationNumber
	return i
}

func (i *IndividualBuilder) WithEngagementContext(engagementContext string) *IndividualBuilder {
	i.individual.EngagementContext = engagementContext
	return i
}

func (i *IndividualBuilder) WithInternalID(internalID string) *IndividualBuilder {
	i.individual.InternalID = internalID
	return i
}

func (i *IndividualBuilder) WithIsMinor(isMinor bool) *IndividualBuilder {
	i.individual.IsMinor = isMinor
	return i
}

func (i *IndividualBuilder) WithMobilityDisabilityLevel(level api.DisabilityLevel) *IndividualBuilder {
	i.individual.MobilityDisabilityLevel = level
	return i
}

func (i *IndividualBuilder) WithNationality1(nationality string) *IndividualBuilder {
	i.individual.Nationality1 = nationality
	return i
}

func (i *IndividualBuilder) WithNationality2(nationality string) *IndividualBuilder {
	i.individual.Nationality2 = nationality
	return i
}

func (i *IndividualBuilder) WithNormalizedPhoneNumber1(phoneNumber string) *IndividualBuilder {
	i.individual.NormalizedPhoneNumber1 = phoneNumber
	return i
}

func (i *IndividualBuilder) WithNormalizedPhoneNumber2(phoneNumber string) *IndividualBuilder {
	i.individual.NormalizedPhoneNumber2 = phoneNumber
	return i
}

func (i *IndividualBuilder) WithNormalizedPhoneNumber3(phoneNumber string) *IndividualBuilder {
	i.individual.NormalizedPhoneNumber3 = phoneNumber
	return i
}

func (i *IndividualBuilder) WithPhoneNumber1(phoneNumber string) *IndividualBuilder {
	i.individual.PhoneNumber1 = phoneNumber
	return i
}

func (i *IndividualBuilder) WithPhoneNumber2(phoneNumber string) *IndividualBuilder {
	i.individual.PhoneNumber2 = phoneNumber
	return i
}

func (i *IndividualBuilder) WithPhoneNumber3(phoneNumber string) *IndividualBuilder {
	i.individual.PhoneNumber3 = phoneNumber
	return i
}

func (i *IndividualBuilder) WithPreferredContactMethod(method string) *IndividualBuilder {
	i.individual.PreferredContactMethod = method
	return i
}

func (i *IndividualBuilder) WithPreferredContactMethodComments(comments string) *IndividualBuilder {
	i.individual.PreferredContactMethodComments = comments
	return i
}

func (i *IndividualBuilder) WithPreferredName(preferredName string) *IndividualBuilder {
	i.individual.PreferredName = preferredName
	return i
}

func (i *IndividualBuilder) WithPreferredCommunicationLanguage(language string) *IndividualBuilder {
	i.individual.PreferredCommunicationLanguage = language
	return i
}

func (i *IndividualBuilder) WithPrefersToRemainAnonymous(prefersToRemainAnonymous bool) *IndividualBuilder {
	i.individual.PrefersToRemainAnonymous = prefersToRemainAnonymous
	return i
}

func (i *IndividualBuilder) WithPresentsProtectionConcerns(presentsProtectionConcerns bool) *IndividualBuilder {
	i.individual.PresentsProtectionConcerns = presentsProtectionConcerns
	return i
}

func (i *IndividualBuilder) WithSelfCareDisabilityLevel(level api.DisabilityLevel) *IndividualBuilder {
	i.individual.SelfCareDisabilityLevel = level
	return i
}

func (i *IndividualBuilder) WithSpokenLanguage1(language string) *IndividualBuilder {
	i.individual.SpokenLanguage1 = language
	return i
}

func (i *IndividualBuilder) WithSpokenLanguage2(language string) *IndividualBuilder {
	i.individual.SpokenLanguage2 = language
	return i
}

func (i *IndividualBuilder) WithSpokenLanguage3(language string) *IndividualBuilder {
	i.individual.SpokenLanguage3 = language
	return i
}

func (i *IndividualBuilder) WithUpdatedAt(updatedAt time.Time) *IndividualBuilder {
	i.individual.UpdatedAt = updatedAt
	return i
}

func (i *IndividualBuilder) WithVisionDisabilityLevel(level api.DisabilityLevel) *IndividualBuilder {
	i.individual.VisionDisabilityLevel = level
	return i
}

func ValidIndividual() *IndividualBuilder {
	bd := time.Now().AddDate(-10, 0, 0)
	return NewIndividualBuilder().
		WithEmail1("email@email.com").
		WithPhoneNumber1("1234567890").
		WithFullName("John Doe").
		WithDisplacementStatus("idp").
		WithBirthDate(&bd).
		WithCountryID(uuid.New().String()).
		WithPreferredName("John").
		WithSex("male").
		WithCollectionAgentTitle("Collection Agent Title").
		WithCollectionAgentName("Collection Agent Name").
		WithCollectionTime(time.Now()).
		WithPreferredContactMethod("email")
}

func TestValidateIndividual(t *testing.T) {
	futureDate := time.Now().AddDate(1, 0, 0)
	emptyDate := time.Time{}
	email1Path := validation.NewPath("email1")
	email2Path := validation.NewPath("email2")
	email3Path := validation.NewPath("email3")
	birthDatePath := validation.NewPath("birthDate")
	countryIDPath := validation.NewPath("countryId")
	displacementStatusPath := validation.NewPath("displacementStatus")
	sexPath := validation.NewPath("sex")
	tests := []struct {
		name string
		i    *api.Individual
		want validation.ErrorList
	}{
		{
			name: "valid",
			i:    ValidIndividual().Build(),
			want: validation.ErrorList{},
		}, {
			name: "address (too long)",
			i:    ValidIndividual().WithAddress(bigstr(individualAddressMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("address"), bigstr(individualAddressMaxLength+1), individualAddressMaxLength)},
		}, {
			name: "birthdate (in the future)",
			i:    ValidIndividual().WithBirthDate(&futureDate).Build(),
			want: validation.ErrorList{validation.Invalid(validation.NewPath("birthDate"), &futureDate, "birthdate cannot be in the future")},
		}, {
			name: "birthdate (empty)",
			i:    ValidIndividual().WithBirthDate(&emptyDate).Build(),
			want: validation.ErrorList{validation.Invalid(birthDatePath, &emptyDate, "must be a valid date")},
		}, {
			name: "cognitiveDisabilityLevel (invalid)",
			i:    ValidIndividual().WithCognitiveDisabilityLevel("invalid").Build(),
			want: validation.ErrorList{validation.NotSupported(validation.NewPath("cognitiveDisabilityLevel"), api.DisabilityLevel("invalid"), allowedDisabilityLevelsStr)},
		}, {
			name: "collectionAdministrativeArea1 (too long)",
			i:    ValidIndividual().WithCollectionAdministrativeArea1(bigstr(individualAdministrativeAreaMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("collectionAdministrativeArea1"), bigstr(individualAdministrativeAreaMaxLength+1), individualAdministrativeAreaMaxLength)},
		}, {
			name: "collectionAdministrativeArea2 (too long)",
			i:    ValidIndividual().WithCollectionAdministrativeArea2(bigstr(individualAdministrativeAreaMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("collectionAdministrativeArea2"), bigstr(individualAdministrativeAreaMaxLength+1), individualAdministrativeAreaMaxLength)},
		}, {
			name: "collectionAdministrativeArea3 (too long)",
			i:    ValidIndividual().WithCollectionAdministrativeArea3(bigstr(individualAdministrativeAreaMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("collectionAdministrativeArea3"), bigstr(individualAdministrativeAreaMaxLength+1), individualAdministrativeAreaMaxLength)},
		}, {
			name: "collectionAgentName (too long)",
			i:    ValidIndividual().WithCollectionAgentName(bigstr(individualCollectionAgentNameMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("collectionAgentName"), bigstr(individualCollectionAgentNameMaxLength+1), individualCollectionAgentNameMaxLength)},
		}, {
			name: "collectionAgentTitle (too long)",
			i:    ValidIndividual().WithCollectionAgentTitle(bigstr(individualCollectionAgentTitleMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("collectionAgentTitle"), bigstr(individualCollectionAgentTitleMaxLength+1), individualCollectionAgentTitleMaxLength)},
		}, {
			name: "collectionTime (in the future)",
			i:    ValidIndividual().WithCollectionTime(futureDate).Build(),
			want: validation.ErrorList{validation.Invalid(validation.NewPath("collectionTime"), futureDate, "collection time cannot be in the future")},
		}, {
			name: "comments (too long)",
			i:    ValidIndividual().WithComments(bigstr(maxTextLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("comments"), bigstr(maxTextLength+1), maxTextLength)},
		}, {
			name: "communicationDisabilityLevel (invalid)",
			i:    ValidIndividual().WithCommunicationDisabilityLevel("invalid").Build(),
			want: validation.ErrorList{validation.NotSupported(validation.NewPath("communicationDisabilityLevel"), api.DisabilityLevel("invalid"), allowedDisabilityLevelsStr)},
		}, {
			name: "communityId (too long)",
			i:    ValidIndividual().WithCommunityID(bigstr(individualCommunityIDMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("communityId"), bigstr(individualCommunityIDMaxLength+1), individualCommunityIDMaxLength)},
		}, {
			name: "countryId (not a uuid)",
			i:    ValidIndividual().WithCountryID("not a uuid").Build(),
			want: validation.ErrorList{validation.Invalid(countryIDPath, "not a uuid", "must be a valid UUID")},
		}, {
			name: "countryId (empty)",
			i:    ValidIndividual().WithCountryID("").Build(),
			want: validation.ErrorList{validation.Required(countryIDPath, "country id is required")},
		}, {
			name: "displacementStatus (invalid)",
			i:    ValidIndividual().WithDisplacementStatus("invalid").Build(),
			want: validation.ErrorList{validation.NotSupported(displacementStatusPath, api.DisplacementStatus("invalid"), allowedDisplacementStatusesStr)},
		}, {
			name: "email1 (too long)",
			i:    ValidIndividual().WithEmail1(bigstr(individualEmailMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(email1Path, bigstr(individualEmailMaxLength+1), individualEmailMaxLength)},
		}, {
			name: "email1 (invalid)",
			i:    ValidIndividual().WithEmail1("invalid").Build(),
			want: validation.ErrorList{validation.Invalid(email1Path, "invalid", "invalid email address")},
		}, {
			name: "email2 (too long)",
			i:    ValidIndividual().WithEmail2(bigstr(individualEmailMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(email2Path, bigstr(individualEmailMaxLength+1), individualEmailMaxLength)},
		}, {
			name: "email2 (invalid)",
			i:    ValidIndividual().WithEmail2("invalid").Build(),
			want: validation.ErrorList{validation.Invalid(email2Path, "invalid", "invalid email address")},
		}, {
			name: "email3 (too long)",
			i:    ValidIndividual().WithEmail3(bigstr(individualEmailMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(email3Path, bigstr(individualEmailMaxLength+1), individualEmailMaxLength)},
		}, {
			name: "email3 (invalid)",
			i:    ValidIndividual().WithEmail3("invalid").Build(),
			want: validation.ErrorList{validation.Invalid(email3Path, "invalid", "invalid email address")},
		}, {
			name: "fullName (too long)",
			i:    ValidIndividual().WithFullName(bigstr(individualFullNameMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("fullName"), bigstr(individualFullNameMaxLength+1), individualFullNameMaxLength)},
		}, {
			name: "freeField1 (too long)",
			i:    ValidIndividual().WithFreeField1(bigstr(individualFreeFieldMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("freeField1"), bigstr(individualFreeFieldMaxLength+1), individualFreeFieldMaxLength)},
		}, {
			name: "freeField2 (too long)",
			i:    ValidIndividual().WithFreeField2(bigstr(individualFreeFieldMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("freeField2"), bigstr(individualFreeFieldMaxLength+1), individualFreeFieldMaxLength)},
		}, {
			name: "freeField3 (too long)",
			i:    ValidIndividual().WithFreeField3(bigstr(individualFreeFieldMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("freeField3"), bigstr(individualFreeFieldMaxLength+1), individualFreeFieldMaxLength)},
		}, {
			name: "freeField4 (too long)",
			i:    ValidIndividual().WithFreeField4(bigstr(individualFreeFieldMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("freeField4"), bigstr(individualFreeFieldMaxLength+1), individualFreeFieldMaxLength)},
		}, {
			name: "freeField5 (too long)",
			i:    ValidIndividual().WithFreeField5(bigstr(individualFreeFieldMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("freeField5"), bigstr(individualFreeFieldMaxLength+1), individualFreeFieldMaxLength)},
		}, {
			name: "sex (invalid)",
			i:    ValidIndividual().WithSex("bla").Build(),
			want: validation.ErrorList{validation.NotSupported(sexPath, api.Sex("bla"), allowedSexsStr)},
		}, {
			name: "hearingDisabilityLevel (invalid)",
			i:    ValidIndividual().WithHearingDisabilityLevel("invalid").Build(),
			want: validation.ErrorList{validation.NotSupported(validation.NewPath("hearingDisabilityLevel"), api.DisabilityLevel("invalid"), allowedDisabilityLevelsStr)},
		}, {
			name: "householdId (too long)",
			i:    ValidIndividual().WithHouseholdID(bigstr(individualHouseholdIDMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("householdId"), bigstr(individualHouseholdIDMaxLength+1), individualHouseholdIDMaxLength)},
		}, {
			name: "identificationType1 (too long)",
			i:    ValidIndividual().WithIdentificationType1(bigstr(individualIdentificationTypeMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("identificationType1"), bigstr(individualIdentificationTypeMaxLength+1), individualIdentificationTypeMaxLength)},
		}, {
			name: "identificationTypeExplanation1 (too long)",
			i:    ValidIndividual().WithIdentificationTypeExplanation1(bigstr(maxTextLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("identificationTypeExplanation1"), bigstr(maxTextLength+1), maxTextLength)},
		}, {
			name: "identificationNumber1 (too long)",
			i:    ValidIndividual().WithIdentificationNumber1(bigstr(individualIdentificationNumberMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("identificationNumber1"), bigstr(individualIdentificationNumberMaxLength+1), individualIdentificationNumberMaxLength)},
		}, {
			name: "identificationType2 (too long)",
			i:    ValidIndividual().WithIdentificationType2(bigstr(individualIdentificationTypeMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("identificationType2"), bigstr(individualIdentificationTypeMaxLength+1), individualIdentificationTypeMaxLength)},
		}, {
			name: "identificationTypeExplanation2 (too long)",
			i:    ValidIndividual().WithIdentificationTypeExplanation2(bigstr(maxTextLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("identificationTypeExplanation2"), bigstr(maxTextLength+1), maxTextLength)},
		}, {
			name: "identificationNumber2 (too long)",
			i:    ValidIndividual().WithIdentificationNumber2(bigstr(individualIdentificationNumberMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("identificationNumber2"), bigstr(individualIdentificationNumberMaxLength+1), individualIdentificationNumberMaxLength)},
		}, {
			name: "identificationType3 (too long)",
			i:    ValidIndividual().WithIdentificationType3(bigstr(individualIdentificationTypeMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("identificationType3"), bigstr(individualIdentificationTypeMaxLength+1), individualIdentificationTypeMaxLength)},
		}, {
			name: "identificationTypeExplanation3 (too long)",
			i:    ValidIndividual().WithIdentificationTypeExplanation3(bigstr(maxTextLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("identificationTypeExplanation3"), bigstr(maxTextLength+1), maxTextLength)},
		}, {
			name: "identificationNumber3 (too long)",
			i:    ValidIndividual().WithIdentificationNumber3(bigstr(individualIdentificationNumberMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("identificationNumber3"), bigstr(individualIdentificationNumberMaxLength+1), individualIdentificationNumberMaxLength)},
		}, {
			name: "engagementContext (too long)",
			i:    ValidIndividual().WithEngagementContext(bigstr(individualEngagementContextMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("engagementContext"), bigstr(individualEngagementContextMaxLength+1), individualEngagementContextMaxLength)},
		}, {
			name: "internalId (too long)",
			i:    ValidIndividual().WithInternalID(bigstr(individualInternalIDMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("internalId"), bigstr(individualInternalIDMaxLength+1), individualInternalIDMaxLength)},
		}, {
			name: "mobilityDisabilityLevel (invalid)",
			i:    ValidIndividual().WithMobilityDisabilityLevel("invalid").Build(),
			want: validation.ErrorList{validation.NotSupported(validation.NewPath("mobilityDisabilityLevel"), api.DisabilityLevel("invalid"), allowedDisabilityLevelsStr)},
		}, {
			name: "nationality1 (too long)",
			i:    ValidIndividual().WithNationality1(bigstr(individualNationalityMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("nationality1"), bigstr(individualNationalityMaxLength+1), individualNationalityMaxLength)},
		}, {
			name: "nationality2 (too long)",
			i:    ValidIndividual().WithNationality2(bigstr(individualNationalityMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("nationality2"), bigstr(individualNationalityMaxLength+1), individualNationalityMaxLength)},
		}, {
			name: "phoneNumber1 (too long)",
			i:    ValidIndividual().WithPhoneNumber1(bigstr(individualPhoneNumberMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("phoneNumber1"), bigstr(individualPhoneNumberMaxLength+1), individualPhoneNumberMaxLength)},
		}, {
			name: "phoneNumber2 (too long)",
			i:    ValidIndividual().WithPhoneNumber2(bigstr(individualPhoneNumberMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("phoneNumber2"), bigstr(individualPhoneNumberMaxLength+1), individualPhoneNumberMaxLength)},
		}, {
			name: "phoneNumber3 (too long)",
			i:    ValidIndividual().WithPhoneNumber3(bigstr(individualPhoneNumberMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("phoneNumber3"), bigstr(individualPhoneNumberMaxLength+1), individualPhoneNumberMaxLength)},
		}, {
			name: "preferredContactMethod (too long)",
			i:    ValidIndividual().WithPreferredContactMethod(bigstr(individualPreferredContactMethodMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("preferredContactMethod"), bigstr(individualPreferredContactMethodMaxLength+1), individualPreferredContactMethodMaxLength)},
		}, {
			name: "preferredContactMethodComments (too long)",
			i:    ValidIndividual().WithPreferredContactMethodComments(bigstr(maxTextLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("preferredContactMethodComments"), bigstr(maxTextLength+1), maxTextLength)},
		}, {
			name: "preferredName (too long)",
			i:    ValidIndividual().WithPreferredName(bigstr(individualPreferredNameMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("preferredName"), bigstr(individualPreferredNameMaxLength+1), individualPreferredNameMaxLength)},
		}, {
			name: "preferredCommunicationLanguage (too long)",
			i:    ValidIndividual().WithPreferredCommunicationLanguage(bigstr(individualPreferredCommunicationLanguageMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("preferredCommunicationLanguage"), bigstr(individualPreferredCommunicationLanguageMaxLength+1), individualPreferredCommunicationLanguageMaxLength)},
		}, {
			name: "selfCareDisabilityLevel (invalid)",
			i:    ValidIndividual().WithSelfCareDisabilityLevel("invalid").Build(),
			want: validation.ErrorList{validation.NotSupported(validation.NewPath("selfCareDisabilityLevel"), api.DisabilityLevel("invalid"), allowedDisabilityLevelsStr)},
		}, {
			name: "spokenLanguage1 (too long)",
			i:    ValidIndividual().WithSpokenLanguage1(bigstr(individualSpokenLanguageMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("spokenLanguage1"), bigstr(individualSpokenLanguageMaxLength+1), individualSpokenLanguageMaxLength)},
		}, {
			name: "spokenLanguage2 (too long)",
			i:    ValidIndividual().WithSpokenLanguage2(bigstr(individualSpokenLanguageMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("spokenLanguage2"), bigstr(individualSpokenLanguageMaxLength+1), individualSpokenLanguageMaxLength)},
		}, {
			name: "spokenLanguage3 (too long)",
			i:    ValidIndividual().WithSpokenLanguage3(bigstr(individualSpokenLanguageMaxLength + 1)).Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(validation.NewPath("spokenLanguage3"), bigstr(individualSpokenLanguageMaxLength+1), individualSpokenLanguageMaxLength)},
		}, {
			name: "visionDisabilityLevel (invalid)",
			i:    ValidIndividual().WithVisionDisabilityLevel("invalid").Build(),
			want: validation.ErrorList{validation.NotSupported(validation.NewPath("visionDisabilityLevel"), api.DisabilityLevel("invalid"), allowedDisabilityLevelsStr)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateIndividual(tt.i)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValidateIndividualList(t *testing.T) {
	tests := []struct {
		name string
		i    *api.IndividualList
		want validation.ErrorList
	}{
		{
			name: "valid",
			i: &api.IndividualList{
				Items: []*api.Individual{ValidIndividual().Build()},
			},
			want: validation.ErrorList{},
		}, {
			name: "invalid",
			i: &api.IndividualList{
				Items: []*api.Individual{ValidIndividual().WithDisplacementStatus("bla").Build()},
			},
			want: validation.ErrorList{
				validation.NotSupported(
					validation.NewPath("items[0].displacementStatus"),
					api.DisplacementStatus("bla"),
					allowedDisplacementStatusesStr)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateIndividualList(tt.i)
			assert.Equal(t, tt.want, got)
		})
	}
}
