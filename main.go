package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/esvarez/donatello/sketch"
)

var (
	sourceImgName   = "source/img1.jpg"
	outputImgName   = "out/img.png"
	totalCycleCount = 100000
)

func main() {
	rand.NewSource(time.Now().Unix())

	img, err := loadImage(sourceImgName)
	if err != nil {
		log.Panic(err)
	}

	destWith := 2000

	newSketch := sketch.NewSketch(img, sketch.UserParams{
		DestWidth:                destWith,
		DestHeight:               2000,
		StrokeRatio:              0.75,
		StrokeReduction:          0.002,
		InitialAlpha:             0.1,
		AlphaIncrease:            0.36,
		StrokeInventionThreshold: 0.05,
		StrokeJitter:             int(0.1 * float64(destWith)),
		MinEdgeCount:             3,
		MaxEdgeCount:             8,
	})

	for i := 0; i < totalCycleCount; i++ {
		newSketch.Update()
	}

	if err := saveOutputWithCheck(newSketch.Output(), outputImgName); err != nil {
		log.Panic(err)
	}
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

func saveOutputWithCheck(img image.Image, filePath string) error {
	ext := filepath.Ext(filePath)
	baseName := strings.TrimSuffix(filePath, ext)

	if _, err := os.Stat(filePath); err == nil {
		newFilePath, err := getNextFileName(baseName, ext)
		if err != nil {
			return err
		}
		filePath = newFilePath
	}
	return saveOutput(img, filePath)
}

func getNextFileName(baseName, ext string) (string, error) {
	dir := filepath.Dir(baseName)
	base := filepath.Base(baseName)
	base = strings.TrimSuffix(base, ext)

	for i := 1; ; i++ {
		newFileName := fmt.Sprintf("%s_%d%s", base, i, ext)
		newFilePath := filepath.Join(dir, newFileName)
		if _, err := os.Stat(newFilePath); os.IsNotExist(err) {
			return newFilePath, nil
		}
	}
}

func saveOutput(img image.Image, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return nil
	}

	return nil
}
