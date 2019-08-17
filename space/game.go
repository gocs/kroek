package space

import (
	"errors"

	"github.com/gocs/kroek/domain"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var errRegularTermination = errors.New("regular termination")

// Game the game logic sequence struct
type Game struct {
	strokes map[*domain.Stroke]struct{}
	sprites []*domain.Sprite
}

// NewGame generates game struct
func NewGame(strokes map[*domain.Stroke]struct{}, sprites []*domain.Sprite) *Game {
	return &Game{
		strokes: strokes,
		sprites: sprites,
	}
}

func (g *Game) spriteAt(x, y int) *domain.Sprite {
	// As the sprites are ordered from back to front,
	// search the clicked/touched sprite in reverse order.
	for i := len(g.sprites) - 1; i >= 0; i-- {
		s := g.sprites[i]
		if s.In(x, y) {
			return s
		}
	}
	return nil
}

func (g *Game) updateStroke(stroke *domain.Stroke) {
	stroke.Update()
	if !stroke.IsReleased() {
		return
	}

	s := stroke.DraggingObject().(*domain.Sprite)
	if s == nil {
		return
	}

	s.MoveBy(stroke.PositionDiff())

	index := -1
	for i, ss := range g.sprites {
		if ss == s {
			index = i
			break
		}
	}

	// Move the dragged sprite to the front.
	g.sprites = append(g.sprites[:index], g.sprites[index+1:]...)
	g.sprites = append(g.sprites, s)

	stroke.SetDraggingObject(nil)
}

// Update used for ebiten.Run handler
func (g *Game) Update(screen *ebiten.Image) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s := domain.NewStroke(&domain.MouseStrokeSource{})
		s.SetDraggingObject(g.spriteAt(s.Position()))
		g.strokes[s] = struct{}{}
	}
	for _, id := range inpututil.JustPressedTouchIDs() {
		s := domain.NewStroke(&domain.TouchStrokeSource{ID: id})
		s.SetDraggingObject(g.spriteAt(s.Position()))
		g.strokes[s] = struct{}{}
	}

	for s := range g.strokes {
		g.updateStroke(s)
		if s.IsReleased() {
			delete(g.strokes, s)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return errRegularTermination
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	draggingSprites := map[*domain.Sprite]struct{}{}
	for s := range g.strokes {
		if sprite := s.DraggingObject().(*domain.Sprite); sprite != nil {
			draggingSprites[sprite] = struct{}{}
		}
	}

	for _, s := range g.sprites {
		if _, ok := draggingSprites[s]; ok {
			continue
		}
		s.Draw(screen, 0, 0, 1)
	}
	for s := range g.strokes {
		dx, dy := s.PositionDiff()
		if sprite := s.DraggingObject().(*domain.Sprite); sprite != nil {
			sprite.Draw(screen, dx, dy, 0.5)
		}
	}

	// dragginMap := map[*domain.Map]struct{}{}

	ebitenutil.DebugPrint(screen, "Drag & Drop the sprites!")

	return nil
}
