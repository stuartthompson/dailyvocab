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
	"github.com/stuartthompson/dailyvocab/io"
)

// ConfigScreen ...
type ConfigScreen struct {
	Config *configuration.AppConfig
}

// Render ...
// Renders the config screen.
func (s *ConfigScreen) Render() {
	io.ClearScreen(0)
	width, height := io.GetWindowSize()
	io.RenderPaneBorder(0, 0, width-1, height-1, 200, 0)
	io.RenderText("Config", 1, 1, 255, 0)
	io.RenderText("Default language: "+s.Config.DefaultLanguage, 1, 3, 255, 0)
	io.Flush()
}
