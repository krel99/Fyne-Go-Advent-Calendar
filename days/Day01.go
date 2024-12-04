package days

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Day01(rectangles []fyne.CanvasObject) {
	if len(rectangles) == 0 {
		return
	}

	container, ok := rectangles[0].(*fyne.Container)
	if !ok {
		return
	}

	rect, ok := container.Objects[0].(*canvas.Rectangle)
	if !ok {
		return
	}

	size := rect.Size()
	centerX := size.Width / 2
	centerY := size.Height / 2

	outerRadius := size.Width / 3
	innerRadius := outerRadius / 2
	pointCount := 10

	// Create lines for the star
	for i := 0; i < pointCount; i++ {
		angle := float64(i) / float64(pointCount) * math.Pi * 2
		angleNext := float64(i+1) / float64(pointCount) * math.Pi * 2

		// Calculate current point
		radius := outerRadius
		if i%2 == 1 {
			radius = innerRadius
		}
		x1 := centerX + float32(math.Sin(angle)*float64(radius))
		y1 := centerY - float32(math.Cos(angle)*float64(radius))

		// Calculate next point
		nextRadius := innerRadius
		if i%2 == 1 {
			nextRadius = outerRadius
		}
		x2 := centerX + float32(math.Sin(angleNext)*float64(nextRadius))
		y2 := centerY - float32(math.Cos(angleNext)*float64(nextRadius))

		// Create and add line
		// line := canvas.NewLine(color.White)
		// line.StrokeWidth = 2
		// line.Position1 = fyne.NewPos(x1, y1)
		// line.Position2 = fyne.NewPos(x2, y2)
		// container.Add(line)

		go daVincify(x1, y1, x2, y2, 200, container)
	}

	rect.FillColor = color.NRGBA{R: 100, G: 100, B: 0, A: 255}
	container.Refresh()
}

func daVincify(x1, y1, x2, y2, split float32, ctx *fyne.Container) {
	xIncrement := (x2 - x1) / split
	yIncrement := (y2 - y1) / split

	currentX := x1
	currentY := y1

	for i := float32(0); i <= split; i++ {
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		line.Position1 = fyne.NewPos(currentX, currentY)
		line.Position2 = fyne.NewPos(currentX+xIncrement, currentY+yIncrement)
		ctx.Add(line)

		currentX += xIncrement
		currentY += yIncrement

		time.Sleep(15 * time.Millisecond)

		ctx.Refresh()
	}
}
