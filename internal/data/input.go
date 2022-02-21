package data

import (
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/hoops1on1/pkg/input"
)

var GameInput = &input.Input{
	Buttons: map[string]*input.ButtonSet{
		"menuUp": {
			Keys:    []pixelgl.Button{pixelgl.KeyW, pixelgl.KeyUp, pixelgl.KeyKP8},
			Buttons: []pixelgl.GamepadButton{pixelgl.ButtonDpadUp},
		},
		"menuDown": {
			Keys:    []pixelgl.Button{pixelgl.KeyS, pixelgl.KeyDown, pixelgl.KeyKP5},
			Buttons: []pixelgl.GamepadButton{pixelgl.ButtonDpadDown},
		},
		"menuLeft": {
			Keys:    []pixelgl.Button{pixelgl.KeyA, pixelgl.KeyLeft, pixelgl.KeyKP4},
			Buttons: []pixelgl.GamepadButton{pixelgl.ButtonDpadLeft},
		},
		"menuRight": {
			Keys:    []pixelgl.Button{pixelgl.KeyD, pixelgl.KeyRight, pixelgl.KeyKP6},
			Buttons: []pixelgl.GamepadButton{pixelgl.ButtonDpadRight},
		},
		"menuSelect": {
			Keys:    []pixelgl.Button{pixelgl.KeySpace, pixelgl.KeyEnter, pixelgl.KeyKPEnter},
			Buttons: []pixelgl.GamepadButton{pixelgl.ButtonA},
		},
		"menuBack":   input.New(pixelgl.KeyEscape, pixelgl.ButtonB),
		"inputClear": input.New(pixelgl.KeyF1, pixelgl.ButtonBack),
		"click":      input.NewJoyless(pixelgl.MouseButtonLeft),
		"scrollUp": {
			Scroll: 1,
		},
		"scrollDown": {
			Scroll: -1,
		},
		"pause": input.New(pixelgl.KeyEscape, pixelgl.ButtonStart),
	},
	StickD: true,
	Mode:   input.KeyboardMouse,
}