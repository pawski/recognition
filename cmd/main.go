package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"github.com/disintegration/imaging"
	"os"
)

func main() {
	fmt.Printf("Strating app.\n")

	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	imageSourcePath := os.Args[1];

	src, err := imaging.Open(imageSourcePath)

	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}

	// Create a blurred version of the image.
	img1 := imaging.Blur(src, 2)

	// Create a grayscale version of the image with higher contrast and sharpness.
	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2, 20)
	img2 = imaging.Sharpen(img2, 2)

	// Create an inverted version of the image.
	img3 := imaging.Invert(src)

	// Create an embossed version of the image using a convolution filter.
	img4 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)

	// Create a new image and paste the four produced images into it.
	dst := imaging.New(512, 512, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	dst = imaging.Paste(dst, img2, image.Pt(0, 256))
	dst = imaging.Paste(dst, img3, image.Pt(256, 0))
	dst = imaging.Paste(dst, img4, image.Pt(256, 256))

	// Save the resulting image using JPEG format.
	err = imaging.Save(dst, "./cache/collage.jpg")

	imaging.Save(img1, "./cache/blur.jpg")
	imaging.Save(img2, "./cache/gray_scale.jpg")
	imaging.Save(img3, "./cache/contrast.jpg")
	imaging.Save(img4, "./cache/sharpen.jpg")

	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}
}
