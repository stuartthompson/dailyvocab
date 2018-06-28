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

package screen

// Viewport ...
// Represents a viewport to which a screen can render.
type Viewport struct {
	x      int
	y      int
	width  int
	height int
}

// NewViewport ...
// Creates a new viewport.
func NewViewport(x int, y int, width int, height int) *Viewport {
	return &Viewport{x: x, y: y, width: width, height: height}
}
