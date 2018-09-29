// backgroundRGB 16, 16, 16
// boost_color 78,196,217
// default_color 209,209,209
// hp_color 236,97,64
// trail_color 253,204,14
package main

import (
	"fmt"
	"image/color"
	"time"

	"golang.org/x/image/font/basicfont"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

var (
	prevFrameTime time.Time
	currFrameTime time.Time
	dt            float64
	win           *pixelgl.Window
	startError    error
	astroCount    int
	slowment      float64
	playerDead    bool
)

func run() {
	// set window config
	// monitor to display, if set -> fullscreen mode
	monitor := pixelgl.PrimaryMonitor()
	cfg := pixelgl.WindowConfig{
		Title:   "Asteroids",
		Bounds:  pixel.R(0, 0, 1600, 900),
		Monitor: monitor,
		VSync:   true,
	}

	win, startError = pixelgl.NewWindow(cfg)
	if startError != nil {
		panic(startError)
	}

	slowment = 1

	// score text
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	scoreplace := pixel.V(100, win.Bounds().Size().Y-100)
	scoreTxt := text.New(scoreplace, basicAtlas)

	deadText := text.New(win.Bounds().Center().Add(pixel.V(50, 50)), basicAtlas)

	// init game field
	field := &Area{
		DrawTarget:  win,
		GameObjects: make(map[gameObject]gameObject),
	}

	field.Restart()

	// for calculating dt
	prevFrameTime = time.Now()

	for !win.Closed() {
		// dt counting things
		currFrameTime = time.Now()
		dt = currFrameTime.Sub(prevFrameTime).Seconds()
		prevFrameTime = currFrameTime

		win.Clear(color.RGBA{16, 16, 16, 255})

		// close game on 'esc'
		if win.JustPressed(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		if win.JustPressed(pixelgl.KeyEnter) && playerDead {
			field.Restart()
			playerDead = false
		}

		field.update(dt * slowment)
		field.draw()
		if !playerDead {
			scoreTxt.Clear()
			fmt.Fprintf(scoreTxt, "Score %d", field.Score)
			scoreTxt.Draw(win, pixel.IM.Scaled(scoreTxt.Orig, 2))
		} else {
			scoreTxt.Clear()
			deadText.Clear()
			fmt.Fprintf(deadText, "SHIP CRASHED\n\nScore: %d", field.Score)
			deadText.Draw(win, pixel.IM.Scaled(deadText.Dot, 5))
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
