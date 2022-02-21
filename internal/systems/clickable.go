package systems

import (
	"timsims1717/hoops1on1/internal/data"
	"timsims1717/hoops1on1/internal/myecs"
	"timsims1717/hoops1on1/pkg/object"
	"timsims1717/hoops1on1/pkg/util"
)

func ClickableSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsClickable) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		if okO {
			if util.PointInside(data.GameInput.World, obj.Rect, obj.Mat) &&
				data.GameInput.Get("click").JustPressed() {
				click := result.Components[myecs.Clickable]
				if fn, ok := click.(*data.FrameFunc); ok {
					if fn.Func() {
						myecs.Manager.DisposeEntity(result)
					}
				}
			}
		}
	}
}
