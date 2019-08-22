package domain

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Sprite represents an image.
type Sprite struct {
	screen      Screener
	image       *ebiten.Image
	imageLarger *ebiten.Image
	isLarger    bool
	op          *ebiten.DrawImageOptions
	x           int
	y           int

	err error
}

// NewSprite sprite initializer
func NewSprite(screen Screener, spriteImage, spriteImageLarger *ebiten.Image, x, y int) Spriter {
	return &Sprite{
		screen:      screen,
		image:       spriteImage,
		imageLarger: spriteImageLarger,
		x:           x,
		y:           y,
	}
}

// Screener struct must have width and height; basically screen struct to avoid cyclic dep
type Screener interface {
	Width() int
	Height() int
}

// Spriter struct must have sprites methods to be called a sprite
type Spriter interface {
	In(x, y int) bool
	MoveBy(x, y int)
	DrawingBuilder
}

// In returns true if (x, y) is in the sprite, and false otherwise.
func (s *Sprite) In(x, y int) bool {
	// Check the actual color (alpha) value at the specified position
	// so that the result of In becomes natural to users.
	//
	// Note that this is not a good manner to use At for logic
	// since color from At might include some errors on some machines.
	// As this is not so important logic, it's ok to use it so far.

	// todo: recognize color even if resized
	// change s.image size by replacing image
	if s.isLarger {
		return s.imageLarger.At(x-s.x, y-s.y).(color.RGBA).A > 0
	}
	return s.image.At(x-s.x, y-s.y).(color.RGBA).A > 0
}

// MoveBy moves the sprite by (x, y).
func (s *Sprite) MoveBy(x, y int) {
	var w, h int
	if s.isLarger {
		w, h = s.imageLarger.Size()
	} else {
		w, h = s.image.Size()
	}

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

// DrawingBuilder Drawing options region
// considering struct of functions for better formatting
// anything between InitDrawingOptions and Draw are unordered
type DrawingBuilder interface {
	InitDrawingOptions() DrawingBuilder
	Move(dx, dy int) DrawingBuilder
	SwapLarger(isLarger bool) DrawingBuilder
	Draw(screen *ebiten.Image, alpha float64) error
}

// InitDrawingOptions every method in drawing must be preceded to keep the geom and colorm
func (s *Sprite) InitDrawingOptions() DrawingBuilder {
	s.op = &ebiten.DrawImageOptions{}
	return s
}

// Move Moves the image to a location; basically Translate
func (s *Sprite) Move(dx, dy int) DrawingBuilder {
	if s.op == nil {
		s.err = errors.New("add a &ebiten.DrawImageOptions{} to s.op")
		return s
	}
	s.op.GeoM.Translate(float64(s.x+dx), float64(s.y+dy))
	return s
}

// SwapLarger change image to its bigger alternative if true
func (s *Sprite) SwapLarger(isLarger bool) DrawingBuilder {
	s.isLarger = isLarger
	return s
}

// Draw draws the sprite.
func (s *Sprite) Draw(screen *ebiten.Image, alpha float64) error {
	if s.op == nil {
		return errors.New("add a &ebiten.DrawImageOptions{} to s.op")
	}
	s.op.ColorM.Scale(1, 1, 1, alpha)
	if s.isLarger {
		s.err = screen.DrawImage(s.imageLarger, s.op)
		return s.err
	}
	s.err = screen.DrawImage(s.image, s.op)
	return s.err
}
