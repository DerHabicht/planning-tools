package calendar

import (
	"github.com/ag7if/go-latex"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/internal/config"
)

func configureLaTeXCompiler() (*latex.Compiler, error) {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	compiler := latex.NewCompiler(latex.XeLaTeX, latex.NoBib, *cacheDir)

	return &compiler, nil
}
