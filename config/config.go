package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Config struct {
	RegistryUrl string `toml:"registry_url"`
	Host        string `toml:"host"`
	Port        string `toml:"port"`
}

func ReadConfigFile(filename string) (*Config, error) {
	config := Config{}

	path, err := expand(filename)
	if err != nil {
		return &config, err
	}

	d, err := ioutil.ReadFile(path)
	if err != nil {
		return &config, err
	} else if err = toml.Unmarshal(d, &config); err != nil {
		return &config, err
	}

	return &config, nil
}

func expand(path string) (string, error) {
	// Ignore if path has no leading tilde.
	if path != "~" && !strings.HasPrefix(path, "~"+string(os.PathSeparator)) {
		return path, nil
	}

	// Fetch the current user to determine the home path.
	u, err := user.Current()
	if err != nil {
		return path, err
	} else if u.HomeDir == "" {
		return path, fmt.Errorf("home directory unset")
	}

	if path == "~" {
		return u.HomeDir, nil
	}
	return filepath.Join(u.HomeDir, strings.TrimPrefix(path, "~"+string(os.PathSeparator))), nil
}
