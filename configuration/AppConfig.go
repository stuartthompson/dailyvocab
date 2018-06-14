// Copyright 2018 Stuart Thompson.

// This file is part of DailyVocab.

// DailyVocab is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// DailyVocab is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with DailyVocab. If not, see <http://www.gnu.org/licenses/>.

package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

// configFileName ...
// The name of the application configuration file.
const configFileName = ".dailyvocab"

// AppConfig ...
// Represents configuration for the application.
type AppConfig struct {
	DefaultLanguage string `json:"DefaultLanguage"`
}

// ReadConfiguration ...
// Reads the application configuration from disk.
func (a *AppConfig) ReadConfiguration() error {
	configFilePath, err := a.buildConfigFilePath()
	if err != nil {
		log.Print("Error building config file path.")
		return err
	}

	// Create config if it does not exist
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		a.writeDefaultConfiguration(configFilePath)
	}

	// Read configuration file
	rawConfig, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Unmarshal configuration
	var config *AppConfig
	json.Unmarshal(rawConfig, &config)

	a.DefaultLanguage = config.DefaultLanguage
	return nil
}

// buildConfigFilePath ...
// Builds the file path for the application configuration file.
func (a *AppConfig) buildConfigFilePath() (string, error) {
	// Get user's home directory
	usr, err := user.Current()
	if err != nil {
		log.Print("Error getting current user. Error: ", err)
		return "", err
	}

	// Build configuration file path
	configFilePath := path.Join(usr.HomeDir, configFileName)

	return configFilePath, nil
}

// writeDefaultConfiguration ...
// Writes a default configuration file.
func (a *AppConfig) writeDefaultConfiguration(configFilePath string) {
	log.Print("Writing default configuration")
	config := AppConfig{DefaultLanguage: "en-us"}
	configJSON, err := json.Marshal(config)
	if err != nil {
		log.Print("Error marshaling json.")
		log.Fatal(err)
	}
	err = ioutil.WriteFile(configFilePath, configJSON, 0666)
	if err != nil {
		log.Print("Error writing configuration file.")
		log.Fatal("Error is: ", err)
	}
}
