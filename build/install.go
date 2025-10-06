package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/ag7if/go-files"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/internal/logging"
)

func CreateConfigDirectories() error {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}

	logging.Info().Str("dir", cfgDir.Path()).Msg("creating config directory")
	err = os.Mkdir(cfgDir.Path(), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cfgDir.Path()).Msg("failed to create config directory")
	}

	logging.Info().Str("subdir", "cfg").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir.Path(), "cfg"), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "cfg").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "assets").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir.Path(), "assets"), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "assets").Msg("failed to create subdirectory")
	}

	return nil
}

func CreateCacheDirectory() error {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return errors.WithStack(err)
	}

	logging.Info().Str("dir", cacheDir.Path()).Msg("creating cache directory")
	err = os.Mkdir(cacheDir.Path(), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cacheDir.Path()).Msg("failed to create directory")
	}

	return nil
}

func CopyAssets(projectPath string) error {
	defaultCfgDir := filepath.Join(projectPath, "config")
	imgDir := filepath.Join(projectPath, "assets", "logos", "personal_symbols", "personal_achievement")

	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}
	destAssetDir := filepath.Join(cfgDir.Path(), "assets")
	destCfgDir := filepath.Join(cfgDir.Path(), "cfg")

	logging.Info().Str("file", "full_achievement-color.png").Msg("copying asset")
	achievement, err := files.NewFile(filepath.Join(imgDir, "full_achievement-color.png"))
	if err != nil {
		logging.Warn().Err(err).Str("file", "full_achievement-color.png").Msg("failed to copy asset")
	}
	_, err = achievement.Copy(destAssetDir)
	if err != nil {
		logging.Warn().Err(err).Str("file", "full_achievement-color.png").Msg("failed to copy asset")
	}
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "cap_command_emblem.jpg").Msg("failed to copy asset")
	}

	logging.Info().Str("file", "default.yaml").Msg("copying config")
	defaultCfg, err := files.NewFile(filepath.Join(defaultCfgDir, "default.yaml"))
	if err != nil {
		logging.Warn().Err(err).Str("file", "duty_assignments.yaml").Msg("failed to copy config")
	}
	_, err = defaultCfg.Copy(destCfgDir)
	if err != nil {
		logging.Warn().Err(err).Str("file", "duty_assignments.yaml").Msg("failed to copy config")
	}
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "duty_assignments.yaml").Msg("failed to copy config")
	}

	return nil
}

func ClearFileExistsError(err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()
	if strings.Contains(msg, "file exists") {
		return nil
	}

	return err
}

func ClearFileDoesNotExistError(err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()
	if strings.Contains(msg, "no such file or directory") {
		return nil
	}

	return err
}

func main() {
	logging.InitLogging("info", true)

	err := CreateConfigDirectories()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create config directories")
		os.Exit(1)
	}

	err = CreateCacheDirectory()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create cache directories")
		os.Exit(1)
	}

	err = CopyAssets(os.Args[1])
	if err != nil {
		logging.Error().Err(err).Msg("failed to copy assets")
		os.Exit(1)
	}
}
