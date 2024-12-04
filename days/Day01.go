package days

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Day01(rectangles []fyne.CanvasObject) {
	if len(rectangles) > 0 {
		if rect, ok := rectangles[0].(*canvas.Rectangle); ok {
			rect.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255} // Bright red
			rect.Refresh()
		}
	}
}
