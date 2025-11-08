package config

import (
	"os"
	"path/filepath"

	"github.com/derhabicht/planning-tools/internal/logging"
)

func init() {
	SetDefault(HomeLocationLat, 0.0)
	SetDefault(HomeLocationLong, 0.0)
	SetDefault(HomeLocationTz, "UTC")
	SetDefault(CoverLogo, "")
	SetDefault(VisualCrossingAPIKey, "")
	SetDefault(Birthday, "0001-01-01")

	cwd, _ := os.Getwd()
	logging.Trace().Str("cwd", cwd).Msg("")

	cfgDir, err := ConfigDir()
	if err != nil {
		logging.Error().Err(err).Msg("could not find user config dir")
	} else {
		logging.Debug().Str("usrCfgDir", cfgDir.Path()).Msg("user config dir found")
	}

	SetConfigType("yaml")
	SetConfigName("config")
	AddConfigPath(filepath.Join(cfgDir.Path(), "cfg"))
	err = ReadInConfig()
	if err != nil {
		path := filepath.Join(cfgDir.Path(), "cfg", "config.yaml")
		logging.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = WriteConfigAs(path)
		if err != nil {
			logging.Error().Err(err).Msg("failed to create default config file")
		}
	}
}
