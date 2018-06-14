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

package entities

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// WordList ...
// Represents a list of words.
type WordList struct {
	Words []Word
}

const wordListFileName = "wordlist.json"

// Load ...
// Loads the word list.
func (w *WordList) Load() error {
	// Read word list
	rawContent, err := ioutil.ReadFile(wordListFileName)
	if err != nil {
		log.Print(err)
		return err
	}

	// Unmarshal configuration
	err = json.Unmarshal(rawContent, &w.Words)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// GetWord ...
// Gets a word in a specific language.
func (w *WordList) GetWord(id int, languageCode string) string {
	// Find the requested word
	var word Word
	// TODO: Store a hashmap for lookup by word id
	for i := 0; i < len(w.Words); i++ {
		if w.Words[i].ID == id {
			word = w.Words[i]
		}
		continue
	}

	// If no word found with that id then return nil
	if w == nil {
		// TODO: Log this and handle this better
		return ""
	}

	// Get the requested translation for this word
	var translated string
	for i := 0; i < len(word.Translations); i++ {
		if word.Translations[i].LanguageCode == languageCode {
			translated = word.Translations[i].Word
		}
	}

	return translated
}
