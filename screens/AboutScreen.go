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

package screens

import (
	"github.com/stuartthompson/dailyvocab/configuration"
	"github.com/stuartthompson/dailyvocab/io/screen"
)

// AboutScreen ...
type AboutScreen struct {
	screen        *screen.Screen
	configuration *configuration.AppConfig
}

// NewAboutScreen ...
// Instantiates a new about screen.
func NewAboutScreen(config *configuration.AppConfig) *AboutScreen {
	return &AboutScreen{configuration: config}
}

// Render ...
// Renders the about screen.
func (s *AboutScreen) Render() {
	s.screen.Clear()

	s.screen.RenderText("About", 1, 1, 255, 0)
	s.screen.RenderText("DailyVocab presents a word of the day in different languages.", 1, 3, 255, 0)
}
