package generics

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// HandleImage loads and displays an image in a specified container
func HandleImage(rectangles []fyne.CanvasObject, dayIndex int, imagePath string) {
	if len(rectangles) <= dayIndex {
		log.Printf("Error: Not enough rectangles provided. Required: %d, Got: %d", dayIndex+1, len(rectangles))
		return
	}

	container, ok := rectangles[dayIndex].(*fyne.Container)
	if !ok {
		log.Printf("Error: Failed to cast rectangle at index %d to container", dayIndex)
		return
	}

	fyneImage := canvas.NewImageFromFile(imagePath)
	fyneImage.FillMode = canvas.ImageFillContain
	fyneImage.Resize(container.Size())

	container.Add(fyneImage)

	container.Refresh()

	log.Printf("Image '%s' successfully loaded and displayed in container %d.", imagePath, dayIndex)
}
