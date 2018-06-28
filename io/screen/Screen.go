package screen

import "github.com/stuartthompson/dailyvocab/io"

// Screen ...
// Represents a screen that can be rendered.
type Screen struct {
	x      int
	y      int
	width  int
	height int
	style  *Style // A style object describing the appearance
}

// NewScreen ...
// Creates a new screen.
func NewScreen(x int, y int, width int, height int, style *Style) *Screen {
	return &Screen{x: x, y: y, width: width, height: height, style: style}
}

// Clear ...
// Clears the screen, ready for rendering.
func (s *Screen) Clear() {
	io.ClearArea(s.x, s.y, s.x+s.width, s.y+s.height, 0)
	if s.style.showBorder == true {
		io.RenderPaneBorder(s.x, s.y, s.width-1, s.height-1, 0, s.style.borderColor)
	}
}

// RenderText ...
// Renders a string at relative coordinates within the canvas using the supplied colors.
func (s *Screen) RenderText(text string, x int, y int, fgColor int, bgColor int) {
	// TODO: Check that text fits within the pane
	// TODO: Clean up calculation of x and y position (too confusing)
	io.RenderText(text, s.x+x+1, s.y+y+1, fgColor, bgColor)
}

// MoveAndResize ...
// Moves the screen and resizes it.
func (s *Screen) MoveAndResize(x int, y int, width int, height int) {
	// Set the new canvas position and size
	s.x = x
	s.y = y
	s.width = width
	s.height = height
}
