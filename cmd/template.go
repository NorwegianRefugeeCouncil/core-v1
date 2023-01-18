/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Commented out fields are currently blank, but left in place in case we want to change values later
		individual := &api.Individual{
			Address:                       "123 Blvd. Drive",
			Age:                           pointers.Int(34),
			BirthDate:                     pointers.Time(time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC)),
			CognitiveDisabilityLevel:      api.DisabilityLevelMild,
			CollectionAdministrativeArea1: "Area1",
			CollectionAdministrativeArea2: "Area2",
			CollectionAdministrativeArea3: "Area3",
			CollectionOffice:              "Office",
			CollectionAgentName:           "Mary J.",
			CollectionAgentTitle:          "Field Officer",
			CollectionTime:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			Comments:                      "",
			CommunicationDisabilityLevel:  api.DisabilityLevelMild,
			CommunityID:                   "community-1234",
			CreatedAt:                     time.Now(),
			DisplacementStatus:            api.DisplacementStatusIDP,
			DisplacementStatusComment:     "",
			Email1:                        "john@email.com",
			Email2:                        "",
			Email3:                        "",
			FullName:                      "John Doe",
			FreeField1:                    "",
			FreeField2:                    "",
			FreeField3:                    "",
			FreeField4:                    "",
			FreeField5:                    "",
			Sex:                           api.SexMale,
			HasCognitiveDisability:        true,
			HasCommunicationDisability:    true,
			HasConsentedToRGPD:            true,
			HasConsentedToReferral:        true,
			HasHearingDisability:          false,
			HasMobilityDisability:         false,
			HasSelfCareDisability:         false,
			HasVisionDisability:           false,
			// HearingDisabilityLevel:         nil,
			HouseholdID:                    "household-1234",
			ID:                             "",
			IdentificationType1:            "passport",
			IdentificationTypeExplanation1: "",
			IdentificationNumber1:          "123456789",
			// IdentificationType2:            nil,
			// IdentificationTypeExplanation2: nil,
			// IdentificationNumber2:          nil,
			// IdentificationType3:            nil,
			// IdentificationTypeExplanation3: nil,
			// IdentificationNumber3:          nil,
			EngagementContext:       "fieldActivity",
			InternalID:              "Internal-id-1234",
			IsHeadOfCommunity:       false,
			IsHeadOfHousehold:       true,
			IsFemaleHeadedHousehold: false,
			IsMinorHeadedHousehold:  false,
			IsMinor:                 false,
			// MobilityDisabilityLevel:        nil,
			Nationality1:                   "AFG",
			Nationality2:                   "",
			PhoneNumber1:                   "123-456-1233",
			PhoneNumber2:                   "",
			PhoneNumber3:                   "",
			PreferredContactMethod:         "phone",
			PreferredContactMethodComments: "",
			PreferredName:                  "John",
			PreferredCommunicationLanguage: "fra",
			PrefersToRemainAnonymous:       false,
			PresentsProtectionConcerns:     false,
			// SelfCareDisabilityLevel:        nil,
			SpokenLanguage1: "fra",
			SpokenLanguage2: "en",
			SpokenLanguage3: "",
			UpdatedAt:       time.Now(),
			// VisionDisabilityLevel:          nil,
		}
		individualList := []*api.Individual{individual}
		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)
		templateFile, err := os.OpenFile(path.Join(basepath, "..", "web", "static", "nrc_grf_template.xlsx"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}
		defer func() {
			templateFile.Close()
		}()
		if err := api.MarshalIndividualsExcel(templateFile, individualList); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
