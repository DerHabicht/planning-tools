package plancal

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/derhabicht/planning-tools/internal/calendar"
	"github.com/derhabicht/planning-tools/internal/logging"
)

var tabsCmd = &cobra.Command{
	Use:   "tabs [fiscal_year] [start_tab]",
	Short: "Create Avery 16282 tabs for the planning calendar",
	Long:  ``,
	Args:  cobra.RangeArgs(2, 3),
	Run:   runTabsCmd,
}

func runTabsCmd(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

	year, err := strconv.Atoi(args[0])
	if err != nil {
		logger.Error().Err(err).Str("fy", args[0]).Msg("first argument is not a valid fiscal year")
		os.Exit(1)
	}

	startTab, err := strconv.Atoi(args[1])
	if err != nil {
		logger.Error().Err(err).Str("fy", args[0]).Msg("second value is not a valid start tab")
		os.Exit(1)
	}

	if startTab < 1 || startTab > 20 {
		logger.Error().Err(err).Str("fy", args[0]).Msg("second value is not a valid start tab")
		os.Exit(1)
	}

	var outputFile files.File
	if len(args) == 3 {
		outputFile, err = files.NewFile(args[2])
	} else {
		outputFile, err = files.NewFile(fmt.Sprintf("PlanningCalendarTabs-FY%d.pdf", year))
	}
	if err != nil {
		logger.Error().Err(err).Msg("failed to create output file")
		os.Exit(1)
	}

	err = calendar.BuildCalendarTabs(year, startTab, outputFile)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate calendar tabs")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(tabsCmd)
}
