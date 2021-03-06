package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var skin = `
	iVBORw0KGgoAAAANSUhEUgAAAJAAAAB6CAMAAABnRypuAAADAFBMVEUAAACA
	AAAAgACAgAAAAICAAIAAgIDAwMCAgID/AAAA/wD//wAAAP//AP8A//////8Q
	EBARERESEhITExMUFBQVFRUWFhYXFxcYGBgZGRkaGhobGxscHBwdHR0eHh4f
	Hx8gICAhISEiIiIjIyMkJCQlJSUmJiYnJycoKCgpKSkqKiorKyssLCwtLS0u
	Li4vLy8wMDAxMTEyMjIzMzM0NDQ1NTU2NjY3Nzc4ODg5OTk6Ojo7Ozs8PDw9
	PT0+Pj4/Pz9AQEBBQUFCQkJDQ0NERERFRUVGRkZHR0dISEhJSUlKSkpLS0tM
	TExNTU1OTk5PT09QUFBRUVFSUlJTU1NUVFRVVVVWVlZXV1dYWFhZWVlaWlpb
	W1tcXFxdXV1eXl5fX19gYGBhYWFiYmJjY2NkZGRlZWVmZmZnZ2doaGhpaWlq
	ampra2tsbGxtbW1ubm5vb29wcHBxcXFycnJzc3N0dHR1dXV2dnZ3d3d4eHh5
	eXl6enp7e3t8fHx9fX1+fn5/f3+AgICBgYGCgoKDg4OEhISFhYWGhoaHh4eI
	iIiJiYmKioqLi4uMjIyNjY2Ojo6Pj4+QkJCRkZGSkpKTk5OUlJSVlZWWlpaX
	l5eYmJiZmZmampqbm5ucnJydnZ2enp6fn5+goKChoaGioqKjo6OkpKSlpaWm
	pqanp6eoqKipqamqqqqrq6usrKytra2urq6vr6+wsLCxsbGysrKzs7O0tLS1
	tbW2tra3t7e4uLi5ubm6urq7u7u8vLy9vb2+vr6/v7/AwMDBwcHCwsLDw8PE
	xMTFxcXGxsbHx8fIyMjJycnKysrLy8vMzMzNzc3Ozs7Pz8/Q0NDR0dHS0tLT
	09PU1NTV1dXW1tbX19fY2NjZ2dna2trb29vc3Nzd3d3e3t7f39/g4ODh4eHi
	4uLj4+Pk5OTl5eXm5ubn5+fo6Ojp6enq6urr6+vs7Ozt7e3u7u7v7+/w8PDx
	8fHy8vLz8/P09PT19fX29vb39/f4+Pj5+fn6+vr7+/v8/Pz9/f3+/v7///9Q
	NrN6AAAE5UlEQVR4nO2bgZKjIAxAY7ezVu2M//+3pxAgCUHQUnFvzM61QGJ4
	Bgjo7sHrYgKvnkvr+t8Ber9D/YGC9cGJt38+n+afrXdeUP+Lgnrwgno3XFtA
	bwL0eBAiCjSUAf1mgMIESgO9BVDfu69gH4AMCwNi/hCFj4Dh8kB97750oDcD
	QmkJxIfM4TxCnc2hp5eeD1kAYkNGeQ7NIcR5kHoZkJxDv7WAAo4+ZEbCkKFQ
	oLXOgZDn0JARnFZA77ef1kXLngOVDJnjKVz2hUAsIgqQ0yuTmgIVzaFm9Rso
	CzRzITNMl9r2oxCI7DN3kLUfuOTsFSDmLgkEHki3d3oHBB5o238tILvmjb0t
	ArbqQOA2ijpAYI8wplNj75KQtV/KVr82vgyMqRsoY7/S2L3d+XeEByMUzlMG
	aOBAS90CmMxtwgPYkALyDQcjtFg5IiVC/dLzvHYPCOTt6Rz6KhCbQ6vezBnY
	Aorn0GEgoOJXGXh717XVvwZm74AApP9P5hA9lMdARm+BfITCJEoBuc21EhBf
	xvGQRUCR/4+A5CoLfGhvGPA4oq0ybcj6TyIk8pAEkss+ykO1h8w4wG/F3oBA
	3/vEWLJ1pIHk/psDiuwzW0fGfwQUnQcSQPQmN+3Fbp+zj4BkQ2sB8PcCS6XD
	n6VM2/falJSJHwbkBxuGsXNWpkzbQ3kgNuHajpW5vV4OfXUCCK2WzxHQCsu0
	3ZeHIdj4aztgZWGvln1fnYyQTyIrdZctU3tX7tYIkXLOnvrs4ggFI7TZLGsd
	2AiFcs6e+rwj9D9ECNztBOqNMrX3ZbvKoggl7TcjxHJPSR7SyjIPablK5qFg
	zzM1zbxfyMIl9kURSmXeksilIqRfWziHUnOiZG6VtO9eZa5drpqS1VfSvjsP
	uXaZV0ryU0n7HaH/KUKAK5Gfh1LnG2oDiXOP6ic+Y4F+HsrmD5mH9Gy+Pw85
	G5mpBycVz9Q7bRjQeDEB+dgUnp/GM1UEaNalf431VUnJAy0PnCnXH6iix2D3
	uMuBVBPrurLqrwOZzKm7rqTaBQQwrRLeLI31VQSIwqpA9mrjQbqupyKvYCms
	BhQutw6I64oq//aLwypA9HJjFVwvqh/4mfBTqmBaLzWfUqU4fEkeFz4VCMKH
	7BVVWq8Ye/0qKxIogo2BYGJAEwTXMLEISRWJkFBZFk8kgNjdJ4AYNev1iMqy
	4AcDiu/+DCAz82ynP2Y1CSA+ZgeAACCl2sG6C8jZxbNhWu92MjNIUeEsqQuU
	WS/YOmmqjatilQfid5/KQ/5eYyAfoTgPec/KbUiHZJVRniOZOnj+PFNHsM33
	soJMHe13393t83sZejjtPJTf7aV8/cS4cR6KfoHkngTqq4pOjKnLl+urq0qA
	kpe/xjNVAWi8mLA/aBqXH/Jot/GkyaUfo99UfgFoffSFMpmr4WwAmUdfIC9N
	IF2pDeQn+xJ5njeAvFaCdOVEIPLiDdKVUyPkX01CunJyhLBfSFfuCLWJkCEI
	2037CCFQPzeO0JaybR6yL2f7mUfo/Ey9pWwcoR73squssn5WgO48dGfqT4Eu
	GaGLnYdaZmplc215ptY31yL5JhDby8T/KklLTaCKvqoIj9A4imoDoCa9bgg/
	D80XiNDW8aMtkHL8aAwUZ+ob6Aa6geoAhZe1I/+LmkZA5OXxJYDoy+wrAP0D
	7BQA5IVg6IYAAAAASUVORK5CYII=`

