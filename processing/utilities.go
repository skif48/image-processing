package processing

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
)

func RgbToGrayscale(img image.Image, buff *bytes.Buffer) error {
	var err error

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			oldPixel := img.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			newImg.Set(x, y, pixel)
		}
	}

	if err = jpeg.Encode(buff, newImg, nil); err != nil {
		return err
	}

	return nil
}
