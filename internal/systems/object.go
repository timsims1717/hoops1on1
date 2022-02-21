package systems

import (
	"github.com/faiface/pixel"
	"math"
	"timsims1717/hoops1on1/internal/myecs"
	"timsims1717/hoops1on1/pkg/object"
)

func TransformSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsObject) {
		if trans, ok := result.Components[myecs.Object].(*object.Object); ok {
			trans.APos = trans.Pos.Add(trans.Offset)
			trans.APos.X = math.Round(trans.APos.X)
			trans.APos.Y = math.Round(trans.APos.Y)
			trans.Mat = pixel.IM
			if trans.Flip && trans.Flop {
				trans.Mat = trans.Mat.Scaled(pixel.ZV, -1.)
			} else if trans.Flip {
				trans.Mat = trans.Mat.ScaledXY(pixel.ZV, pixel.V(-1., 1.))
			} else if trans.Flop {
				trans.Mat = trans.Mat.ScaledXY(pixel.ZV, pixel.V(1., -1.))
			}
			trans.Mat = trans.Mat.ScaledXY(pixel.ZV, trans.Scalar)
			trans.Mat = trans.Mat.Rotated(pixel.ZV, math.Pi*trans.Rot)
			trans.Mat = trans.Mat.Moved(trans.APos)
		}
	}
}

func ParentSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasParent) {
		trans, okT := result.Components[myecs.Object].(*object.Object)
		parent := result.Components[myecs.Parent]
		if okT {
			if pos, ok := parent.(*pixel.Vec); ok {
				trans.Pos = *pos
			} else if pos, ok := parent.(pixel.Vec); ok {
				trans.Pos = pos // using a non-pointer to a pixel.Vec will freeze the item forever
			} else if par, ok := parent.(*object.Object); ok {
				trans.Pos = par.Pos
			}
		}
	}
}
