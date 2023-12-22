package locales

import (
	_ "embed"
	"fmt"
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

//go:embed ar.toml
var localeAR string

var localeFiles = map[string]string{
	"en": localeEN,
	"ja": localeJA,
	"ar": localeAR,
}

var (
	AvailableLangs = containers.NewStringSet()
	CurrentLang    = DefaultLang
	DefaultLang    = language.English
)
var Localizers map[language.Tag]*i18n.Localizer
var l *locales

type locales struct {
	localizers    map[language.Tag]*i18n.Localizer
	currentLocale string
}

func Init() {
	loc := locales{localizers: Localizers, currentLocale: DefaultLang.String()}
	l = &loc
}

func LoadTranslations() error {
	locs := map[language.Tag]*i18n.Localizer{}

	for localeKey, locale := range localeFiles {
		localeTag := language.Make(localeKey)
		bundle := i18n.NewBundle(localeTag)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		AvailableLangs.Add(localeKey)
		_, err := bundle.ParseMessageFileBytes([]byte(locale), filepath.Join(localeKey+".toml"))
		if err != nil {
			return err
		}
		locs[localeTag] = i18n.NewLocalizer(bundle, locale)
	}

	Localizers = locs
	return nil
}

func GetLocales() *locales {
	return l
}

func GetTranslator() func(id string, args ...interface{}) string {
	return l.Translate
}

func SetLocalizer(lang string) {
	CurrentLang = language.Make(lang)
}

func TranslateSlice(ids []string, args ...[]interface{}) []string {
	translations := make([]string, len(ids))
	for i := range ids {
		translations[i] = l.Translate(ids[i])
	}
	return translations
}

func GetDBColumns(values []string) ([]string, []error) {
	dbCols := make([]string, len(values))
	errs := []error{}
	for i, v := range values {
		foundKey := false
		val := strings.Trim(v, " \t\n\r")
		for _, c := range constants.IndividualFileColumns {
			for _, lang := range AvailableLangs.Items() {
				tra := l.TranslateFrom(c, lang)
				if tra == val {
					dbCols[i] = constants.IndividualFileToDBMap[c]
					foundKey = true
					continue
				}
			}
			if foundKey {
				break
			}
		}
		if !foundKey {
			if constants.IndividualDBColumns.Contains(val) {
				dbCols[i] = val
			} else {
				errs = append(errs, fmt.Errorf(l.Translate("error_unknown_column_detail", val)))
			}
		}
	}
	return dbCols, errs
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
	str, _, err := l.localizers[CurrentLang].LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
	})
	if str == "" && err != nil {
		return "[TL err: " + err.Error() + "]"
	}
	return str
}

func (l locales) TranslateFrom(id string, lang string, args ...interface{}) string {
	var data = map[string]interface{}{}
	if len(args) > 0 {
		data = make(map[string]interface{}, len(args))
		for n, iface := range args {
			data["v"+strconv.Itoa(n)] = iface
		}
	}
	str, _, err := l.localizers[language.Make(lang)].LocalizeWithTag(&i18n.LocalizeConfig{
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
	str, _, err := l.localizers[CurrentLang].LocalizeWithTag(&i18n.LocalizeConfig{
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
	availableLocales := make(map[string]string, len(l.localizers))
	for loc, _ := range localeFiles {
		availableLocales[loc] = l.TranslateFrom("language", loc)
	}
	return availableLocales
}
