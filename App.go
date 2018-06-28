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
	"log"

	termbox "github.com/nsf/termbox-go"
	"github.com/stuartthompson/dailyvocab/configuration"
	"github.com/stuartthompson/dailyvocab/entities"
	"github.com/stuartthompson/dailyvocab/io"
	"github.com/stuartthompson/dailyvocab/io/canvas"
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
	wordList        *entities.WordList
	currentScreen   Screen
	dailyWordScreen *screens.DailyWordScreen
	wordListScreen  *screens.WordListScreen
	configScreen    *screens.ConfigScreen
	aboutScreen     *screens.AboutScreen
	bottomBar       *screens.BottomBarComponent
}

// NewApp ...
// Initializes a new application instance.
func NewApp() *App {
	app := &App{
		isRunning:     true,
		configuration: &configuration.AppConfig{},
		wordList:      &entities.WordList{},
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

	// Read configuration
	err = a.configuration.ReadConfiguration()
	if err != nil {
		log.Print("Unable to read configuration. Exiting.")
		return
	}

	// Read word list
	err = a.wordList.Load()
	if err != nil {
		log.Print("Unable to read word list. Exiting.")
		return
	}

	// Initialize canvas
	width, height := io.GetWindowSize()

	// TODO: Use flex-box logic to size canvases
	bottomBarHeight := 3 // Height without borders
	mainCanvas := canvas.NewCanvas(0, 0, width, height-bottomBarHeight)
	bottomBarCanvas := canvas.NewCanvas(0, height-bottomBarHeight, width, bottomBarHeight)

	// Initialize screens
	a.wordListScreen = screens.NewWordListScreen(a.configuration, a.wordList, mainCanvas)
	a.configScreen = screens.NewConfigScreen(a.configuration, mainCanvas)
	a.dailyWordScreen = screens.NewDailyWordScreen(a.configuration, mainCanvas)
	a.aboutScreen = screens.NewAboutScreen(a.configuration, mainCanvas)
	a.bottomBar = screens.NewBottomBarComponent(a.configuration, bottomBarCanvas)

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
	// Render current screen
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

	// Render bottom bar
	a.bottomBar.Render()

	io.Flush()
}

// registerKeypressHandlers ...
// Registers the key press handlers.
func (a *App) registerKeypressHandlers() {
	// TODO: Screens should really register their own list of keys vs. having a single global list
	a.eventListener.RegisterKeypressHandler('?', a.showAboutScreen)
	a.eventListener.RegisterKeypressHandler('w', a.showDailyWordScreen)
	a.eventListener.RegisterKeypressHandler('l', a.showWordListScreen)
	a.eventListener.RegisterKeypressHandler('c', a.showConfigScreen)
	a.eventListener.RegisterKeypressHandler('q', a.onQuit)
}

func (a *App) showDailyWordScreen() {
	a.currentScreen = DailyWordScreen
}

func (a *App) showWordListScreen() {
	a.currentScreen = WordListScreen
}

func (a *App) showConfigScreen() {
	a.currentScreen = ConfigScreen
}

func (a *App) showAboutScreen() {
	a.currentScreen = AboutScreen
}

// onQuit ...
// Called when the application should quit.
func (a *App) onQuit() {
	a.isRunning = false
}
