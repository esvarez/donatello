package sketch

type UserParams struct {
	DesWith int
	DesHeight int
	StrokeRatio float64
	StrokeReduction float64
	StrokeJitter float64
	StrokeInventionThreshold float64
	InitialAlpha float64
	AlphaIncrease float64
	MinEdgeCount int
	MaxEdgeCount int
}

type Sketch struct {
	UserParams 
	source image.Image
	dc *gg.Context
	sourceWidth int 
	sourceHeight int
	strokeSize float64
	initialStrokeSize float64
}

func NewSketch(source image.Image, userParams UserParams) *Sketch {
	s := &Sketch{UserParams: userParams}
	bounds := source.Bounds()
	s.sourceWidth, s.sourceHeight = bounds.Max.X, bounds.Max.Y
	initialStrokeSize = s.StrokeRatio * float64(s.DesWidth)
	s.strokeSize = initialStrokeSize 

	canvas := gg.NewContext(s.DesWidth, s.DesHeight)
	canvas.SetColor(color.black)
	canvas.DrawRectangle(0,0, float64(s.DestWith), float64(s.DestHeight))
	canvas.Fill()

	s.source = source 
	s.dc = canvas 
	return s
}

func (s *Sketch) Update() {
	// TODO
}

func (s *Sketch) Output() {
	// TODO
}
