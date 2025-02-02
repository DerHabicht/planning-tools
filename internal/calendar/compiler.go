package calendar

import (
	"path/filepath"

	"github.com/ag7if/go-latex"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/internal/logging"
)

func configureLaTeXCompiler(logger logging.Logger) (*latex.Compiler, error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find config directory")
	}

	cacheDir, err := config.CacheDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find cache directory")
	}

	assetDir := filepath.Join(cfgDir, "assets")

	compiler := latex.NewCompiler(assetDir, cacheDir, logger.DefaultLogger())

	return &compiler, nil
}
