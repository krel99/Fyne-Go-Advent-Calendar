package days

import "fyne.io/fyne/v2"

type DayFunction func(rectangles []fyne.CanvasObject)

var Registry = map[int]DayFunction{
	1: Day01,
	2: Day02,
	3: Day03,
	4: Day04,
}
