package intcode

import (
	"fmt"
	"image"

	"github.com/dds/aoc2019/lib"
	"github.com/gdamore/tcell"
)

var directions = []image.Point{
	image.Point{0, 1},
	image.Point{0, -1},
	image.Point{-1, 0},
	image.Point{1, 0},
}

const (
	NoDirection = iota
	North
	South
	West
	East
)

func Direction(p, q image.Point) int {
	d := q.Sub(p)
	if d.X != 0 {
		d.X = lib.Sign(d.X)
	}
	if d.Y != 0 {
		d.Y = lib.Sign(d.Y)
	}
	for i, v := range directions {
		if v == d {
			return i + 1
		}
	}
	return NoDirection
}

func Neighbors(p image.Point) (r []image.Point) {
	for _, q := range directions {
		r = append(r, p.Add(q))
	}
	return
}

type Cell struct {
	Rune rune
	tcell.Style
}

type Cells map[image.Point]Cell

func (c Cells) Size() image.Rectangle {
	var minX, minY, maxX, maxY int
	for p, _ := range c {
		maxX = lib.Max(maxX, p.X)
		maxY = lib.Max(maxY, p.Y)
		minX = lib.Min(minX, p.X)
		minY = lib.Min(minY, p.Y)
	}
	return image.Rect(minX, minY, maxX, maxY)
}

// NextUnexplored returns the next unexplored direction or a zero point and NoDirection.
func (c Cells) NextUnexplored(p image.Point) (image.Point, int) {
	// Find the first unexplored direction.
	for i, q := range Neighbors(p) {
		if _, ok := c[q]; !ok {
			return q, i + 1
		}
	}
	return image.Point{}, NoDirection
}

type Screen interface {
	tcell.Screen
	Cells() Cells
	Draw()
}

type Printer interface {
	Print(string)
}

// Implements Screen and Logger.
type grid struct {
	tcell.Screen
	cells Cells
	scene int
	msg   string
}

func (g *grid) Print(s string) { g.msg = s }
func (g *grid) Cells() Cells   { return g.cells }

func (g *grid) Draw() {
	g.Clear()
	for p, cell := range g.cells {
		g.SetContent(p.X, p.Y+4, cell.Rune, nil, tcell.StyleDefault)
	}
	g.scene++
	for i, c := range fmt.Sprintf("Scene: %v", g.scene) {
		g.SetContent(i, 0, c, nil, tcell.StyleDefault)
	}
	var i, line int = 0, 1
	for _, ch := range g.msg {
		if ch == '\n' {
			line++
			i = 0
			continue
		}
		g.SetContent(i, line, ch, nil, tcell.StyleDefault)
		i++
	}
	g.Show()
}

func NewGrid() Screen {
	sc, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := sc.Init(); err != nil {
		panic(err)
	}
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	sc.SetStyle(style)

	g := &grid{
		Screen: sc,
		cells:  make(Cells),
	}
	return g
}
