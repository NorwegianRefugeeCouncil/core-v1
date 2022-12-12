package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/nrc-no/notcore/pkg/views/forms"
	"time"
)

type Cache interface {
	UpdateLanguageOptions(options []forms.SelectInputFieldOption) error
	ReadLanguageOptions() ([]forms.SelectInputFieldOption, error)
	DeleteLanguageOptions()
}

type cache struct {
	languageOptions *bigcache.BigCache
}

func (c *cache) UpdateLanguageOptions(options []forms.SelectInputFieldOption) error {
	bytes, err := json.Marshal(&options)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	return c.languageOptions.Set("languageOptions", bytes)
}

func (c *cache) ReadLanguageOptions() ([]forms.SelectInputFieldOption, error) {
	bytes, err := c.languageOptions.Get("languageOptions")
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return []forms.SelectInputFieldOption{}, err
		}

		return []forms.SelectInputFieldOption{}, fmt.Errorf("get: %w", err)
	}

	var options []forms.SelectInputFieldOption
	err = json.Unmarshal(bytes, &options)
	if err != nil {
		return []forms.SelectInputFieldOption{}, fmt.Errorf("unmarshal: %w", err)
	}

	return options, nil
}

func (c *cache) DeleteLanguageOptions() {
	c.languageOptions.Delete("languageOptions")
}

func NewCache() (Cache, error) {
	newCache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(24*time.Hour))
	if err != nil {
		return nil, fmt.Errorf("new big cache: %w", err)
	}

	return &cache{
		languageOptions: newCache,
	}, nil
}
