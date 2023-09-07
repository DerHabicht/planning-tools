package main

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/config"
	"github.com/derhabicht/planning-calendar/files"
	"github.com/derhabicht/planning-calendar/logging"
)

func CreateConfigDirectories() error {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}

	logging.Info().Str("dir", cfgDir).Msg("creating config directory")
	err = os.Mkdir(cfgDir, 0700)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cfgDir).Msg("failed to create config directory")
	}

	logging.Info().Str("subdir", "cfg").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "cfg"), 0700)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "cfg").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "assets").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "assets"), 0700)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "assets").Msg("failed to create subdirectory")
	}

	return nil
}

func CopyAssets(projectPath string) error {
	defaultCfgDir := filepath.Join(projectPath, "config")
	imgDir := filepath.Join(projectPath, "assets", "logos", "personal_symbols", "personal_achievement")
	texDir := filepath.Join(projectPath, "assets", "tex")

	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}
	destAssetDir := filepath.Join(cfgDir, "assets")
	destCfgDir := filepath.Join(cfgDir, "cfg")

	logging.Info().Str("file", "full_achievement-color.png").Msg("copying asset")
	err = files.Copy(filepath.Join(imgDir, "full_achievement-color.png"), filepath.Join(destAssetDir, "full_achievement-color.png"))
	if err != nil {
		logging.Warn().Err(err).Str("file", "cap_command_emblem.jpg").Msg("failed to copy asset")
	}

	logging.Info().Str("file", "calendar.tex").Msg("copying asset")
	err = files.Copy(filepath.Join(texDir, "calendar.tex"), filepath.Join(destAssetDir, "calendar.tex"))
	if err != nil {
		logging.Warn().Err(err).Str("file", "calendar.tex").Msg("failed to copy asset")
	}

	logging.Info().Str("file", "month.tex").Msg("copying asset")
	err = files.Copy(filepath.Join(texDir, "month.tex"), filepath.Join(destAssetDir, "month.tex"))
	if err != nil {
		logging.Warn().Err(err).Str("file", "month.tex").Msg("failed to copy asset")
	}

	logging.Info().Str("file", "default.yaml").Msg("copying config")
	err = files.Copy(filepath.Join(defaultCfgDir, "default.yaml"), filepath.Join(destCfgDir, "config.yaml"))
	if err != nil {
		logging.Warn().Err(err).Str("file", "duty_assignments.yaml").Msg("failed to copy config")
	}

	return nil
}

func main() {
	logging.InitLogging("info", true)

	err := CreateConfigDirectories()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create config directories")
		os.Exit(1)
	}

	err = CopyAssets(os.Args[1])
	if err != nil {
		logging.Error().Err(err).Msg("failed to copy assets")
		os.Exit(1)
	}

	config.InitConfig()
}
