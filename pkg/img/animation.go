package img

import (
	"github.com/faiface/pixel"
)

type Animation struct {
	Loop bool
	Hold bool
	S    []*pixel.Sprite
	dur  float64
}

func NewAnimation(spriteSheet *SpriteSheet, a []pixel.Rect, loop, hold bool, dur float64) *Animation {
	var spr []*pixel.Sprite
	for _, r := range a {
		spr = append(spr, pixel.NewSprite(spriteSheet.Img, r))
	}
	return &Animation{
		Loop: loop,
		Hold: hold,
		S:    spr,
		dur:  dur,
	}
}