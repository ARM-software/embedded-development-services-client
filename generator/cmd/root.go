package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ARM-software/embedded-development-services-client/generator/codegen"
	"github.com/ARM-software/golang-utils/utils/commonerrors"
	configUtils "github.com/ARM-software/golang-utils/utils/config"
	"github.com/ARM-software/golang-utils/utils/logs/logrimp"
)

const (
	app             = "generate_code"
	inputArg        = "input"
	outputArg       = "output"
	templateArg     = "template"
	GenerateType    = "generate"
	clientPath      = "client-path"
	disableLinksArg = "disable-links"
)

var (
	extensionConfig codegen.ExtensionsConfig
	viperSession    = viper.New()
)

var rootCmd = &cobra.Command{
	Use: "extension-generator",
	RunE: func(_ *cobra.Command, _ []string) error {
		return RunCLI(context.Background())
	},
	SilenceUsage: true, // otherwise 'Usage' is printed after any error
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP(inputArg, "i", "", "path to the OpenAPI spec")
	rootCmd.Flags().StringP(outputArg, "o", "", "destination to output the file to")
	rootCmd.Flags().StringP(templateArg, "t", "", "path to the template file if defaults need overriding")
	rootCmd.Flags().StringP(GenerateType, "g", "", "what to generate")
	rootCmd.Flags().StringP(clientPath, "c", "", "path to the client go package generated by OpenAPI Generator")
	rootCmd.Flags().StringP(disableLinksArg, "d", "", "disable the generation of HAL link data")

	_ = configUtils.BindFlagToEnv(viperSession, app, "GENERATE_CODE_INPUT", rootCmd.Flags().Lookup(inputArg))
	_ = configUtils.BindFlagToEnv(viperSession, app, "GENERATE_CODE_OUTPUT", rootCmd.Flags().Lookup(outputArg))
	_ = configUtils.BindFlagToEnv(viperSession, app, "GENERATE_CODE_TEMPLATE", rootCmd.Flags().Lookup(templateArg))
	_ = configUtils.BindFlagToEnv(viperSession, app, "GENERATE_CODE_GENERATE_TYPE", rootCmd.Flags().Lookup(GenerateType))
	_ = configUtils.BindFlagToEnv(viperSession, app, "GENERATE_CODE_CLIENT_PATH", rootCmd.Flags().Lookup(clientPath))
	_ = configUtils.BindFlagToEnv(viperSession, app, "GENERATE_CODE_DISABLE_LINKS", rootCmd.Flags().Lookup(disableLinksArg))
}

func RunCLI(ctx context.Context) (err error) {
	logger := logrimp.NewStdOutLogr()

	err = configUtils.LoadFromViper(viperSession, app, &extensionConfig, codegen.DefaultExtensionsConfig())
	if err != nil {
		logger.Error(err, "Failed to initialise CLI with error")
		return
	}

	d, err := codegen.GenerateDataStruct(extensionConfig)
	if err != nil {
		return
	}

	generateParams, ok := codegen.ValidGenerationTargets[extensionConfig.GenerateType]
	if ok {
		err = generateParams(d)
	} else {
		err = fmt.Errorf("%w. '%s' is not a valid generation type", commonerrors.ErrNotFound, extensionConfig.GenerateType)
	}
	if err != nil {
		return
	}

	staticFileConfig, err := codegen.GenerateStaticFileConfigStruct(extensionConfig)
	if err != nil {
		return
	}

	err = codegen.CopyStaticFiles(ctx, staticFileConfig)
	if err != nil {
		return
	}

	err = codegen.GenerateTemplateFile(ctx, d)

	return
}
