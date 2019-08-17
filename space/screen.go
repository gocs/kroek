package space

// Screen sets the screen
type Screen struct {
	width, height int
}

// NewScreen sets the screen width and height
func NewScreen(width, height int) *Screen {
	return &Screen{width: width, height: height}
}

// Width gets the once set readonly width
func (scr *Screen) Width() int {
	return scr.width
}

// Height gets the once set readonly height
func (scr *Screen) Height() int {
	return scr.height
}
