package days

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Day03(rectangles []fyne.CanvasObject) {
	if len(rectangles) <= 1 {
		return
	}

	container, ok := rectangles[2].(*fyne.Container)
	if !ok {
		return
	}

	rect, ok := container.Objects[0].(*canvas.Rectangle)
	if !ok {
		return
	}

	size := rect.Size()
	xC := size.Width / 2
	yC := size.Height / 2

	// COOKIE BASE
	cookieW := size.Width / 10 * 4
	cookieH := size.Height / 10 * 4

	cookieColor := color.NRGBA{R: 205, G: 155, B: 100, A: 255} // Beige/cookie color
	cookie := canvas.NewCircle(cookieColor)
	cookie.Resize(fyne.NewSize(cookieW, cookieH))
	cookie.Move(fyne.NewPos(xC-cookieW/2, yC-cookieH/2))
	container.Add(cookie)

	// CHOCOLATE CHIPS
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chipColor := color.NRGBA{R: 60, G: 30, B: 10, A: 255} // Exkrementova barva

	chipSize := cookieW / 8
	for i := 0; i < 6; i++ {
		// Random position within cookie bounds, keeping some margin from edges
		margin := chipSize
		xPos := xC - cookieW/2 + margin + rnd.Float32()*(cookieW-2*margin)
		yPos := yC - cookieH/2 + margin + rnd.Float32()*(cookieH-2*margin)

		chip := canvas.NewCircle(chipColor)
		chip.Resize(fyne.NewSize(chipSize, chipSize))
		chip.Move(fyne.NewPos(xPos, yPos))
		container.Add(chip)
	}

	// Warm background color
	rect.FillColor = color.NRGBA{R: 240, G: 220, B: 180, A: 255}

	container.Refresh()
}
