package palette

import (
	"image"
	"image/color"
	"os"
	"sort"

	_ "github.com/gonutz/bmp"
	_ "image/jpeg"
	_ "image/png"
)

// Palette is a list of colors and their occurrence counts.
type Palette []ColorCount

// ColorCount contains a color and the number of its occurrences.
type ColorCount struct {
	Color color.RGBA
	Count int
}

// ExtractPaletteFromImageFiles looks at all pixels in the given list of image
// files and enumerates all different colors into a Palette. The Palette is
// sorted from most common (most occurrences) color to least common color.
func ExtractPaletteFromImageFiles(paths ...string) (Palette, error) {
	colors := make(map[color.RGBA]int)

	for _, path := range paths {
		img, err := loadImage(path)
		if err != nil {
			return Palette{}, err
		}

		b := img.Bounds()
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				r, g, b, a := img.At(x, y).RGBA()
				colors[color.RGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: uint8(a >> 8),
				}]++
			}
		}
	}

	pal := make(Palette, 0, len(colors))
	for col, count := range colors {
		pal = append(pal, ColorCount{Color: col, Count: count})
	}
	sort.Slice(pal, func(i, j int) bool {
		return pal[i].Count > pal[j].Count
	})

	return pal, nil
}

func loadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	return img, err
}
