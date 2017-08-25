package main

import (
	"image"
	//"image/color"
	"log"
	"github.com/disintegration/imaging"
	"os"
)

func main() {
	log.Print("Strating app.\n")

	argsWithoutProg := os.Args[1:]
	log.Print(argsWithoutProg)

	imageSourcePath := os.Args[1];

	src, err := imaging.Open(imageSourcePath)

	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}

	done := make(chan bool, 1)

	go func(src image.Image){
		transform(src, "blur")
		done <- true
	}(src)

	go func(src image.Image, done chan bool) {
		transform(src, "gray_scale")
		done <- true
	}(src, done)

	go func(src image.Image) {
		transform(src, "contrast")
		done <- true
	}(src)

	go func(src image.Image) {
		transform(src, "sharpen")
		done <- true
	}(src)

	go func(src image.Image) {
		transform(src, "invert")
		done <- true
	}(src)

	go func(src image.Image) {
		transform(src, "emboss")
		done <- true
	}(src)

	for i := 0; i < 6; i++ {
		select {
		case <-done:
		}
	}
}

func transform(img image.Image, processType string) {

	log.Printf("%s\n", processType + " start")

	var processedImage image.Image

	switch processType {
	case "gray_scale":
		processedImage = imaging.Grayscale(img)
	case "blur":
		processedImage = imaging.Blur(img, 2)
	case "contrast":
		processedImage = imaging.AdjustContrast(img, 20)
	case "sharpen":
		processedImage = imaging.Sharpen(img, 2)
	case "invert":
		processedImage = imaging.Invert(img)
	case "emboss":
		processedImage = imaging.Convolve3x3(
		img,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)
	default:
		log.Panic("Unknown transform type")
	}

	err := imaging.Save(processedImage, "./cache/" + processType + ".jpg")

	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}

	log.Printf("%s\n", processType)
}