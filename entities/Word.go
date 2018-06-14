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

// LocalizedWord ...
// Represents a word in a specific language.
type LocalizedWord struct {
	LanguageCode string `json:"language-code"`
	Word         string `json:"word"`
}

// Word ...
// Represents a word.
type Word struct {
	ID           int             `json:"id"`
	Translations []LocalizedWord `json:"translations"`
	Meaning      string          `json:"meaning"`
}
