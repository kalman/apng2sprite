package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"

	"github.com/kettek/apng"
)

func main() {
	os.Exit(main2())
}

func main2() int {
	inFilename := flag.String("i", "", "input file")
	outFilename := flag.String("o", "", "output file, leave empty for stdout")
	flag.Parse()

	if *inFilename == "" || flag.NArg() != 0 {
		flag.Usage()
		return 1
	}

	inFile, err := os.Open(*inFilename)
	defer inFile.Close()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	apngDecode, err := apng.DecodeAll(inFile)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	images := make([]image.Image, len(apngDecode.Frames))
	for i, frame := range apngDecode.Frames {
		images[i] = frame.Image
	}
	ssi := newSpriteSheetImage(images)

	var outFile io.Writer
	if *outFilename != "" {
		var f *os.File
		if f, err = os.Create(*outFilename); err == nil {
			outFile = f
			defer f.Close()
		}
	} else {
		outFile = os.Stdout
	}
	if err != nil {
		fmt.Println(err)
		return 1
	}

	err = png.Encode(outFile, ssi)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func newSpriteSheetImage(images []image.Image) *spriteSheetImage {
	bounds := images[0].Bounds()
	return &spriteSheetImage{
		images: images,
		width:  bounds.Max.X - bounds.Min.X,
		height: bounds.Max.Y - bounds.Min.Y,
	}
}

type spriteSheetImage struct {
	images        []image.Image
	width, height int
}

func (ssi *spriteSheetImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (ssi *spriteSheetImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: ssi.width * len(ssi.images), Y: ssi.height},
	}
}

func (ssi *spriteSheetImage) At(x, y int) color.Color {
	return ssi.images[x/ssi.width].At(x%ssi.width, y)
}
