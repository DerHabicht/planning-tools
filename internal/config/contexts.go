package config

import (
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

func GetContextDefs() (map[string]string, error) {
	cfgDir, err := ConfigDir()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	f, err := files.NewFile(filepath.Join(cfgDir.Path(), "cfg", "contexts.yaml"))

	ctxDefRaw, err := f.ReadFile()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var defs map[string]string
	err = yaml.Unmarshal(ctxDefRaw, &defs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return defs, nil
}
