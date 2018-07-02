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

package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const wordListFileName = "wordlist.json"

// LocalizedWord ...
// Represents a word in a specific language.
type LocalizedWord struct {
	LanguageCode string `json:"languageCode"`
	Native       string `json:"native"`
	Anglicized   string `json:"anglicized"`
}

// Word ...
// Represents a word.
type Word struct {
	ID           int             `json:"id"`
	Translations []LocalizedWord `json:"translations"`
	Usage        []WordUsage     `json:"usage"`
}

// Vocabulary ...
// Represents a list of words.
type Vocabulary struct {
	Words []Word
}

// WordUsage ...
// Describes how a word is used.
type WordUsage struct {
	Type    string `json:"type"`
	Meaning string `json:"meaning"`
}

// Load ...
// Loads the word list.
func (v *Vocabulary) Load() error {
	// Read word list
	rawContent, err := ioutil.ReadFile(wordListFileName)
	if err != nil {
		log.Print(err)
		return err
	}

	// Unmarshal configuration
	err = json.Unmarshal(rawContent, &v.Words)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// GetWord ...
// Gets a word by id.
func (v *Vocabulary) GetWord(id int) *Word {
	// Find the requested word
	var word *Word
	// TODO: Store a hashmap for lookup by word id
	for i := 0; i < len(v.Words); i++ {
		if v.Words[i].ID == id {
			word = &v.Words[i]
		}
		continue
	}

	// If no word found with that id then return nil
	if word == nil {
		// TODO: Log this and handle this better
		return nil
	}

	return word
}

// GetWordInLanguage ...
// Gets a word in a specific language.
func (v *Vocabulary) GetWordInLanguage(id int, languageCode string) string {
	// Get the requested word
	word := v.GetWord(id)
	if word == nil {
		// TODO: Log this
		return ""
	}

	// Get the requested translation for this word
	var translated string
	for i := 0; i < len(word.Translations); i++ {
		if word.Translations[i].LanguageCode == languageCode {
			translated = word.Translations[i].Native
		}
	}

	return translated
}
