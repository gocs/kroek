package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/gocs/kroek/domain"

	"github.com/gocs/kroek/space"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

var (
	screen  *space.Screen
	theGame *space.Game
	w, h    = 0, 0
)

func init() {
	rand.Seed(time.Now().UnixNano())
	// ebiten.SetFullscreen(true)
	// w, h = ebiten.ScreenSizeInFullscreen()

	// On mobiles, ebiten.MonitorSize is not available so far.
	// Use arbitrary values.
	if w == 0 || h == 0 {
		// w = 300
		w = 800
		h = 450
	}

	screen = space.NewScreen(w, h)

	img, _, err := image.Decode(bytes.NewReader(images.Tile_png))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	// Initialize the sprites.
	sprites := []*domain.Sprite{}
	middlePosX, middlePosY := ebitenImage.Size()
	mapSprite := domain.NewSprite(
		screen,
		ebitenImage,
		screen.Width()/2-middlePosX/2, screen.Height()/2-middlePosY/2)

	sprites = append(sprites, mapSprite)

	// Initialize the game.
	theGame = space.NewGame(
		map[*domain.Stroke]struct{}{},
		sprites,
	)
}

func main() {
	if err := ebiten.Run(theGame.Update, screen.Width(), screen.Height(), 2, "Drag & Drop (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
