package main

import (
	"advent/days"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Advent Calendar")

	// Create a slice to hold our rectangles
	var rectangles []fyne.CanvasObject
	var buttons []fyne.CanvasObject

	// Create empty grids
	canvasGrid := container.New(layout.NewGridWrapLayout(fyne.NewSize(125, 125)))
	canvasGrid.Resize(fyne.NewSize(800, 600))

	uiPositions := container.NewWithoutLayout()
	uiPositions.Resize(fyne.NewSize(800, 600))

	// Set window size

	// show window, before rendering starts
	window.Show()

	go func() {
		for i := 0; i < 24; i++ {
			slotContainer := container.NewWithoutLayout()
			rect := canvas.NewRectangle(color.White)
			rect.Resize(fyne.NewSize(125, 125))
			slotContainer.Add(rect)
			rectangles = append(rectangles, slotContainer)
			canvasGrid.Add(slotContainer)
			time.Sleep(45 * time.Millisecond) // 200ms delay
		}

		column := 0
		row := 0

		for row = 0; row < 4 && len(buttons) < 24; column++ {
			butt := widget.NewButton("Reveal", nil)

			x := column*125 + column*4 // 6 buttons per row, 125px width
			y := row*125 + row*4

			butt.Resize(fyne.NewSize(125, 125))
			butt.Move(fyne.NewPos(float32(x), float32(y)))

			currentDay := len(buttons) + 1

			butt.OnTapped = func() {
				// check, jestli funkce s danym indexem existuje
				if fn, exists := days.Registry[currentDay]; exists {
					fn(rectangles)
				}
				butt.Hide()
				butt.Disable()
			}

			buttons = append(buttons, butt)

			uiPositions.Add(butt)

			if column == 5 {
				column = -1
				row++
			}
		}
		uiPositions.Move(fyne.NewPos(0, 0))
	}()

	content := container.NewWithoutLayout(canvasGrid, uiPositions)

	window.SetContent(content)
	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}
