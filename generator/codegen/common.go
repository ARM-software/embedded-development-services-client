package codegen

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"go/format"
	"path/filepath"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/tools/imports"

	"github.com/ARM-software/golang-utils/utils/commonerrors"
	configUtils "github.com/ARM-software/golang-utils/utils/config"
	"github.com/ARM-software/golang-utils/utils/filesystem"
	"github.com/ARM-software/golang-utils/utils/parallelisation"
	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	//go:embed templates/*
	templates embed.FS
)

type Data struct {
	Params          any
	SpecPath        string
	TemplatePath    string
	DestinationPath string
}

var ValidGenerationTargets = map[string]func(*Data) error{
	"collections": AddCollectionsToParams,
	"jobs":        AddJobItemsToParams,
	"links":       GenerateFollowLinks,
}

type ExtensionsConfig struct {
	Input        string `mapstructure:"input"`
	Output       string `mapstructure:"output"`
	Template     string `mapstructure:"template"`
	GenerateType string `mapstructure:"generate_type"`
}

func DefaultExtensionsConfig() *ExtensionsConfig {
	return &ExtensionsConfig{
		Input:    "",
		Output:   "",
		Template: "",
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

func GenerateDataStruct(cfg ExtensionsConfig) (d *Data, err error) {
	return &Data{
		DestinationPath: cfg.Output,
		SpecPath:        cfg.Input,
		TemplatePath:    cfg.Template,
	}, nil
}

func GenerateTemplateFile(ctx context.Context, d *Data) (err error) {
	t, err := template.
		New(filepath.Base(d.TemplatePath)).
		ParseFS(templates, d.TemplatePath)
	if err != nil {
		return
	}

	err = generateSourceCode(ctx, d, t)
	return
}

func formatImports(path string, src []byte) (bytes []byte, err error) {
	bytes, err = imports.Process(path, src, &imports.Options{
		Comments:   true,
		TabIndent:  true,
		TabWidth:   4,
		Fragment:   true,
		FormatOnly: false,
	})
	if err != nil {
		err = fmt.Errorf("%w: could not cleanup imports for file '%v': %v", commonerrors.ErrUnexpected, path, err.Error())
		return
	}

	return
}

func generateSourceCode(ctx context.Context, data *Data, template *template.Template) (err error) {
	if data == nil {
		err = fmt.Errorf("%w: data must be defined", commonerrors.ErrUndefined)
		return
	}
	if template == nil {
		err = fmt.Errorf("%w: template must be defined", commonerrors.ErrUndefined)
		return
	}

	err = parallelisation.DetermineContextError(ctx)
	if err != nil {
		return
	}

	var tempBuffer bytes.Buffer
	err = template.Execute(&tempBuffer, data.Params)
	if err != nil {
		return
	}
	content, err := format.Source(tempBuffer.Bytes())
	if err != nil {
		return
	}
	content, err = formatImports(data.DestinationPath, content)
	if err != nil {
		return
	}

	file, err := filesystem.CreateFile(data.DestinationPath)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()
	_, err = file.Write(content)
	if err != nil {
		return
	}

	err = file.Close()
	return
}

func loadAPIDefinition(specPath string) (swagger *openapi3.T, err error) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	swagger, err = loader.LoadFromFile(specPath)
	return
}

func AddValuesToParams(d *Data, getValues func(*openapi3.T) (interface{}, error), generationType string) (err error) {
	if d == nil {
		err = fmt.Errorf("%w input variable", commonerrors.ErrUndefined)
		return
	}
	if d.TemplatePath == "" {
		d.TemplatePath = fmt.Sprintf("templates/%s.tmpl", generationType)
	}
	swagger, err := loadAPIDefinition(d.SpecPath)
	if err != nil {
		return
	}
	params, err := getValues(swagger)
	if err != nil {
		return
	}
	d.Params = params
	return nil
}
