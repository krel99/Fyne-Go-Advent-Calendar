package days

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Day04(rectangles []fyne.CanvasObject) {
	if len(rectangles) == 0 {
		return
	}

	container, ok := rectangles[3].(*fyne.Container)
	if !ok {
		return
	}

	rect, ok := container.Objects[0].(*canvas.Rectangle)
	if !ok {
		return
	}

	// unniversal coordinates
	size := rect.Size()
	eightX := size.Width / 8
	eightY := size.Height / 8

	// horizontal line
	line := canvas.NewLine(color.Black)
	line.Position1 = fyne.NewPos(eightX*1, eightY*6)
	line.Position2 = fyne.NewPos(eightX*7, eightY*6)
	line.StrokeWidth = 5
	container.Add(line)

	// curved lines up from the base
	// ! this is suboptimal approach - too many variables declared and improper naming

	startX_right := eightX * 7
	startY := eightY * 6
	endX_right := eightX * 5.5
	endY := eightY * 4
	startX_left := eightX * 1
	endX_left := eightX * 2.5

	go drawSideKurvatures(6, 1, eightX, startX_right, startY, endX_right, endY, &*container)
	go drawSideKurvatures(6, -1, eightX, startX_left, startY, endX_left, endY, &*container)

	// draw relatively straight horizontals

	thEndRight := eightX * 5
	thEndLeft := eightX * 3
	thEnd := eightY * 2

	go drawTiltedHorizontals(endX_right, endY, thEndRight, thEnd, &*container)
	go drawTiltedHorizontals(endX_left, endY, thEndLeft, thEnd, &*container)

	// top curvature
	numSegmentsCircle := 10
	centerX := eightX * 4 // Center of the circle
	topY := eightY * 2    // Y coordinate where our straight line ended
	radius := eightX * 1  // Radius of 1/8th of width

	for i := 0; i < numSegmentsCircle; i++ {
		// Go from 0 to π for the top half circle
		t := float64(i) / float64(numSegmentsCircle) * math.Pi
		t_next := float64(i+1) / float64(numSegmentsCircle) * math.Pi

		x1 := centerX + radius*float32(math.Cos(t))
		y1 := topY - radius*float32(math.Sin(t))

		x2 := centerX + radius*float32(math.Cos(t_next))
		y2 := topY - radius*float32(math.Sin(t_next))

		segment := canvas.NewLine(color.Black)
		segment.Position1 = fyne.NewPos(x1, y1)
		segment.Position2 = fyne.NewPos(x2, y2)
		segment.StrokeWidth = 5

		container.Add(segment)
	}

	// Refresh the context
	container.Refresh()

}

func drawSideKurvatures(segments, curveIx int, eightX, startX, startY, endX, endY float32, container *fyne.Container) {

	for i := 0; i < segments; i++ {
		t := float64(i) / float64(segments)
		t_next := float64(i+1) / float64(segments)

		// curving here
		curveX := func(t float64) float32 {
			return startX - (startX-endX)*float32(t) - float32(curveIx)*eightX*float32(t*(1-t))
		}

		x1 := curveX(t)
		y1 := startY - (startY-endY)*float32(t)

		x2 := curveX(t_next)
		y2 := startY - (startY-endY)*float32(t_next)

		segment := canvas.NewLine(color.Black)
		segment.Position1 = fyne.NewPos(x1, y1)
		segment.Position2 = fyne.NewPos(x2, y2)
		segment.StrokeWidth = 5

		container.Add(segment)
	}

}

func drawTiltedHorizontals(startX, startY, endX, endY float32, container *fyne.Container) {
	// relatively straight up lines
	horizontal := canvas.NewLine(color.Black)
	horizontal.Position1 = fyne.NewPos(startX, startY)
	horizontal.Position2 = fyne.NewPos(endX, endY)
	horizontal.StrokeWidth = 5

	container.Add(horizontal)
}
