/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/locales"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Create the default template for users to download",
	Long:  `Create an excel file that contains an example participant that users can download in the app`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Commented out fields are currently blank, but left in place in case we want to change values later
		individual := &api.Individual{
			Address:                         "123 Blvd. Drive",
			Age:                             pointers.Int(34),
			BirthDate:                       pointers.Time(time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC)),
			CognitiveDisabilityLevel:        enumTypes.DisabilityLevelMild,
			CollectionAdministrativeArea1:   "Area1",
			CollectionAdministrativeArea2:   "Area2",
			CollectionAdministrativeArea3:   "Area3",
			CollectionOffice:                "Office",
			CollectionAgentName:             "Mary J.",
			CollectionAgentTitle:            "Field Officer",
			CollectionTime:                  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			Comments:                        "",
			CommunicationDisabilityLevel:    enumTypes.DisabilityLevelMild,
			CommunityID:                     "community-1234",
			CommunitySize:                   pointers.Int(5),
			CreatedAt:                       time.Now(),
			DisplacementStatus:              enumTypes.DisplacementStatusIDP,
			DisplacementStatusComment:       "",
			Email1:                          "john@email.com",
			Email2:                          "",
			Email3:                          "",
			FullName:                        "John Joe Doe",
			FirstName:                       "John",
			MiddleName:                      "Joe",
			LastName:                        "Doe",
			NativeName:                      "جون",
			MothersName:                     "Jane Doe",
			FreeField1:                      "",
			FreeField2:                      "",
			FreeField3:                      "",
			FreeField4:                      "",
			FreeField5:                      "",
			Sex:                             enumTypes.SexMale,
			HasMedicalCondition:             nil,
			NeedsLegalAndPhysicalProtection: pointers.Bool(false),
			IsSeparatedChild:                nil,
			IsSingleParent:                  pointers.Bool(true),
			IsPregnant:                      pointers.Bool(false),
			IsLactating:                     pointers.Bool(true),
			IsChildAtRisk:                   nil,
			IsElderAtRisk:                   nil,
			IsWomanAtRisk:                   pointers.Bool(false),
			HasCognitiveDisability:          pointers.Bool(true),
			HasCommunicationDisability:      pointers.Bool(true),
			HasConsentedToRGPD:              pointers.Bool(true),
			HasConsentedToReferral:          nil,
			HasDisability:                   pointers.Bool(false),
			HasHearingDisability:            pointers.Bool(false),
			HasMobilityDisability:           pointers.Bool(false),
			HasSelfCareDisability:           pointers.Bool(false),
			HasVisionDisability:             nil,
			HouseholdID:                     "household-1234",
			HouseholdSize:                   pointers.Int(5),
			ID:                              "",
			IdentificationType1:             enumTypes.IdentificationTypePassport,
			IdentificationTypeExplanation1:  "",
			IdentificationNumber1:           "123456789",
			EngagementContext:               enumTypes.EngagementContextFieldActivity,
			InternalID:                      "Internal-id-1234",
			IsHeadOfCommunity:               pointers.Bool(false),
			IsHeadOfHousehold:               pointers.Bool(true),
			IsFemaleHeadedHousehold:         pointers.Bool(false),
			IsMinorHeadedHousehold:          pointers.Bool(false),
			IsMinor:                         pointers.Bool(false),
			Nationality1:                    "AFG",
			Nationality2:                    "",
			PhoneNumber1:                    "123-456-1233",
			PhoneNumber2:                    "",
			PhoneNumber3:                    "",
			PreferredContactMethod:          enumTypes.ContactMethodPhone,
			PreferredContactMethodComments:  "",
			PreferredName:                   "John",
			PreferredCommunicationLanguage:  "fra",
			PrefersToRemainAnonymous:        pointers.Bool(false),
			PresentsProtectionConcerns:      pointers.Bool(false),
			PWDComments:                     "",
			VulnerabilityComments:           "",
			SpokenLanguage1:                 "fra",
			SpokenLanguage2:                 "eng",
			SpokenLanguage3:                 "",
			UpdatedAt:                       time.Now(),
			ServiceCC1:                      enumTypes.ServiceCCShelter,
			ServiceRequestedDate1:           pointers.Time(time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)),
			ServiceDeliveredDate1:           pointers.Time(time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)),
			ServiceComments1:                "Service comment",
		}
		individualList := []*api.Individual{individual}

		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)

		locales.LoadTranslations()
		locales.Init()

		for _, lang := range locales.AvailableLangs.Items() {
			locales.SetLocalizer(lang)

			templateFile, err := os.OpenFile(path.Join(basepath, "..", "web", "static", fmt.Sprintf("nrc_grf_template.%s.xlsx", lang)), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				return err
			}
			defer func() {
				templateFile.Close()
			}()
			if err := api.MarshalIndividualsExcel(templateFile, individualList); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
