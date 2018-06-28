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

	"github.com/stuartthompson/dailyvocab/entities"
)

// WordListScreen ...
type WordListScreen struct {
	screen        *screen.Screen
	configuration *configuration.AppConfig
	wordList      *entities.WordList
}

// NewWordListScreen ...
// Instantiates a new word list screen.
func NewWordListScreen(config *configuration.AppConfig, wordList *entities.WordList) *WordListScreen {
	return &WordListScreen{configuration: config, wordList: wordList}
}

// Render ...
// Renders the about screen.
func (s *WordListScreen) Render() {
	s.screen.Clear()

	s.screen.RenderText("Word List", 1, 1, 255, 0)
	s.screen.RenderText("The word list is shown here.", 1, 3, 255, 0)

	// Render word list
	for i := 0; i < len(s.wordList.Words); i++ {
		w := s.wordList.Words[i]
		// Get the word in the default language
		word := s.wordList.GetWordInLanguage(w.ID, s.configuration.DefaultLanguage)
		str := fmt.Sprintf("Word %d: %d %s", i, w.ID, word)
		s.screen.RenderText(str, 1, 5+i, 255, 0)
	}
}
