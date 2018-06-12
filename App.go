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
	termbox "github.com/nsf/termbox-go"
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
	AboutScreen
)

// App ...
// Encapsulates main application logic.
type App struct {
	isRunning       bool
	eventListener   *io.EventListener
	currentScreen   Screen
	dailyWordScreen *screens.DailyWordScreen
	wordListScreen  *screens.WordListScreen
	aboutScreen     *screens.AboutScreen
}

// NewApp ...
// Initializes a new application instance.
func NewApp() *App {
	app := &App{
		isRunning:       true,
		dailyWordScreen: &screens.DailyWordScreen{},
		wordListScreen:  &screens.WordListScreen{},
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
	case AboutScreen:
		a.aboutScreen.Render()
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
