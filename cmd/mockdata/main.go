package mockdata

import (
	"encoding/csv"
	"math/rand"
	"os"
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
		"email",
		"address",
		"phone_number",
		"birth_date",
		"gender",
	}); err != nil {
		return err
	}
	for i := 0; i < 5000; i++ {

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

		var gender string
		if randBool(45) {
			gender = "male"
		} else if randBool(45) {
			gender = "female"
		}

		if err := writer.Write([]string{
			xid.New().String(),
			f.Name(),
			email,
			address,
			phoneNumber,
			birthDate,
			gender,
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
