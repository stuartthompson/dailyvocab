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

// BottomBarComponent ...
type BottomBarComponent struct {
	screen *screen.Screen
	config *configuration.AppConfig
}

// NewBottomBarComponent ...
// Instantiates a new bottom bar component.
func NewBottomBarComponent(config *configuration.AppConfig) *BottomBarComponent {
	return &BottomBarComponent{config: config}
}

// Render ...
// Renders the bottom bar component.
func (c *BottomBarComponent) Render() {
	c.screen.Clear()

	c.screen.RenderText("Bottom Bar", 0, 0, 255, 0)
}
