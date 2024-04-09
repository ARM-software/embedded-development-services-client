package codegen

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"path/filepath"
	"text/template"

	"github.com/ARM-software/golang-utils/utils/commonerrors"
	configUtils "github.com/ARM-software/golang-utils/utils/config"
	"github.com/ARM-software/golang-utils/utils/filesystem"

	"github.com/getkin/kin-openapi/openapi3"
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/tools/imports"
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

type ExtensionsConfig struct {
	Input    string `mapstructure:"input"`
	Output   string `mapstructure:"output"`
	Template string `mapstructure:"template"`
}

func DefaultExtensionsConfig() *ExtensionsConfig {
	return &ExtensionsConfig{
		Input:    "",
		Output:   "",
		Template: "templates/collections.go.tmpl",
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

func GenerateTemplateFile(d *Data) (err error) {
	t, err := template.
		New(filepath.Base(d.TemplatePath)).
		ParseFS(templates, d.TemplatePath)
	if err != nil {
		return
	}

	err = generateSourceCode(d, t)
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

func generateSourceCode(data *Data, template *template.Template) (err error) {
	if data == nil {
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

func LoadAPIDefinition(specPath string) (swagger *openapi3.T, err error) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	swagger, err = loader.LoadFromFile(specPath)
	return
}
