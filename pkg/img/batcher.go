package img

import (
	"encoding/json"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/pkg/errors"
	"image/color"
	"os"
	"path/filepath"
)

var (
	IM       = pixel.IM
	Flip     = pixel.IM.ScaledXY(pixel.ZV, pixel.V(-1., 1.))
	Flop     = pixel.IM.ScaledXY(pixel.ZV, pixel.V(1., -1.))
	FlipFlop = pixel.IM.ScaledXY(pixel.ZV, pixel.V(-1., -1.))
	Batchers = map[string]*Batcher{}
	batchers []*Batcher
)

func ClearBatches() {
	for _, batcher := range batchers {
		if batcher.AutoClear {
			batcher.Clear()
		}
	}
}

func DrawBatches(target pixel.Target) {
	for _, batcher := range batchers {
		if batcher.AutoDraw {
			batcher.Draw(target)
		}
	}
}

type Batcher struct {
	Index      int
	Sprites    map[string]*pixel.Sprite
	Animations map[string]*Animation
	batch      *pixel.Batch
	AutoDraw   bool
	AutoClear  bool
}

func AddBatcher(key string, sheet *SpriteSheet, autoDraw, autoClear bool) {
	if _, ok := Batchers[key]; ok {
		Batchers[key].SetSpriteSheet(sheet)
		Batchers[key].AutoDraw = autoDraw
		Batchers[key].AutoClear = autoClear
	} else {
		Batchers[key] = NewBatcher(sheet, autoDraw, autoClear)
		batchers = append(batchers, Batchers[key])
	}
}

func NewBatcher(sheet *SpriteSheet, autoDraw, autoClear bool) *Batcher {
	b := &Batcher{
		Index:     len(batchers),
		AutoDraw:  autoDraw,
		AutoClear: autoClear,
	}
	b.SetSpriteSheet(sheet)
	return b
}

func (b *Batcher) GetFrame(key string, index int) *pixel.Sprite {
	if a, ok := b.Animations[key]; ok {
		if len(a.S) > index {
			return a.S[index]
		}
	}
	return nil
}

func (b *Batcher) GetSprite(key string) *pixel.Sprite {
	if s, ok := b.Sprites[key]; ok {
		return s
	}
	return nil
}

func (b *Batcher) GetAnimation(key string) *Animation {
	if a, ok := b.Animations[key]; ok {
		return a
	}
	return nil
}

func (b *Batcher) SetSpriteSheet(sheet *SpriteSheet) {
	b.batch = pixel.NewBatch(&pixel.TrianglesData{}, sheet.Img)
	b.Sprites = make(map[string]*pixel.Sprite)
	b.Animations = make(map[string]*Animation)
	for k, r := range sheet.SpriteMap {
		b.Sprites[k] = pixel.NewSprite(sheet.Img, r)
	}
	for k, a := range sheet.AnimMap {
		b.Animations[k] = NewAnimation(sheet, a.Sprites, a.Loop, a.Hold, a.dur)
	}
}

func (b *Batcher) Clear() {
	b.batch.Clear()
}

func (b *Batcher) Batch() *pixel.Batch {
	return b.batch
}

func (b *Batcher) DrawSprite(key string, mat pixel.Matrix) {
	b.Sprites[key].Draw(b.batch, mat)
}

func (b *Batcher) DrawSpriteColor(key string, mat pixel.Matrix, mask color.Color) {
	b.Sprites[key].DrawColorMask(b.batch, mat, mask)
}

func (b *Batcher) Draw(target pixel.Target) {
	b.batch.Draw(target)
}

type SpriteSheet struct {
	Img       pixel.Picture
	Sprites   []pixel.Rect
	SpriteMap map[string]pixel.Rect
	AnimMap   map[string]AnimDef
}

type AnimDef struct {
	Loop    bool
	Hold    bool
	Sprites []pixel.Rect
	dur     float64
}

type spriteFile struct {
	ImgFile   string   `json:"img"`
	Sprites   []sprite `json:"sprites"`
	Width     float64  `json:"width"`
	Height    float64  `json:"height"`
	SingleRow bool     `json:"singleRow"`
}

type sprite struct {
	K string  `json:"key"`
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`

	Loop   bool    `json:"loop"`
	Hold   bool    `json:"hold"`
	Dur    float64 `json:"dur"`
	Anim   bool    `json:"anim"`
	Frames int     `json:"frames"`
}

func LoadSpriteSheet(path string) (*SpriteSheet, error) {
	errMsg := "load sprite sheet"
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var fileSheet spriteFile
	err = decoder.Decode(&fileSheet)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	img, err := LoadImage(fmt.Sprintf("%s/%s", filepath.Dir(path), fileSheet.ImgFile))
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	sheet := &SpriteSheet{
		Img:       img,
		Sprites:   make([]pixel.Rect, 0),
		SpriteMap: make(map[string]pixel.Rect, 0),
		AnimMap:   make(map[string]AnimDef, 0),
	}
	x := 0.0
	for _, r := range fileSheet.Sprites {
		var rect pixel.Rect
		w := fileSheet.Width
		h := fileSheet.Height
		if r.W > 0.0 {
			w = r.W
		}
		if fileSheet.SingleRow {
			rect = pixel.R(x, 0.0, x+w, h)
			x += w
		} else {
			if r.H > 0.0 {
				h = r.H
			}
			rect = pixel.R(r.X, r.Y, r.X+w, r.Y+h)
		}
		sheet.Sprites = append(sheet.Sprites, rect)
		if r.K != "" {
			if def, ok := sheet.AnimMap[r.K]; ok {
				def.Sprites = append(def.Sprites, rect)
				for i := 1; i < r.Frames; i++ {
					def.Sprites = append(def.Sprites, rect)
				}
				sheet.AnimMap[r.K] = def
			} else {
				if r.Dur != 0.0 || r.Anim {
					spr := []pixel.Rect{rect}
					for i := 1; i < r.Frames; i++ {
						spr = append(spr, rect)
					}
					sheet.AnimMap[r.K] = AnimDef{
						Loop:    r.Loop,
						Hold:    r.Hold,
						Sprites: spr,
						dur:     r.Dur,
					}
				}
				sheet.SpriteMap[r.K] = rect
			}
		}
	}
	return sheet, nil
}
