package clips

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mevdschee/minesweeper.go/sprites"
)

// Clip is a set of frames
type Clip struct {
	name          string
	x, y          int
	width, height int
	frame         int
	frames        []*ebiten.Image
}

// GetName gets the name of the clip
func (c *Clip) GetName() string {
	return c.name
}

// NewSlice creates a new slice sprite based clip
func NewSlice(spriteMap *sprites.SpriteMap, name string, x, y, width, height int) *Clip {
	frame0 := ebiten.NewImage(width, height)
	sprite := spriteMap.Sliced[name]

	srcY := sprite.Y
	dstY := 0
	for h := 0; h < 3; h++ {
		srcHeight := sprite.Heights[h]
		dstHeight := sprite.Heights[h]
		if h == 1 {
			dstHeight = height - sprite.Heights[0] - sprite.Heights[2]
		}
		srcX := sprite.X
		dstX := 0
		for w := 0; w < 3; w++ {
			srcWidth := sprite.Widths[w]
			dstWidth := sprite.Widths[w]
			if w == 1 {
				dstWidth = width - sprite.Widths[0] - sprite.Widths[2]
			}

			img := spriteMap.Image.SubImage(image.Rect(srcX, srcY, srcX+srcWidth, srcY+srcHeight)).(*ebiten.Image)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(float64(dstWidth)/float64(srcWidth), float64(dstHeight)/float64(srcHeight))
			op.GeoM.Translate(float64(dstX), float64(dstY))
			frame0.DrawImage(img, op)

			srcX += srcWidth + sprite.Gap
			dstX += dstWidth
		}
		srcY += srcHeight + sprite.Gap
		dstY += dstHeight
	}

	return &Clip{
		name:   name,
		x:      x,
		y:      y,
		width:  width,
		height: height,
		frames: []*ebiten.Image{frame0},
	}
}

// Draw draws the layer
func (c *Clip) Draw(screen *ebiten.Image) {
	img := c.frames[c.frame]
	srcWidth, srcHeight := img.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(c.width)/float64(srcWidth), float64(c.height)/float64(srcHeight))
	op.GeoM.Translate(float64(c.x), float64(c.y))
	screen.DrawImage(img, op)
}

// Update updates the clip
func (c *Clip) Update() (err error) {
	//if playing do:
	//c.frame = (c.frame + 1) % len(c.frames)
	//if moving do:
	//c.x++
	//if resizing do:
	//c.height--
	return nil
}