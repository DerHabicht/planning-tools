package plancal

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/derhabicht/planning-tools/internal/calendar"
	"github.com/derhabicht/planning-tools/internal/logging"
)

var kfmCmd = &cobra.Command{
	Use:   "kfm <fiscal_year>",
	Short: "Generate a blank YAML file for specifying KFM lessons for the given FY",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run:   runKFMCmd,
}

func runKFMCmd(_ *cobra.Command, args []string) {
	logger := logging.Logger{}

	year, err := strconv.Atoi(args[0])
	if err != nil {
		logger.Error().Err(err).Str("fy", args[0]).Msg("argument is not a valid fiscal year")
		os.Exit(1)
	}

	err = calendar.GenerateKFM(year)
}

func init() {
	rootCmd.AddCommand(kfmCmd)
}
