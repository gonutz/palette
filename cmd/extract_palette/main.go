package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/gonutz/palette"
)

func main() {
	paletteFile := flag.String(
		"pal",
		"",
		"Palette image file to be written, no file created if empty string",
	)
	flag.Parse()

	pal, err := palette.ExtractPaletteFromImageFiles(flag.Args()...)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
		return
	}

	fmt.Println("Palette:")
	fmt.Println("index) occurrences:  R   G   B   A")
	for i, c := range pal {
		fmt.Printf(
			"%5d) %11d: %3d %3d %3d %3d\n",
			i,
			c.Count,
			c.Color.R,
			c.Color.G,
			c.Color.B,
			c.Color.A,
		)
	}

	if *paletteFile != "" {
		palImage := image.NewRGBA(image.Rect(0, 0, len(pal), 1))
		for i, c := range pal {
			palImage.Set(i, 0, c.Color)
		}

		f, err := os.Create(*paletteFile)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
			return
		}
		defer f.Close()

		if err := png.Encode(f, palImage); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
			return
		}
	}
}
