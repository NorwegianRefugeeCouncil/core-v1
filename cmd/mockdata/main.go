package mockdata

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/manveru/faker"
	"github.com/rs/xid"
)

func Generate() error {

	f, err := faker.New("en")
	if err != nil {
		return err
	}

	file, err := os.Create("generated.csv")
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	if err := writer.Write([]string{
		"id",
		"full_name",
		"preferred_name",
		"email",
		"address",
		"phone_number",
		"birth_date",
		"is_minor",
		"gender",
		"presents_protection_concerns",
		"physical_impairment",
		"sensory_impairment",
		"mental_impairment",
		"displacement_status",
		"countries",
	}); err != nil {
		return err
	}
	for i := 0; i < 200; i++ {

		var name = f.Name()
		var preferredName = name
		if randBool(5) {
			preferredName = f.Name()
		}

		var email string
		if randBool(80) {
			email = f.Email()
		}

		var address string
		if randBool(80) {
			address = f.StreetAddress() + "\n" + f.PostCode() + " " + f.City() + "\n" + f.Country()
		}
		var phoneNumber string
		if randBool(80) {
			phoneNumber = f.PhoneNumber()
		}

		var birthDate string
		if randBool(80) {
			start, _ := time.Parse("2006-01-02", "1900-01-01")
			end := time.Now()
			birthDate = start.Add(time.Duration(rand.Int63n(end.Unix()-start.Unix())) * time.Second).Format("2006-01-02")
		}

		var isMinor = "false"
		if randBool(20) {
			isMinor = "true"
		}

		var gender string
		if randBool(45) {
			gender = "male"
		} else if randBool(45) {
			gender = "female"
		}

		var protectionConcerns = "false"
		if randBool(20) {
			protectionConcerns = "true"
		}

		var physicalImpairment = ""
		if randBool(10) {
			physicalImpairment = "moderate"
		} else if randBool(5) {
			physicalImpairment = "severe"
		}

		var sensoryImpairment = ""
		if randBool(10) {
			sensoryImpairment = "moderate"
		} else if randBool(5) {
			sensoryImpairment = "severe"
		}

		var mentalImpairment = ""
		if randBool(10) {
			mentalImpairment = "moderate"
		} else if randBool(5) {
			mentalImpairment = "severe"
		}

		var displacementStatus = ""
		if randBool(20) {
			displacementStatus = "refugee"
		} else if randBool(20) {
			displacementStatus = "idp"
		} else if randBool(20) {
			displacementStatus = "host_community"
		}

		var countries = []string{}
		if randBool(20) {
			countries = append(countries, "kenya")
		}
		if randBool(20) {
			countries = append(countries, "uganda")
		}
		if randBool(20) {
			countries = append(countries, "colombia")
		}
		if randBool(20) {
			countries = append(countries, "ukraine")
		}
		var countriesStr = ""
		if len(countries) > 0 {
			countriesStr = strings.Join(countries, ",")
		}

		if err := writer.Write([]string{
			xid.New().String(),
			name,
			preferredName,
			email,
			address,
			phoneNumber,
			birthDate,
			isMinor,
			gender,
			protectionConcerns,
			physicalImpairment,
			sensoryImpairment,
			mentalImpairment,
			displacementStatus,
			countriesStr,
		}); err != nil {
			return err
		}

	}

	writer.Flush()
	return nil

}

func randBool(probability int) bool {
	return rand.Intn(100) < probability
}
