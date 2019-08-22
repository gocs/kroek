package main

import (
	"bytes"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/gocs/kroek/domain"

	"github.com/gocs/kroek/space"
	"github.com/hajimehoshi/ebiten"
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

	content, err := ioutil.ReadFile("assets/pirate-treasure-map-sticker.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	content1, err := ioutil.ReadFile("assets/pirate-treasure-map-sticker-bigger.png")
	if err != nil {
		log.Fatal(err)
	}
	img1, _, err := image.Decode(bytes.NewReader(content1))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage1, _ := ebiten.NewImageFromImage(img1, ebiten.FilterDefault)

	// Initialize the sprites.
	sprites := []domain.Spriter{}
	middlePosX, middlePosY := ebitenImage.Size()
	var cities []*domain.City
	mapSprite := domain.NewMapSprite(
		screen,
		ebitenImage,
		ebitenImage1,
		screen.Width()/2-middlePosX/2, screen.Height()/2-middlePosY/2,
		cities)

	sprites = append(sprites, mapSprite)

	// Initialize the game.
	theGame = space.NewGame(
		map[*domain.Stroke]struct{}{},
		sprites,
	)
}

func main() {
	if err := ebiten.Run(theGame.Update, screen.Width(), screen.Height(), 2, "Kroek"); err != nil {
		log.Fatal(err)
	}
}
