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
	"fmt"

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
func NewAboutScreen(config *configuration.AppConfig, viewport *screen.Viewport) *AboutScreen {
	screenStyle := &screen.Style{ShowBorder: true, BorderColor: 218}
	screen := screen.NewScreen(viewport, screenStyle)
	return &AboutScreen{screen: screen, configuration: config}
}

// Render ...
// Renders the about screen.
func (s *AboutScreen) Render() {
	s.screen.Clear()

	s.screen.RenderText("About", 1, 1, 255, 0)
	s.screen.RenderText("DailyVocab presents a word of the day in different languages.", 1, 3, 255, 0)

	// Render color grid
	s.screen.RenderText("Colors", 1, 5, 255, 0)

	rowIndex := 0
	columnIndex := 0
	xOffset := 1
	yOffset := 6
	for colorCode := 0; colorCode < 256; colorCode++ {
		str := fmt.Sprintf("%d", colorCode)
		s.screen.RenderText(str, columnIndex+xOffset, rowIndex+yOffset, colorCode, 0)
		rowIndex++
		if rowIndex+yOffset > s.screen.GetHeight()-4 {
			columnIndex += 4
			rowIndex = 0
		}
	}

	// Render background color grid
	rowIndex = 0
	columnIndex += 6
	for colorCode := 0; colorCode < 256; colorCode++ {
		s.screen.RenderText(fmt.Sprintf("%d", colorCode), columnIndex+xOffset, rowIndex+yOffset, 255, colorCode)
		rowIndex++
		if rowIndex+yOffset > s.screen.GetHeight()-4 {
			columnIndex += 4
			rowIndex = 0
		}
	}
}
