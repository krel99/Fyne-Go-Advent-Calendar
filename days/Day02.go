package days

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Day02(rectangles []fyne.CanvasObject) {
	if len(rectangles) <= 1 {
		return
	}

	container, ok := rectangles[1].(*fyne.Container)
	if !ok {
		return
	}

	rect, ok := container.Objects[0].(*canvas.Rectangle)
	if !ok {
		return
	}

	size := rect.Size()

	xC := size.Width / 2
	yC := size.Height/2 + size.Height*0.2

	// BODY PARTS
	unolegW := size.Width / 10 * 4
	unolegH := size.Height / 10 * 4

	bodyW := size.Width / 10 * 2.8
	bodyH := size.Height / 10 * 2.8

	headW := size.Width / 10 * 1.9
	headH := size.Height / 10 * 1.9

	snowColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255} // Pure white for snowman
	unoleg := canvas.NewCircle(snowColor)
	body := canvas.NewCircle(snowColor)
	head := canvas.NewCircle(snowColor)

	unoleg.Move(fyne.NewPos(xC-unolegW/2, yC-unolegH/2))
	unoleg.Resize(fyne.NewSize(unolegW, unolegH))

	body.Move(fyne.NewPos(xC-bodyW/2, yC-unolegH))
	body.Resize(fyne.NewSize(bodyW, bodyH))

	head.Move(fyne.NewPos(xC-headW/2, yC-unolegH-bodyH/2))
	head.Resize(fyne.NewSize(headW, headH))

	container.Add(unoleg)
	container.Add(body)
	container.Add(head)

	// DECORATIONS
	// ↪ eyes (a little bit crooked, because I don't adjust for their size in positioning)
	leftEye := canvas.NewCircle(color.Black)
	rightEye := canvas.NewCircle(color.Black)

	leftEye.Move(fyne.NewPos(xC-headW/7, yC-unolegH-bodyH/2.9))
	leftEye.Resize(fyne.NewSize(headW/7, headH/7))

	rightEye.Move(fyne.NewPos(xC+headW/7, yC-unolegH-bodyH/2.9))
	rightEye.Resize(fyne.NewSize(headW/7, headH/7))

	container.Add(leftEye)
	container.Add(rightEye)

	// ↪ shirt buttons
	bodyTop := yC - unolegH
	for i := 0; i < 4; i++ {
		shirtButt := canvas.NewCircle(color.Black)
		shirtButt.Resize(fyne.NewSize(headW/8, headH/8))
		buttonY := bodyTop + (bodyH * float32(i+1) / 4)
		shirtButt.Move(fyne.NewPos(xC-headW/16, buttonY))
		container.Add(shirtButt)
	}

	// Dark blue background for night sky
	rect.FillColor = color.NRGBA{R: 25, G: 25, B: 112, A: 255}

	container.Refresh()
}
