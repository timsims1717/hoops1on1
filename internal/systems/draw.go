package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/hoops1on1/internal/myecs"
	"timsims1717/hoops1on1/pkg/object"
)

func DrawSystem(win *pixelgl.Window) {
	for _, result := range myecs.Manager.Query(myecs.IsDrawable) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		if okO {
			draw := result.Components[myecs.Drawable]
			if spr, ok := draw.(*pixel.Sprite); ok {
				spr.Draw(win, obj.Mat)
			}
		}
	}
}