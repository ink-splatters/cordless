package main

import (
	"fmt"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/ink-splatters/cordless/tview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content tview.Primitive) {
	textView := tview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/ink-splatters/cordless/tview"
	fmt.Fprint(textView, url)
	return "End", Center(len(url), 1, textView)
}
