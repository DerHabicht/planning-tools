package plancal

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/derhabicht/planning-tools/internal/calendar"
	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/internal/logging"
)

var logLevel string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: viper.GetString("version"),
	Use:     "plancal <fiscal_year> [output_file]",
	Short:   "Generate an AG7IF Planning Calendar for the given fiscal year",
	Long:    ``,
	Args:    cobra.RangeArgs(1, 2),
	Run:     runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

	year, err := strconv.Atoi(args[0])
	if err != nil {
		logger.Error().Err(err).Str("fy", args[0]).Msg("specified value is not a valid fiscal year")
		os.Exit(1)
	}

	var outputFilePath string
	if len(args) == 2 {
		outputFilePath = args[1]
	} else {
		outputFilePath = fmt.Sprintf("PlanningCalendar-FY%d.pdf", year)
	}
	outputFile, err := files.NewFile(outputFilePath)
	if err != nil {
		logger.Error().Err(err).Str("filename", outputFilePath).Msg("failed to create reference to output file")
		os.Exit(1)
	}

	err = calendar.BuildCalendar(year, outputFile)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate planning_calendar")
		os.Exit(1)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	config.Set(config.Version, version)
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
