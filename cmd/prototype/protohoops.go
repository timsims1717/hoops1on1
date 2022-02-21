package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"timsims1717/hoops1on1/internal/data"
	"timsims1717/hoops1on1/internal/myecs"
	"timsims1717/hoops1on1/internal/systems"
	"timsims1717/hoops1on1/pkg/camera"
	"timsims1717/hoops1on1/pkg/img"
	"timsims1717/hoops1on1/pkg/object"
	"timsims1717/hoops1on1/pkg/timing"
)

func run() {
	conf := pixelgl.WindowConfig{
		Title:  "Hoops! 1 on 1 - prototype",
		Bounds: pixel.R(0, 0, 1600, 900),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(conf)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	camera.Cam = camera.New(true)
	camera.Cam.Opt.WindowScale = 900.
	camera.Cam.SetZoom(1.)
	camera.Cam.SetSize(1600., 900.)

	cardImg, err := img.LoadImage("assets/img/testcard.png")
	if err != nil {
		panic(err)
	}

	card := object.New()
	spr := pixel.NewSprite(cardImg, cardImg.Bounds())
	card.Rect = spr.Frame()
	e := myecs.Manager.NewEntity()
	e.AddComponent(myecs.Object, card).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Clickable, data.NewFrameFunc(func() bool {
			if e.HasComponent(myecs.Parent) {
				e.RemoveComponent(myecs.Parent)
			} else {
				e.AddComponent(myecs.Parent, &data.GameInput.World)
			}
			return false
		}))

	timing.Reset()
	for !win.Closed() {
		timing.Update()

		data.GameInput.Update(win)
		systems.ParentSystem()
		systems.TransformSystem()
		systems.ClickableSystem()
		camera.Cam.Update(win)

		win.Clear(color.RGBA{
			R: 100,
			G: 100,
			B: 100,
			A: 255,
		})

		systems.DrawSystem(win)

		win.Update()
	}
}

// fire the run function (the real main function)
func main() {
	pixelgl.Run(run)
}