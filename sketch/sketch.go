package sketch

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type UserParams struct {
	DestWidth                int
	DestHeight               int
	StrokeRatio              float64
	StrokeReduction          float64
	StrokeJitter             int
	StrokeInventionThreshold float64
	InitialAlpha             float64
	AlphaIncrease            float64
	MinEdgeCount             int
	MaxEdgeCount             int
}

type Sketch struct {
	UserParams
	source            image.Image
	dc                *gg.Context
	sourceWidth       int
	sourceHeight      int
	strokeSize        float64
	initialStrokeSize float64
}

func NewSketch(source image.Image, userParams UserParams) *Sketch {
	s := &Sketch{UserParams: userParams}
	bounds := source.Bounds()
	s.sourceWidth, s.sourceHeight = bounds.Max.X, bounds.Max.Y
	initialStrokeSize := s.StrokeRatio * float64(s.DestWidth)
	s.strokeSize = initialStrokeSize

	canvas := gg.NewContext(s.DestWidth, s.DestHeight)
	canvas.SetColor(color.Black)
	canvas.DrawRectangle(0, 0, float64(s.DestWidth), float64(s.DestHeight))
	canvas.Fill()

	s.source = source
	s.dc = canvas
	return s
}

func (s *Sketch) Update() {
	//	Obtain Information color of the source
	rndX := rand.Float64() * float64(s.sourceWidth)
	rndY := rand.Float64() * float64(s.sourceHeight)
	r, g, b := rgb255(s.source.At(int(rndX), int(rndY)))

	// Determine the destination in the output image
	destX := rndX * float64(s.DestWidth) / float64(s.sourceWidth)
	destX += float64(randRange(s.StrokeJitter))
	destY := rndY * float64(s.DestHeight) / float64(s.sourceHeight)
	destY += float64(randRange(int(s.StrokeJitter)))

	// Draw a stroke using the parameters
	edges := s.MinEdgeCount + rand.Intn(s.MaxEdgeCount-s.MinEdgeCount)

	s.dc.SetRGBA255(r, g, b, int(s.InitialAlpha))
	s.dc.DrawRegularPolygon(edges, destX, destY, s.strokeSize, rand.Float64())
	s.dc.FillPreserve()

	if s.strokeSize <= s.StrokeInventionThreshold*s.initialStrokeSize {
		if (r+g+b)/3 < 128 {
			s.dc.SetRGBA255(255, 255, 255, int(s.InitialAlpha*2))
		} else {
			s.dc.SetRGBA255(0, 0, 0, int(s.InitialAlpha*2))
		}
	}
	s.dc.Stroke()

	// Update the parameter state for the next execution
	s.strokeSize -= s.StrokeReduction * s.strokeSize
	s.InitialAlpha += s.AlphaIncrease
}

func rgb255(c color.Color) (int, int, int) {
	r, g, b, _ := c.RGBA()
	return int(r / 255), int(g / 255), int(b / 255)
}

func randRange(max int) int {
	return -max + rand.Intn(2*max)
}

func (s *Sketch) Output() image.Image {
	return s.dc.Image()
}
