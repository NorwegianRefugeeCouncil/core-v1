package locales

import (
	"context"
	_ "embed"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/nrc-no/notcore/internal/containers"
	"golang.org/x/text/language"
	"path/filepath"
	"strconv"
)

//go:embed en.toml
var localeEN string

//go:embed ja.toml
var localeJA string

var localeFiles = map[string]string{
	"en": localeEN,
	"ja": localeJA,
}

var (
	DefaultLang = language.English
)

var Translations *i18n.Bundle
var DefaultLocalizer *i18n.Localizer

var AvailableLangs = containers.NewStringSet()

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

type Interface interface {
	Translate(id string, args ...interface{}) string
	TranslateCount(id string, ct int, args ...interface{}) string
}

type locales struct {
	localizer *i18n.Localizer
}

func New(ctx context.Context) Interface {
	localizer, ok := localizerFrom(ctx)
	if !ok {
		localizer = DefaultLocalizer
	}
	return locales{localizer: localizer}
}

func (l locales) Translate(id string, args ...interface{}) string {
	var data map[string]interface{}
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

type localizerKey struct{}

func WithLocalizer(ctx context.Context, loc *i18n.Localizer) context.Context {
	return context.WithValue(ctx, localizerKey{}, loc)
}

func localizerFrom(ctx context.Context) (*i18n.Localizer, bool) {
	loc, ok := ctx.Value(localizerKey{}).(*i18n.Localizer)
	return loc, ok
}
