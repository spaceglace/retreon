package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"go.uber.org/zap"
)

var settings Settings

func GetSettings() Settings {
	return settings
}

type Settings struct {
	Profile  string              `json:"profile"`
	Profiles map[string]Profile  `json:"profiles"`
	Layout   string              `json:"layout"`
	Layouts  map[string]string   `json:"layouts"`
	Orders   map[string][]string `json:"orders"`
	Mode     string              `json:"mode"`
	Refresh  int                 `json:"refresh"`
}

func getConfigFilename() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(cwd, ".retro_config"), nil
}

func Initialize(l *zap.Logger) error {
	name, err := getConfigFilename()
	if err != nil {
		return err
	}
	// check for existance of config file
	_, err = os.Stat(name)
	if os.IsNotExist(err) {
		l.Warn("No config file found, using defaults")
		settings = Settings{
			Profile:  "",
			Profiles: map[string]Profile{},
			Layout:   "",
			Layouts:  map[string]string{},
			Mode:     "DateEarned",
			Orders:   map[string][]string{},
		}
		return nil
	}
	// if we got past the existance, assume we can read it
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		l.Error("Error opening config file",
			zap.Error(err),
		)
		return err
	}
	err = json.Unmarshal(contents, &settings)
	if err != nil {
		return err
	}

	return nil
}

func save() error {
	config, err := getConfigFilename()
	if err != nil {
		return err
	}
	output, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(config, output, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
