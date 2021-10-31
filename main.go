package unicodesemigraphics

import (
	"io"
	"strings"
)

// https://www.compart.com/en/unicode/block/U+2580

var bitToBlockTable = []string{
	" ", "\u2598", "\u259D", "\u2580",
	"\u2596", "\u258C", "\u259E", "\u259B",
	"\u2597", "\u259A", "\u2590", "\u259C",
	"\u2584", "\u2599", "\u259F", "\u2588"}

type Cell byte

func (c Cell) String() string {
	return bitToBlockTable[c]
}

func _bits(x, y int) Cell {
	return Cell(1 << (x + (y << 1)))
}

func (c Cell) Get(x, y int) bool {
	return (c & _bits(x, y)) > 0
}

func (c *Cell) Set(x, y int, value bool) {
	if value {
		*c |= _bits(x, y)
	} else {
		*c &^= _bits(x, y)
	}
}

type Bitmap struct {
	lines  [][]Cell
	width  int
	height int
}

func NewBitmap(width, height int) *Bitmap {
	lines := make([][]Cell, (height+1)/2)
	length := (width + 1) / 2
	for i := range lines {
		lines[i] = make([]Cell, length)
	}
	return &Bitmap{
		width:  width,
		height: height,
		lines:  lines,
	}
}

func (bmp *Bitmap) Get(x, y int) bool {
	return bmp.lines[y/2][x/2].Get(x%2, y%2)
}

func (bmp *Bitmap) Set(x, y int, value bool) {
	bmp.lines[y/2][x/2].Set(x%2, y%2, value)
}

func (bmp *Bitmap) WriteTo(w io.Writer) (int64, error) {
	rs := ""
	n := 0
	for _, line := range bmp.lines {
		m, _ := io.WriteString(w, rs)
		n += m
		rs = "\n"
		for _, cell := range line {
			m, _ := io.WriteString(w, cell.String())
			n += m
		}
	}
	return int64(n), nil
}

func (bmp *Bitmap) String() string {
	var buffer strings.Builder
	bmp.WriteTo(&buffer)
	return buffer.String()
}
