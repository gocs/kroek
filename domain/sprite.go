package domain

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Sprite represents an image.
type Sprite struct {
	screen Screener
	image  *ebiten.Image
	x      int
	y      int
	cities []struct {
		x int
		y int
	}
	Width  int
	Height int
}

// NewSprite sprite initializer
func NewSprite(screen Screener, spriteImage *ebiten.Image, x, y int) *Sprite {
	w, h := spriteImage.Size()
	return &Sprite{
		screen: screen,
		image:  spriteImage,
		x:      x,
		y:      y,
		Width:  w,
		Height: h,
	}
}

// Screener struct must have width and height; basically screen struct to avoid cyclic dep
type Screener interface {
	Width() int
	Height() int
}

// Spriter struct must have sprites methods
type Spriter interface {
	In(x, y int) bool
	MoveBy(x, y int)
	Draw(screen *ebiten.Image, dx, dy int, alpha float64)
}

// In returns true if (x, y) is in the sprite, and false otherwise.
func (s *Sprite) In(x, y int) bool {
	// Check the actual color (alpha) value at the specified position
	// so that the result of In becomes natural to users.
	//
	// Note that this is not a good manner to use At for logic
	// since color from At might include some errors on some machines.
	// As this is not so important logic, it's ok to use it so far.
	return s.image.At(x-s.x, y-s.y).(color.RGBA).A > 0
}

// MoveBy moves the sprite by (x, y).
func (s *Sprite) MoveBy(x, y int) {
	w, h := s.image.Size()

	s.x += x
	s.y += y

	// if x, y is <1/2 of w, h { set to 1/2 w, h}
	// w favorably gives buffer
	if s.x < s.screen.Width()/2-w {
		s.x = s.screen.Width()/2 - w
	}
	if s.x > s.screen.Width()/2 {
		s.x = s.screen.Width() / 2
	}
	if s.y < s.screen.Height()/2-h {
		s.y = s.screen.Height()/2 - h
	}
	if s.y > s.screen.Height()/2 {
		s.y = s.screen.Height() / 2
	}
}

// Draw draws the sprite.
func (s *Sprite) Draw(screen *ebiten.Image, dx, dy int, alpha float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x+dx), float64(s.y+dy))
	op.ColorM.Scale(1, 1, 1, alpha)

	screen.DrawImage(s.image, op)
}