const (
	iconClosed = iota
	iconOpened
	iconBomb
	iconMarked
	iconAnswerNoBomb
	iconAnswerIsBomb
	iconQuestionMark
	iconQuestionPressed
)

const (
	buttonPlaying = iota
	buttonEvaluate
	buttonLost
	buttonWon
	buttonPressed
)

const (
	stateThinking = iota
	statePlaying
	stateLost
	stateWon
)

type sprites struct {
	numbers    []*ebiten.Image
	icons      []*ebiten.Image
	digits     []*ebiten.Image
	buttons    []*ebiten.Image
	background *ebiten.Image
}

type config struct {
	scale   int
	width   int
	height  int
	bombs   int
	holding int
}

type game struct {
	c           config
	sprites     sprites
	tiles       [][]int
	bombs       [][]bool
	numbers     [][]int
	bombsLeft   int
	buttonState int
	tilePressed *image.Point
	hitAreas    map[string]image.Rectangle
	startTime   time.Time
}

type transform struct {
	src, dst image.Rectangle
}

func loadImageFromString(b64 string) (*ebiten.Image, error) {
	b64 = strings.ReplaceAll(b64, "\n", "")
	b64 = strings.ReplaceAll(b64, "\t", "")
	b64 = strings.ReplaceAll(b64, " ", "")
	bin, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(bytes.NewReader(bin))
	if err != nil {
		return nil, err
	}
	img2 := ebiten.NewImageFromImage(img)
	return img2, err
}

func (g *game) getHitArea() string {
	x, y := ebiten.CursorPosition()
	cursor := image.Point{x, y}
	for area, rect := range g.hitAreas {
		if cursor.In(rect) {
			return area
		}
	}
	return ""
}

func (g *game) getTileCoords() *image.Point {
	x, y := ebiten.CursorPosition()
	cursor := image.Point{x, y}
	rect := g.hitAreas["tiles"]
	if !cursor.In(rect) {
		return nil
	}
	cx := (x - rect.Min.X) / 16
	cy := (y - rect.Min.Y) / 16
	return &image.Point{cx, cy}
}

func (g *game) clickButton() {
	g.init()
}

func (g *game) gameOver(x, y int) {

}

func (g *game) leftClickTile(x, y int) {
	log.Println("left-click")
	icon := g.tiles[x][y]
	switch icon {
	case iconClosed:
		if g.bombs[x][y] {
			icon = iconAnswerIsBomb

		} else {
			icon = iconOpened
		}
	}
	g.tiles[x][y] = icon
}

