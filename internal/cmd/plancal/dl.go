package plancal

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/derhabicht/planning-tools/internal/calendar"
	"github.com/derhabicht/planning-tools/internal/logging"
)

var dlCmd = &cobra.Command{
	Use:   "dl [sprint|week]",
	Short: "Create day labels",
	Long:  ``,
	Args:  cobra.RangeArgs(0, 1),
	Run:   runDLCmd,
}

func runDLCmd(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

	cards, err := cmd.Flags().GetBool("cards")
	if err != nil {
		logger.Debug().Err(err).Msg("failed to parse --cards flag")
	}

	labels, err := cmd.Flags().GetBool("labels")
	if err != nil {
		logger.Debug().Err(err).Msg("failed to parse --labels flag")
	}

	contexts, err := cmd.Flags().GetStringSlice("contexts")
	if err != nil {
		logger.Debug().Err(err).Msg("failed to parse --contexts flag")
	}

	if !(cards || labels) {
		cards = true
		labels = true
	}

	sprint := false
	var year int
	var period int
	if len(args) > 0 {
		weekre := `(\d{4})(S|W)(\d{2})`
		re, err := regexp.Compile(weekre)
		if err != nil {
			panic(err)
		}

		res := re.FindStringSubmatch(strings.ToUpper(args[0]))

		year, err = strconv.Atoi(res[1])
		if err != nil {
			logger.Error().Err(err).Msg("malformed sprint/week designator")
		}

		if res[2] == "S" {
			sprint = true
		}

		period, err = strconv.Atoi(res[3])
		if err != nil {
			logger.Error().Err(err).Msg("malformed sprint/week designator")
		}
	} else {
		now := time.Now()
		year, period = now.ISOWeek()
	}

	err = calendar.BuildDL(year, period, sprint, cards, labels, contexts)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate labels")
		os.Exit(1)
	}
}

func init() {
	dlCmd.Flags().BoolP("cards", "c", false, "Generate cards")
	dlCmd.Flags().BoolP("labels", "l", false, "Generate labels")
	dlCmd.Flags().StringSliceP("contexts", "x", nil, "Contexts for generated cards")
	rootCmd.AddCommand(dlCmd)
}
