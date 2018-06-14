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

	"github.com/stuartthompson/dailyvocab/entities"
	"github.com/stuartthompson/dailyvocab/io"
)

// WordListScreen ...
type WordListScreen struct {
	Configuration *configuration.AppConfig
	WordList      *entities.WordList
}

// Render ...
// Renders the about screen.
func (s *WordListScreen) Render() {
	io.ClearScreen(0)
	width, height := io.GetWindowSize()
	io.RenderPaneBorder(0, 0, width-1, height-1, 72, 0)
	io.RenderText("Word List", 1, 1, 255, 0)
	io.RenderText("The word list is shown here.", 1, 3, 255, 0)

	// Render word list
	for i := 0; i < len(s.WordList.Words); i++ {
		w := s.WordList.Words[i]
		// Get the word in the default language
		word := s.WordList.GetWord(w.ID, s.Configuration.DefaultLanguage)
		s := fmt.Sprintf("Word %d: %d %s", i, w.ID, word)
		io.RenderText(s, 1, 5+i, 255, 0)
	}
	io.Flush()
}
