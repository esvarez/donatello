package main

import "github.com/fogleman/gg"


var (
	sourceImgName = "source/img1.jpg"
	outImgName = "out/img.ong"
	totalCycleCount = 500
)

func main() {
	rand.Seed(time.Now().Unix())


	img, err := loadImage(sourceImage)
	if err != nil {
		log.Panic(err)
	}

	desWith := 2000

	sketch := sketch.NewSketch(img, sketch.UserParams{
		DesWidth: desWith,
		DesHeight: 2000,
		StrokeRatio: 0.75,
		StrokeReduction: 0.002,
		StrokeInventionThreshold : 0.05,
		StrokeJitter: int(0.1 * float64(desWidth)),
		InitialAlpha: 0.1,
		AlphaIncrease: 0.36,
		MinEdgeCount: 3,
		MaxEdgeCount:4,
	})

	for i := 0; i<totalCycleCount; i++ {
		sketch.Update()
	}

	saveOutput(s.Output(), outImageName)
}

func saveOutput (img image.Image, filePath string) error {
	f, err := onCreate(filePath)
	if err != nil {
		return err 
	}
	defer.Close()

	if err := png.Encode(f, img); err != nil {
		return nil
	}

	return nil
}
