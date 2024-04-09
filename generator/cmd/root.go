package cmd

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ARM-software/embedded-development-services-client/generator/codegen"
	configUtils "github.com/ARM-software/golang-utils/utils/config"
)

const (
	app          = "generate_code"
	inputArg     = "input"
	outputArg    = "output"
	templateArg  = "template"
	GenerateType = "generate"
)

var (
	extensionConfig codegen.ExtensionsConfig
	viperSession    = viper.New()
)

var rootCmd = &cobra.Command{
	Use: "extension-generator",
	RunE: func(_ *cobra.Command, _ []string) (err error) {
		ctx := context.Background()

		err = configUtils.LoadFromViper(viperSession, app, &extensionConfig, codegen.DefaultExtensionsConfig())
		if err != nil {
			fmt.Printf("%+v\n", extensionConfig)
			log.Errorf("Failed to initialise CLI with error: %s", err)
			return
		}
		d, err := codegen.GenerateDataStruct(extensionConfig)
		if err != nil {
			return
		}

		swagger, err := codegen.LoadAPIDefinition(d.SpecPath)
		if err != nil {
			return
		}

		d.Params, err = codegen.GetCollections(swagger)
		if err != nil {
			return
		}

		err = codegen.GenerateTemplateFile(ctx, d)
		return
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

	_ = configUtils.BindFlagToEnv(viperSession, app, "GENERATE_CODE_INPUT", rootCmd.Flags().Lookup(inputArg))
	_ = configUtils.BindFlagToEnv(viperSession, app, "GENERATE_CODE_OUTPUT", rootCmd.Flags().Lookup(outputArg))
	_ = configUtils.BindFlagToEnv(viperSession, app, "GENERATE_CODE_TEMPLATE", rootCmd.Flags().Lookup(templateArg))
}
