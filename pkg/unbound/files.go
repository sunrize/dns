package unbound

import (
	"os"
	"path/filepath"
)

const includeConfFilename = "include.conf"
const includeToplevelConfFilename = "include_toplevel.conf"

func (c *configurator) createEmptyIncludeConf() error {
	filepath := filepath.Join(c.unboundEtcDir, includeConfFilename)
	file, err := os.OpenFile(filepath, os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	return file.Close()
}

func (c *configurator) createEmptyIncludeToplevelConf() error {
	filepath := filepath.Join(c.unboundEtcDir, includeToplevelConfFilename)
	file, err := os.OpenFile(filepath, os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	return file.Close()
}
