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

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"

	termbox "github.com/nsf/termbox-go"
	"github.com/stuartthompson/dailyvocab/configuration"
	"github.com/stuartthompson/dailyvocab/io"
	"github.com/stuartthompson/dailyvocab/screens"
)

// Screen ...
// Typedef for screen types.
type Screen int

// Defines screen types.
const (
	DailyWordScreen = iota
	WordListScreen
	ConfigScreen
	AboutScreen
)

// configFileName ...
// The name of the application configuration file.
const configFileName = ".dailyvocab"

// App ...
// Encapsulates main application logic.
type App struct {
	isRunning       bool
	eventListener   *io.EventListener
	configuration   *configuration.AppConfig
	currentScreen   Screen
	dailyWordScreen *screens.DailyWordScreen
	wordListScreen  *screens.WordListScreen
	configScreen    *screens.ConfigScreen
	aboutScreen     *screens.AboutScreen
}

// NewApp ...
// Initializes a new application instance.
func NewApp() *App {
	app := &App{
		isRunning:       true,
		dailyWordScreen: &screens.DailyWordScreen{},
		wordListScreen:  &screens.WordListScreen{},
		configScreen:    &screens.ConfigScreen{},
		aboutScreen:     &screens.AboutScreen{},
	}
	app.eventListener = io.NewEventListener(app.Render)

	return app
}

// Run ...
// Runs the application.
func (a *App) Run() {
	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)
	defer termbox.Close()

	// Build configuration file path
	configFilePath, err := a.buildConfigFilePath()
	if err != nil {
		log.Print("Unable to build configuration file path. Exiting.")
		return
	}
	// Load configuration
	a.configuration, err = a.loadConfiguration(configFilePath)
	if err != nil {
		log.Print("Unable to read configuration file. Exiting.")
		return
	}

	// Register keypress handlers
	a.registerKeypressHandlers()

	// Render screen (initially)
	a.Render()

	// Start main app loop
	for a.isRunning {
		a.eventListener.WaitForEvent()
		a.Render()
	}
}

// Render ...
// Renders the current screen.
func (a *App) Render() {
	switch a.currentScreen {
	case DailyWordScreen:
		a.dailyWordScreen.Render()
	case WordListScreen:
		a.wordListScreen.Render()
	case ConfigScreen:
		a.configScreen.Render()
	case AboutScreen:
		a.aboutScreen.Render()
	}
}

// buildConfigFilePath ...
// Builds the file path for the application configuration file.
func (a *App) buildConfigFilePath() (string, error) {
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

func (a *App) loadConfiguration(configFilePath string) (*configuration.AppConfig, error) {
	// Create config if it does not exist
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		a.writeDefaultConfiguration(configFilePath)
	}

	// Read configuration file
	rawConfig, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Unmarshal configuration
	var config *configuration.AppConfig
	json.Unmarshal(rawConfig, config)
	return config, nil
}

// writeDefaultConfiguration ...
// Writes a default configuration file.
func (a *App) writeDefaultConfiguration(configFilePath string) {
	log.Print("Writing default configuration")
	config := configuration.AppConfig{DefaultLanguage: "english"}
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

func (a *App) registerKeypressHandlers() {
	// TODO: Screens should really register their own list of keys vs. having a single global list
	a.eventListener.RegisterKeypressHandler('?', a.toggleAboutScreen)
	a.eventListener.RegisterKeypressHandler('q', a.onQuit)
}

func (a *App) toggleAboutScreen() {
	if a.currentScreen == AboutScreen {
		// TODO: Restore previous screen instead of always going to daily word screen
		a.currentScreen = DailyWordScreen
	} else {
		a.currentScreen = AboutScreen
	}
}

func (a *App) onQuit() {
	a.isRunning = false
}
