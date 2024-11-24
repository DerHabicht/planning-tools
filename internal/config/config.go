package config

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Configuration keys
const (
	HomeLocationLat      = "home_location.lat"
	HomeLocationLong     = "home_location.long"
	HomeLocationTz       = "home_location.tz"
	CoverLogo            = "cover_logo"
	Version              = "version"
	VisualCrossingAPIKey = "visual_crossing.api_key"
)

// Directory functions

// ConfigDir returns the directory where configuration files and assets are stored.
func ConfigDir() (string, error) {
	hd, err := os.UserConfigDir()
	if err != nil {
		return "", errors.WithMessage(err, "failed to find user home directory")
	}

	return filepath.Join(hd, "planning-tools"), nil
}

// CacheDir returns the directory where LaTeX files are built.
func CacheDir() (string, error) {
	cd, err := os.UserCacheDir()
	if err != nil {
		return "", errors.WithMessage(err, "failed to find user cache directory")
	}

	return filepath.Join(cd, "planning-tools"), nil
}

// Config setup functions

// SetConfigType wraps Viper's SetConfigType function
func SetConfigType(t string) {
	viper.SetConfigType(t)
}

// SetConfigName wraps Viper's SetConfigName function
func SetConfigName(name string) {
	viper.SetConfigName(name)
}

// AddConfigPath wraps Viper's AddConfigPath function
func AddConfigPath(path string) {
	viper.AddConfigPath(path)
}

// SetDefault wraps Viper's SetDefault function.
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

// ReadInConfig wraps Viper's ReadInConfig function.
func ReadInConfig() error {
	return viper.ReadInConfig()
}

// WriteConfigAs wraps Viper's WriteConfigAs function.
func WriteConfigAs(path string) error {
	return viper.WriteConfigAs(path)
}

// Configuration access functions

func Set(key string, value interface{}) {
	viper.Set(key, value)
}

// GetFloat64 wraps Viper's GetFloat64 function.
func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

// GetInt wraps Viper's GetInt function.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetString wraps Viper's GetString function.
func GetString(key string) string {
	return viper.GetString(key)
}
