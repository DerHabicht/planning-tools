package plancal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/derhabicht/planning-tools/internal/calendar"
	"github.com/derhabicht/planning-tools/internal/logging"
)

var dlCmd = &cobra.Command{
	Use:   "dl [week]",
	Short: "Create day labels",
	Long:  ``,
	Args:  cobra.RangeArgs(0, 1),
	Run:   runDLCmd,
}

func runDLCmd(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

	var year int
	var week int
	if len(args) > 0 {
		weekre := `(\d{4})W(\d{2})`
		re, err := regexp.Compile(weekre)
		if err != nil {
			panic(err)
		}

		res := re.FindStringSubmatch(strings.ToUpper(args[0]))

		year, err = strconv.Atoi(res[1])
		if err != nil {
			logger.Error().Err(err).Msg("malformed week designator")
		}

		week, err = strconv.Atoi(res[2])
		if err != nil {
			logger.Error().Err(err).Msg("malformed week designator")
		}
	} else {
		now := time.Now()
		year, week = now.ISOWeek()
	}

	outputFile, err := files.NewFile(fmt.Sprintf("DayLabels-%04dW%02d.pdf", year, week), logger.DefaultLogger())
	if err != nil {
		logger.Error().Err(err).Msg("failed to create output file")
	}

	err = calendar.BuildLabels(year, week, outputFile, logger)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate labels")
	}

}

func init() {
	rootCmd.AddCommand(dlCmd)
}
