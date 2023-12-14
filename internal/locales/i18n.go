package locales

import (
	_ "embed"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"golang.org/x/text/language"
	"path/filepath"
	"strconv"
	"strings"
)

//go:embed en.toml
var localeEN string

//go:embed ja.toml
var localeJA string

var localeFiles = map[string]string{
	"en": localeEN,
	"ja": localeJA,
}

var localeNames = map[string]string{
	"en": "English",
	"ja": "Debugging",
}

var (
	DefaultLang    = language.English
	CurrentLang    = DefaultLang
	AvailableLangs = containers.NewStringSet()
)

var Translations *i18n.Bundle
var DefaultLocalizer *i18n.Localizer
var l *locales

type locales struct {
	localizer *i18n.Localizer
}

func Init() {
	loc := locales{localizer: DefaultLocalizer}
	l = &loc
}

func LoadTranslations() error {
	bundle := i18n.NewBundle(DefaultLang)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	for localeKey, locale := range localeFiles {
		AvailableLangs.Add(localeKey)
		_, err := bundle.ParseMessageFileBytes([]byte(locale), filepath.Join(localeKey+".toml"))
		if err != nil {
			return err
		}
	}

	Translations = bundle
	DefaultLocalizer = i18n.NewLocalizer(bundle, DefaultLang.String())
	return nil
}

func GetLocales() *locales {
	return l
}

func GetTranslator() func(id string, args ...interface{}) string {
	return l.Translate
}

func SetLocalizer(lang string) {
	loc := i18n.NewLocalizer(Translations, lang)
	CurrentLang = language.Make(lang)
	l.localizer = loc
}

func TranslateSlice(ids []string, args ...[]interface{}) []string {
	translations := make([]string, len(ids))
	for i := range ids {
		translations[i] = l.Translate(ids[i])
	}
	return translations
}

func GetTranslationKeys(values []string) []string {
	translationKeys := make([]string, len(values))
	for i, v := range values {
		for _, c := range constants.IndividualFileColumns {
			tra := l.Translate(c)
			val := strings.Trim(v, " \t\n\r")
			if tra == val {
				translationKeys[i] = c
			}
		}
	}
	return translationKeys
}

type Interface interface {
	Translate(id string, args ...interface{}) string
	TranslateCount(id string, ct int, args ...interface{}) string
	GetAvailableLocales() []string
}

type Translator func(id string, args ...interface{}) string

func (l locales) Translate(id string, args ...interface{}) string {
	var data = map[string]interface{}{}
	if len(args) > 0 {
		data = make(map[string]interface{}, len(args))
		for n, iface := range args {
			data["v"+strconv.Itoa(n)] = iface
		}
	}
	str, _, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
	})
	if str == "" && err != nil {
		return "[TL err: " + err.Error() + "]"
	}
	return str
}

func (l locales) TranslateCount(id string, ct int, args ...interface{}) string {
	data := make(map[string]interface{}, len(args)+1)
	if len(args) > 0 {
		for n, iface := range args {
			data["v"+strconv.Itoa(n)] = iface
		}
	}
	data["ct"] = ct
	str, _, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
		PluralCount:  ct,
	})
	if str == "" && err != nil {
		return "[TL err: " + err.Error() + "]"
	}
	return str
}

func (l locales) GetAvailableLocales() map[string]string {
	return localeNames
}
