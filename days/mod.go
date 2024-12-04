package days

import "fyne.io/fyne/v2"

// Define the function type first
type DayFunction func(rectangles []fyne.CanvasObject)

var Registry = map[int]DayFunction{
	1: Day01, // Make sure this matches your function name exactly
}
