package screen

import "github.com/stuartthompson/dailyvocab/io"

// Screen ...
// Represents a screen that can be rendered.
type Screen struct {
	viewport *Viewport
	style    *Style // A style object describing the appearance
}

// NewScreen ...
// Creates a new screen.
func NewScreen(viewport *Viewport, style *Style) *Screen {
	return &Screen{viewport: viewport, style: style}
}

// Clear ...
// Clears the screen, ready for rendering.
func (s *Screen) Clear() {
	io.ClearArea(s.viewport.x, s.viewport.y, s.viewport.x+s.viewport.width, s.viewport.y+s.viewport.height, 0)
	if s.style.ShowBorder == true {
		io.RenderPaneBorder(s.viewport.x, s.viewport.y, s.viewport.width-1, s.viewport.height-1, s.style.BorderColor, 0)
	}
}

// RenderText ...
// Renders a string at relative coordinates within the canvas using the supplied colors.
func (s *Screen) RenderText(text string, x int, y int, fgColor int, bgColor int) {
	// TODO: Check that text fits within the pane
	// TODO: Clean up calculation of x and y position (too confusing)
	io.RenderText(text, s.viewport.x+x+1, s.viewport.y+y+1, fgColor, bgColor)
}

// GetHeight ...
// Gets the height of this screen's viewport.
func (s *Screen) GetHeight() int {
	return s.viewport.height
}

// GetWidth ...
// Gets the width of this screen's viewport.
func (s *Screen) GetWidth() int {
	return s.viewport.width
}

// MoveAndResize ...
// Moves the screen and resizes it.
func (s *Screen) MoveAndResize(viewport Viewport) {
	// Set the new canvas position and size
	s.viewport.x = viewport.x
	s.viewport.y = viewport.y
	s.viewport.width = viewport.width
	s.viewport.height = viewport.height
}
