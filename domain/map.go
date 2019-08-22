package domain

import (
	"github.com/hajimehoshi/ebiten"
)

// Point ...
type Point struct {
	x, y int
}

// MapSprite represents an image.
type MapSprite struct {
	Cities []*City

	// MapSprite need these implementations
	Spriter
}

// Mapper struct using Map must have these methods
type Mapper interface {
	Padding(ceiling, left, floor, right int)
	Margin(ceiling, left, floor, right int)
}

// MapSpriter also implements spriter
type MapSpriter interface {
	Spriter
}

// NewMapSprite generates new map sprite
func NewMapSprite(screen Screener, spriteImage, spriteImageLarger *ebiten.Image, x, y int, cities []*City) MapSpriter {
	return &MapSprite{
		Cities:  cities,
		Spriter: NewSprite(screen, spriteImage, spriteImageLarger, x, y),
	}
}

// City basic city struct
type City struct {
	name     string
	location Point
}