func (g *game) rightClickTile(x, y int) {
	log.Println("right-click")
	icon := g.tiles[x][y]
	switch icon {
	case iconClosed:
		icon = iconMarked
	case iconMarked:
		icon = iconClosed
	}
	g.tiles[x][y] = icon
}

func (g *game) Update() error {
	area := g.getHitArea()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		switch area {
		case "button":
			g.buttonState = buttonPressed
		case "tiles":
			g.tilePressed = g.getTileCoords()
			g.buttonState = buttonEvaluate
		}
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		switch area {
		case "button":
			g.buttonState = buttonPlaying
			g.clickButton()
		case "tiles":
			if g.tilePressed != nil {
				coords := g.getTileCoords()
				g.tilePressed = nil
				g.buttonState = buttonPlaying
				g.leftClickTile(coords.X, coords.Y)
			}
		}
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		switch area {
		case "tiles":
			coords := g.getTileCoords()
			g.tilePressed = nil
			g.buttonState = buttonPlaying
			g.rightClickTile(coords.X, coords.Y)
		}
	}
	if g.buttonState == buttonPressed {
		if area != "button" {
			g.buttonState = buttonPlaying
		}
	}
	if g.tilePressed != nil {
		if area == "tiles" {
			g.tilePressed = g.getTileCoords()
		} else {
			g.tilePressed = nil
			g.buttonState = buttonPlaying
		}
	}
	return nil
}

func (g *game) loadBackgroundTile(spritesImage *ebiten.Image) {
	w, h := g.c.width, g.c.height
	width, height := g.getSize()
	background := ebiten.NewImage(width, height)
	transforms := []transform{
		{image.Rect(0, 82, 12, 93), image.Rect(0, 0, 12, 11)},
		{image.Rect(13, 82, 14, 93), image.Rect(12, 0, 12+w*16, 11)},
		{image.Rect(15, 82, 27, 93), image.Rect(12+w*16, 0, 12+w*16+12, 11)},
		{image.Rect(0, 94, 12, 95), image.Rect(0, 11, 12, 11+33)},
		{image.Rect(15, 94, 27, 95), image.Rect(12+w*16, 11, 12+w*16+12, 11+33)},
		{image.Rect(0, 96, 12, 107), image.Rect(0, 11+33, 12, 11+33+11)},
		{image.Rect(13, 96, 14, 107), image.Rect(12, 11+33, 12+w*16, 11+33+11)},
		{image.Rect(15, 96, 27, 107), image.Rect(12+w*16, 11+33, 12+w*16+12, 11+33+11)},
		{image.Rect(0, 108, 12, 109), image.Rect(0, 11+33+11, 12, 11+33+11+h*16)},
		{image.Rect(15, 108, 27, 109), image.Rect(12+w*16, 11+33+11, 12+w*16+12, 11+33+11+h*16)},
		{image.Rect(0, 110, 12, 121), image.Rect(0, 11+33+11+h*16, 12, 11+33+11+h*16+11)},
		{image.Rect(13, 110, 14, 121), image.Rect(12, 11+33+11+h*16, 12+w*16, 11+33+11+h*16+11)},
		{image.Rect(15, 110, 27, 121), image.Rect(12+w*16, 11+33+11+h*16, 12+w*16+12, 11+33+11+h*16+11)},
		{image.Rect(28, 82, 69, 107), image.Rect(12+4, 11+4, 12+4+41, 11+4+25)},
		{image.Rect(28, 82, 69, 107), image.Rect(12+w*16-4-41, 11+4, 12+w*16-4, 11+4+25)},
	}
	for _, t := range transforms {
		source := spritesImage.SubImage(t.src).(*ebiten.Image)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(t.dst.Size().X)/float64(t.src.Size().X), float64(t.dst.Size().Y)/float64(t.src.Size().Y))
		op.GeoM.Translate(float64(t.dst.Min.X), float64(t.dst.Min.Y))
		background.DrawImage(source, op)
	}
	g.sprites.background = background
}

func (g *game) init() *game {
	spritesImage, err := loadImageFromString(skin)
	if err != nil {
		log.Fatalln("Could not load skin")
	}
	g.sprites.numbers = make([]*ebiten.Image, 9)
	for x := range g.sprites.numbers {
		g.sprites.numbers[x] = spritesImage.SubImage(image.Rect(x*16, 0, x*16+16, 16)).(*ebiten.Image)
	}
	g.sprites.icons = make([]*ebiten.Image, 8)
	for x := range g.sprites.icons {
		g.sprites.icons[x] = spritesImage.SubImage(image.Rect(x*16, 16, x*16+16, 32)).(*ebiten.Image)
	}
	g.sprites.digits = make([]*ebiten.Image, 11)
	for x := range g.sprites.digits {
		g.sprites.digits[x] = spritesImage.SubImage(image.Rect(x*12, 33, x*12+11, 54)).(*ebiten.Image)
	}
	g.sprites.buttons = make([]*ebiten.Image, 5)
	for x := range g.sprites.buttons {
		g.sprites.buttons[x] = spritesImage.SubImage(image.Rect(x*27, 55, x*27+26, 81)).(*ebiten.Image)
	}
	g.loadBackgroundTile(spritesImage)
	g.initTiles()
	g.initBombs()
	g.initNumbers()
	g.buttonState = buttonPlaying
	g.bombsLeft = g.c.bombs
	g.hitAreas = make(map[string]image.Rectangle)
	g.startTime = time.Now()
	return g
}

