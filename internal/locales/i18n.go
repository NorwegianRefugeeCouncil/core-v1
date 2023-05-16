package locales

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/nrc-no/notcore/internal/containers"
	"golang.org/x/text/language"
)

var (
	DefaultLang = language.English
)

var Translations *i18n.Bundle
var DefaultLocalizer *i18n.Localizer

var AvailableLangs = containers.NewStringSet()

func LoadTranslations() error {
	dir := filepath.Join("internal", "locales")
	bundle := i18n.NewBundle(DefaultLang)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".toml" {
			return nil
		}
		AvailableLangs.Add(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
		_, err = bundle.LoadMessageFile(path)
		return err
	})
	if err != nil {
		return err
	}

	Translations = bundle
	DefaultLocalizer = i18n.NewLocalizer(bundle, DefaultLang.String())
	return nil
}

type Interface interface {
	Translate(id string, args ...interface{}) string
	TranslateCount(id string, ct int, args ...interface{}) string
	GetAvailableLangs() []string
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

func (l locales) GetAvailableLangs() []string {
	return AvailableLangs.Items()
}

type localizerKey struct{}

func WithLocalizer(ctx context.Context, loc *i18n.Localizer) context.Context {
	return context.WithValue(ctx, localizerKey{}, loc)
}

func localizerFrom(ctx context.Context) (*i18n.Localizer, bool) {
	loc, ok := ctx.Value(localizerKey{}).(*i18n.Localizer)
	return loc, ok
}
