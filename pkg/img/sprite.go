package img

import (
	"github.com/faiface/pixel"
	"image/color"
)

type Handle struct {
	Sprite *pixel.Sprite
	Color  color.Color
	Batch  string
}