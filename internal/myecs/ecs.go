package myecs

import "github.com/bytearena/ecs"

var (
	Manager = ecs.NewManager()

	// Components
	Object = Manager.NewComponent()
	Parent = Manager.NewComponent()

	Clickable = Manager.NewComponent()

	Drawable = Manager.NewComponent()

	// Tags
	IsObject    = ecs.BuildTag(Object)
	HasParent   = ecs.BuildTag(Object, Parent)
	IsClickable = ecs.BuildTag(Object, Clickable)
	IsDrawable  = ecs.BuildTag(Object, Drawable)
)