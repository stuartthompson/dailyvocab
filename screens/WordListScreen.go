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

	"github.com/stuartthompson/dailyvocab/app"
	"github.com/stuartthompson/dailyvocab/configuration"
	"github.com/stuartthompson/dailyvocab/io/screen"
)

// checkmarkColor ...
// Color used for the checkmark displayed next to viewed words.
const checkmarkColor = 3 // Green

// screenBorderColor ...
// Color used for the word list screen border.
const screenBorderColor = 5 // Blue

// WordListScreen ...
type WordListScreen struct {
	screen        *screen.Screen
	configuration *configuration.AppConfig // Application configuration
	viewedWords   map[int]string           // Map of viewed words by id
	vocabulary    *app.Vocabulary          // The word list to render
}

// NewWordListScreen ...
// Instantiates a new word list screen.
func NewWordListScreen(config *configuration.AppConfig, vocabulary *app.Vocabulary, viewport *screen.Viewport) *WordListScreen {
	screenStyle := &screen.Style{ShowBorder: true, BorderColor: screenBorderColor}
	screen := screen.NewScreen(viewport, screenStyle)

	// Create new word list screen
	wordListScreen := &WordListScreen{screen: screen, configuration: config, vocabulary: vocabulary}

	// Build viewed words map
	wordListScreen.viewedWords = buildViewedWordsMap(config.ViewedWords)

	// Return new word list screen
	return wordListScreen
}

// Render ...
// Renders the about screen.
func (s *WordListScreen) Render() {
	s.screen.Clear()

	s.screen.RenderText("Word List", 1, 0, 255, 0)

	// Render header
	totalWords := len(s.vocabulary.Words)
	viewedWords := len(s.configuration.ViewedWords)
	// Determine how many words will be displayed
	wordsPerPage := s.screen.GetHeight() - 5
	pageNum := 1
	startIndex := (wordsPerPage * (pageNum - 1)) + 1
	endIndex := startIndex + wordsPerPage
	if endIndex > totalWords {
		endIndex = totalWords
	}
	// Render header text
	headerText := fmt.Sprintf("Showing %d - %d of %d total words. Viewed %d.", startIndex, endIndex, totalWords, viewedWords)
	s.screen.RenderText(headerText, 1, 2, 255, 0)

	// Render word list
	for i := 0; i < len(s.vocabulary.Words); i++ {
		// Calculate y-coordinate at which to render this line
		y := 4 + i

		w := s.vocabulary.Words[i]
		// Get the word in the default language
		word := s.vocabulary.GetWordInLanguage(w.ID, s.configuration.DefaultLanguage)
		// Get total number of languages
		numLangs := len(s.vocabulary.Words[i].Translations)
		// Render "viewed" checkmark (if word is marked viewed)
		if s.viewedWords[w.ID] != "" {
			s.screen.RenderText("✓", 1, y, checkmarkColor, 0)
		}
		// Render main list item text
		str := fmt.Sprintf("[%d] %s (in %d languages)", w.ID, word, numLangs)
		s.screen.RenderText(str, 3, y, 255, 0)

	}
}

// buildViewedWordsMap ...
// Builds a map that is used to quickly look up words that have been viewed.
// This is an optimization to speed up checking if a word has been viewed during the render cycle.
func buildViewedWordsMap(viewedWords []configuration.ViewedWord) map[int]string {
	viewedWordsMap := make(map[int]string)
	// Build the map of viewed words
	for i := 0; i < len(viewedWords); i++ {
		viewedWordsMap[viewedWords[i].ID] = viewedWords[i].MarkedViewedAt
	}

	return viewedWordsMap
}
