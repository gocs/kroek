package main

import (
	"bytes"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gocs/kroek/domain"

	"github.com/gocs/kroek/space"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

var (
	screen  *space.Screen
	theGame *space.Game
)

func init() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetFullscreen(true)
	w, h := ebiten.ScreenSizeInFullscreen()

	// On mobiles, ebiten.MonitorSize is not available so far.
	// Use arbitrary values.
	if w == 0 || h == 0 {
		w = 300
		h = 450
	}

	screen = space.NewScreen(w, h)

	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	var filename string
	if len(os.Args) > 0 {
		filename = os.Args[1]
	}

	reading, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	img2, _, err := image.Decode(bytes.NewReader(reading))
	if err != nil {
		log.Fatal(err)
	}
	mapImage, _ := ebiten.NewImageFromImage(img2, ebiten.FilterDefault)

	// Initialize the sprites.
	sprites := []*domain.Sprite{}
	w, h = ebitenImage.Size()
	for i := 0; i < 2; i++ {
		s := domain.NewSprite(
			screen,
			ebitenImage,
			rand.Intn(screen.Width()-w),
			rand.Intn(screen.Height()-h),
		)
		sprites = append(sprites, s)
	}
	middlePosX, middlePosY := mapImage.Size()
	mapSprite := domain.NewSprite(screen, mapImage, screen.Width()/2-middlePosX/2, screen.Height()/2-middlePosY/2)

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
