package domain

import (
	"github.com/hajimehoshi/ebiten"
)

// Point ...
type Point struct {
	x, y int
}

// Map this is the greater view of the known world
type Map struct {
	sprite Sprite
	cities []*Point
}

// Mapper struct using Map must have these methods
type Mapper interface {
	Padding(ceiling, left, floor, right int)
	Margin(ceiling, left, floor, right int)
}

// NewMap generates a new map
func NewMap(screen Screener, mapImage *ebiten.Image, x, y int, cities []*Point) *Map {
	return &Map{
		sprite: Sprite{
			screen: screen,
			image:  mapImage,
			x:      x,
			y:      y,
		},
		cities: cities,
	}
}

// In returns true if (x, y) is in the map, and false otherwise.
func (m *Map) In(x, y int) bool {
	return m.sprite.In(x, y)
}

// MoveBy moves the map by (x, y).
func (m *Map) MoveBy(x, y int) {
	m.sprite.MoveBy(x, y)
}

// Draw draws the map.
func (m *Map) Draw(screen *ebiten.Image, dx, dy int, alpha float64) {
	m.sprite.Draw(screen, dx, dy, alpha)
}
