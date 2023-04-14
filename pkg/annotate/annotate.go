package annotate

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/draychev/image-manipulation/pkg/logger"
)

var log = logger.NewPretty("image-manipulation/annotate")

func Annotate(inputFileName, outputFileName, annotationText string) error {
	imgFile, err := os.Open(inputFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Could not decode image file %s", inputFileName)
		return err
	}
	defer imgFile.Close()

	decodedImage, err := png.Decode(imgFile)
	if err != nil {
		log.Error().Err(err).Msgf("Could not decode image file %s", imgFile.Name())
		return err
	}

	// Create a new image with the same dimensions as the original image, plus a white rectangle at the top
	newImg := image.NewRGBA(image.Rect(0, 0, decodedImage.Bounds().Dx(), decodedImage.Bounds().Dy()+100))

	// Fill the top rectangle with white color
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.ZP, draw.Src)

	// Draw the original image on top of the white rectangle
	draw.Draw(newImg, image.Rect(0, 100, decodedImage.Bounds().Dx(), decodedImage.Bounds().Dy()+100), decodedImage, image.Point{0, 0}, draw.Src)

	// Add the "Hello World" text to the white rectangle
	draw.Draw(newImg, image.Rect(0, 0, decodedImage.Bounds().Dx(), 100), &image.Uniform{C: color.Black}, image.ZP, draw.Src)
	drawString(newImg, annotationText, image.Point{X: 20, Y: 40}, color.White)

	// Save the new image to a file
	outputImageFile, err := os.Create(outputFileName)
	if err != nil {
		log.Error().Err(err).Msgf("Could not create file %s", outputFileName)
		return err
	}
	defer outputImageFile.Close()
	if err := png.Encode(outputImageFile, newImg); err != nil {
		log.Error().Err(err).Msgf("Could not encode image into file %s", outputImageFile.Name())
		return err
	}

	return nil
}

func drawString(img *image.RGBA, s string, sp image.Point, c color.Color) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: basicfont.Face7x13,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6(sp.X * 64),
			Y: fixed.Int26_6(sp.Y * 64),
		},
	}
	d.DrawString(s)
}
