package main

import "github.com/fogleman/gg"




func main() {
	params := sketch.UserParams{
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
	}
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