func (g *game) initTiles() {
	g.tiles = make([][]int, g.c.width)
	for x := range g.tiles {
		g.tiles[x] = make([]int, g.c.height)
	}
}

func (g *game) initBombs() {
	g.bombs = make([][]bool, g.c.width)
	for x := range g.tiles {
		g.bombs[x] = make([]bool, g.c.height)
	}
	b := g.c.bombs
	for b > 0 {
		x, y := rand.Intn(g.c.width), rand.Intn(g.c.height)
		if !g.bombs[x][y] {
			g.bombs[x][y] = true
			b--
		}
	}
}

func (g *game) initNumbers() {
	g.numbers = make([][]int, g.c.width)
	for x := range g.tiles {
		g.numbers[x] = make([]int, g.c.height)
	}
	w, h := g.c.width, g.c.height
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			g.numbers[x][y] = g.calucateBombs(x, y)
		}
	}
}

func (g *game) drawBackground(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.RGBA{0xcc, 0xcc, 0xcc, 0xcc})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.sprites.background, op)
}

func (g *game) calucateBombs(x, y int) int {
	bombs := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if x+dx >= 0 && x+dx < g.c.width {
				if y+dy >= 0 && y+dy < g.c.height {
					if g.bombs[x+dx][y+dy] {
						bombs++
					}
				}
			}
		}
	}
	return bombs
}

func (g *game) drawTiles(screen *ebiten.Image) {
	w, h := g.c.width, g.c.height
	g.hitAreas["tiles"] = image.Rect(12, 11+33+11, 12+w*16, 11+33+11+h*16)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(12+x*16), float64(11+33+11+y*16))
			icon := g.tiles[x][y]
			if icon == iconOpened {
				number := g.numbers[x][y]
				screen.DrawImage(g.sprites.numbers[number], op)
			} else {
				if g.tilePressed != nil && g.tilePressed.X == x && g.tilePressed.Y == y {
					icon = iconOpened
				}
				screen.DrawImage(g.sprites.icons[icon], op)
			}
		}
	}
}

func (g *game) drawBombsLeft(screen *ebiten.Image) {
	for i := 0; i < 3; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(12+4+2+i*(11+2)), float64(11+4+2))
		digit := (g.bombsLeft / int(math.Pow(10, float64(2-i)))) % 10
		screen.DrawImage(g.sprites.digits[digit], op)
	}
}

func (g *game) drawButton(screen *ebiten.Image) {
	width, _ := g.getSize()
	g.hitAreas["button"] = image.Rect(width/2-13, 11+4, width/2+13, 11+4+26)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(width/2-13), float64(11+4))
	screen.DrawImage(g.sprites.buttons[g.buttonState], op)
}

func (g *game) drawSecondsPassed(screen *ebiten.Image) {
	for i := 0; i < 3; i++ {
		op := &ebiten.DrawImageOptions{}
		width, _ := g.getSize()
		op.GeoM.Translate(float64(width-(12+4+(3-i)*(11+2))), float64(11+4+2))
		secondsPassed := int((time.Now().UnixNano() - g.startTime.UnixNano()) / 1000000000)
		digit := (secondsPassed / int(math.Pow(10, float64(2-i)))) % 10
		screen.DrawImage(g.sprites.digits[digit], op)
	}
}

func (g *game) Draw(screen *ebiten.Image) {
	g.drawBackground(screen)
	g.drawBombsLeft(screen)
	g.drawButton(screen)
	g.drawSecondsPassed(screen)
	g.drawTiles(screen)
}

func (g *game) getSize() (int, int) {
	return g.c.width*16 + 12*2, g.c.height*16 + 11*3 + 33
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.getSize()
}

func main() {
	ebiten.SetWindowTitle("Minesweeper.go")
	g := &game{c: config{
		scale:   5,
		width:   9,
		height:  9,
		bombs:   10,
		holding: 15,
	}}
	width, height := g.getSize()
	ebiten.SetMaxTPS(6)
	ebiten.SetWindowSize(g.c.scale*width, g.c.scale*height)
	if err := ebiten.RunGame(g.init()); err != nil {
		log.Fatalf("%v\n", err)
	}
}
