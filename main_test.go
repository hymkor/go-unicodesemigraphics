package unicodesemigraphics_test

import (
	"testing"

	"github.com/zetamatta/go-unicodesemigraphics"
)

func TestNewBitmap(t *testing.T) {
	bmp := unicodesemigraphics.NewBitmap(3, 3)
	bmp.Set(0, 0, false)
	bmp.Set(1, 0, true)
	bmp.Set(2, 0, false)

	bmp.Set(0, 1, true)
	bmp.Set(1, 1, false)
	bmp.Set(2, 1, true)

	bmp.Set(0, 2, false)
	bmp.Set(1, 2, true)
	bmp.Set(2, 2, false)

	result := bmp.String()
	expect := "\u259E\u2596\n\u259D "
	if result != expect {
		t.Fatalf("expect `%s` but result is `%s`", expect, result)
	}
}
