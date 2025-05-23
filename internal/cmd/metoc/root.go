package metoc

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/derhabicht/planning-tools/internal/logging"
	"github.com/derhabicht/planning-tools/internal/planwx"
)

var logLevel string

var rootCmd = &cobra.Command{
	Use:   "metoc_report <PLAN YAML> [OUTPUT TEX]",
	Short: "Generate METOC reports for operational planning",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.Logger{}

		planFile, err := files.NewFile(args[0], logger.DefaultLogger())
		if err != nil {
			logger.Error().Err(err).Str("filename", args[0]).Msg("failed to create reference to plan file")
			os.Exit(1)
		}

		var outputFilePath string
		if len(args) == 2 {
			outputFilePath = args[1]
		} else {
			outputFilePath = filepath.Join(planFile.Dir(), fmt.Sprintf("%s.%s", planFile.Base(), "pdf"))
		}
		outputFile, err := files.NewFile(outputFilePath, logger.DefaultLogger())
		if err != nil {
			logger.Error().Err(err).Str("filename", outputFilePath).Msg("failed to create reference to output file")
			os.Exit(1)
		}

		err = planwx.Generate(planFile, outputFile, logger)
		if err != nil {
			logger.Error().Err(err).Msg("failed to generate planning weather report")
		}
	},
}

func Execute(version string) {
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "")

	logging.InitLogging(logLevel, true)
}
