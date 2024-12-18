package tviewutil

import "github.com/ink-splatters/cordless/tview"

func CreateCenteredComponent(component tview.Primitive, width int) tview.Primitive {
	padding := tview.NewFlex().SetDirection(tview.FlexColumn)
	padding.AddItem(tview.NewBox(), 0, 1, false)
	padding.AddItem(component, width, 0, false)
	padding.AddItem(tview.NewBox(), 0, 1, false)

	return padding
}
