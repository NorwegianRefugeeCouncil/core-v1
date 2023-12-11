package enumTypes

import (
	"github.com/nrc-no/notcore/internal/locales"
	"log"
	"os"
	"testing"
)

func setupSuite() func() {
	err := locales.LoadTranslations()
	if err == nil {
		locales.Init()
	}
	return func() {
		log.Println("teardown suite")
	}
}

func TestMain(m *testing.M) {
	teardownSuite := setupSuite()
	defer teardownSuite()
	code := m.Run()
	os.Exit(code)
}
