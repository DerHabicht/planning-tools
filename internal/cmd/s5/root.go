package s5

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/internal/logging"
)

var logLevel string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: viper.GetString("version"),
	Use:     "s5",
	Short:   "",
	Long:    ``,
	//Args:    cobra.RangeArgs(1, 2),
	Run: runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

	logger.Fatal().Msg("s5 command not implemented")
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
