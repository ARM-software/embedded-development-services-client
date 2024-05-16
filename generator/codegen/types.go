package codegen

import (
	validation "github.com/go-ozzo/ozzo-validation"

	configUtils "github.com/ARM-software/golang-utils/utils/config"
)

type Collection struct {
	CollectionRef string
	ItemRef       string
	ModelRef      string
	IteratorRef   string
}

type Collections = []Collection

type MessageCollection struct {
	ItemRef     string
	IteratorRef string
}

type MessageCollections []MessageCollection

type NotificationFeedCollection struct {
	ItemRef     string
	IteratorRef string
}

type NotificationFeedCollections []NotificationFeedCollection

type CollectionParams = struct {
	Collections
	MessageCollections
	NotificationFeedCollections
}

type Data struct {
	Params          any
	SpecPath        string
	TemplatePath    string
	DestinationPath string
}

type ExtensionsConfig struct {
	Input    string `mapstructure:"input"`
	Output   string `mapstructure:"output"`
	Template string `mapstructure:"template"`
}

func DefaultExtensionsConfig() *ExtensionsConfig {
	return &ExtensionsConfig{
		Input:    "",
		Output:   "",
		Template: "templates/entities.go.tmpl",
	}
}

func (cfg *ExtensionsConfig) Validate() error {
	err := configUtils.ValidateEmbedded(cfg)
	if err != nil {
		return err
	}

	return validation.ValidateStruct(cfg,
		validation.Field(&cfg.Input, validation.Required),
		validation.Field(&cfg.Output, validation.Required),
		validation.Field(&cfg.Template, validation.Required),
	)
}
