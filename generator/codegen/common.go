package codegen

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"go/format"
	"path/filepath"
	"slices"
	"strings"
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
	//go:embed static/*
	static embed.FS
)

const (
	schemaPrefix = "#/components/schemas/"
	jsonMIME     = "application/json"
	// See https://github.com/Arm-Debug/API-Uniform-Contract?tab=readme-ov-file#api-extensions
	redactFlag       = "x-redact"
	jobFlag          = "x-job"
	collectionFlag   = "x-collection"
	noPaginationFlag = "x-no-pagination"
)

type Data struct {
	Params            any
	SpecPath          string
	TemplatePath      string
	DestinationPath   string
	ClientPackagePath string
	PackageName       string
	DisableLinks      bool
}

var ValidGenerationTargets = map[string]func(*Data) error{
	"collections": AddCollectionsToParams,
	"jobs":        AddJobItemsToParams,
	"links":       AddLinkFollowersToParams,
}

type ExtensionsConfig struct {
	Input             string `mapstructure:"input"`
	Output            string `mapstructure:"output"`
	Template          string `mapstructure:"template"`
	GenerateType      string `mapstructure:"generate_type"`
	ClientPackagePath string `mapstructure:"client_path"`
	DisableLinks      bool   `mapstructure:"disable_links"`
}

type StaticFileConfig struct {
	ClientName  string
	Destination string
	Exclusions  []string
}

func DefaultExtensionsConfig() *ExtensionsConfig {
	return &ExtensionsConfig{
		Input:             "",
		Output:            "",
		Template:          "",
		GenerateType:      "",
		ClientPackagePath: "",
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

func getPackageNameFromPath(clientPath string) (clientName string, err error) {
	if clientPath == "" {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "missing client package path, provide a client package path using '-c' or '--client_path'")
		return
	}
	isDir, err := filesystem.IsDir(clientPath)
	if err != nil {
		return
	}

	if !isDir {
		err = commonerrors.Newf(commonerrors.ErrInvalid, "client path '%s' is not a directory", clientPath)
		return
	}

	clientName = filepath.Base(clientPath)

	return
}

func GenerateDataStruct(cfg ExtensionsConfig) (d *Data, err error) {
	clientName, err := getPackageNameFromPath(cfg.ClientPackagePath)
	if err != nil {
		return
	}

	return &Data{
		DestinationPath:   cfg.Output,
		SpecPath:          cfg.Input,
		TemplatePath:      cfg.Template,
		ClientPackagePath: cfg.ClientPackagePath,
		PackageName:       clientName,
	}, nil
}

func GenerateStaticFileConfigStruct(cfg ExtensionsConfig) (d *StaticFileConfig, err error) {
	clientName, err := getPackageNameFromPath(cfg.ClientPackagePath)
	if err != nil {
		return
	}

	var exclusions []string
	if cfg.DisableLinks {
		exclusions = append(exclusions, "static/extension_model_hal_link_data.go.static")
	}

	return &StaticFileConfig{
		ClientName:  clientName,
		Destination: filepath.Dir(cfg.Output),
		Exclusions:  exclusions,
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
	err = template.Execute(&tempBuffer, data)
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

func isExtensionFlagSet(props openapi3.ExtensionProps, flagKey string) (isSet bool, err error) {
	if c, ok := props.Extensions[flagKey].(json.RawMessage); ok {
		marshallingErr := json.Unmarshal(c, &isSet)
		if marshallingErr != nil {
			err = commonerrors.Newf(commonerrors.ErrUnexpected, "failed to unmarshal `%s`", flagKey)
			return
		}
	}
	return
}

func CopyStaticFiles(ctx context.Context, staticFileConfig *StaticFileConfig) (err error) {
	efs, fsErr := filesystem.NewEmbedFileSystem(&static)
	if fsErr != nil {
		return commonerrors.Newf(commonerrors.ErrUnexpected, "failed to create a filesystem for directory `%s`: %s", staticFileConfig.Destination, fsErr.Error())
	}

	files, lsErr := efs.FindAll(".", "go.static")
	if lsErr != nil {
		return commonerrors.Newf(commonerrors.ErrUnexpected, "no files with the '.go.static' extension were found in the directory `%s`", staticFileConfig.Destination)
	}

	mkdirErr := filesystem.MkDir(staticFileConfig.Destination)
	if mkdirErr != nil {
		return commonerrors.Newf(commonerrors.ErrUnexpected, "could not create directory `%s`: %s", staticFileConfig.Destination, mkdirErr.Error())
	}

	for _, f := range files {
		// Skip the file if it's contained in exclusions
		if slices.Contains(staticFileConfig.Exclusions, f) {
			continue
		}

		t, tmplErr := template.
			New(filesystem.FilePathBase(efs, f)).
			ParseFS(static, f)
		if tmplErr != nil {
			return tmplErr
		}

		resultFileName := strings.TrimPrefix(f, "static/") // The embedded filesystem always uses Unix-style paths, regardless of the target platform
		resultFileName = strings.TrimSuffix(resultFileName, ".static")

		d := Data{
			DestinationPath: filepath.Join(staticFileConfig.Destination, resultFileName),
			PackageName:     staticFileConfig.ClientName,
		}
		err = generateSourceCode(ctx, &d, t)
		if err != nil {
			return
		}

	}

	return
}
