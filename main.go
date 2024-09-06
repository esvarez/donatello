package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/esvarez/donatello/sketch"
)

var (
	sourceImgName   = "source/img1.jpg"
	outImgName      = "out/img.ong"
	totalCycleCount = 500
)

func main() {
	rand.NewSource(time.Now().Unix())

	img, err := loadImage(sourceImgName)
	if err != nil {
		log.Panic(err)
	}

	desWith := 2000

	sketch := sketch.NewSketch(img, sketch.UserParams{
		DesWidth:                 desWith,
		DesHeight:                2000,
		StrokeRatio:              0.75,
		StrokeReduction:          0.002,
		StrokeInventionThreshold: 0.05,
		StrokeJitter:             int(0.1 * float64(desWidth)),
		InitialAlpha:             0.1,
		AlphaIncrease:            0.36,
		MinEdgeCount:             3,
		MaxEdgeCount:             4,
	})

	for i := 0; i < totalCycleCount; i++ {
		sketch.Update()
	}

	saveOutput(s.Output(), outImgName)
}

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("source file could not be loaded: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("source image format could not be decoded: %w", err)
	}

	return img, nil

}

func saveOutput(img image.Image, filePath string) error {
	f, err := onCreate(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return nil
	}

	return nil
}
