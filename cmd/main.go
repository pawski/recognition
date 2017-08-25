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

	transformations := []string{"blur", "gray_scale", "contrast", "sharpen", "invert", "emboss"}

	for _, transformationType := range transformations {
		log.Println(transformationType)
		go func(src image.Image, transformationType string){
			log.Println(transformationType)
			transform(src, transformationType)
			done <- true
		}(src, transformationType)
	}

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