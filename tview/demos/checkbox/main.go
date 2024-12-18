// Demo code for the Checkbox primitive.
package main

import "github.com/ink-splatters/cordless/tview"

func main() {
	app := tview.NewApplication()
	checkbox := tview.NewCheckbox().SetLabel("Hit Enter to check box: ")
	if err := app.SetRoot(checkbox, true).Run(); err != nil {
		panic(err)
	}
}
