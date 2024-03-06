package tutorials

import (
	"NKNshellgo/config"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func shellScreen(w fyne.Window) fyne.CanvasObject {
	var seed string
	result := widget.NewLabel("")
	cmdEntry := widget.NewEntry()
	cmdsubmit := widget.NewButton("Execute", func() {
		result.SetText(cmdEntry.Text)
	})
	seedDisplay := widget.NewLabel("")
	seedDisplay.Wrapping = fyne.TextWrapWord
	cmdSend := container.NewVBox(cmdEntry, cmdsubmit)
	newSeed := dialog.NewConfirm("NewSeed", "Do you want to new a Seed", func(b bool) {
		if b {
			seed = "a1d87fdj0334kjd7323dsff3sdf9jdsk73jd672jdslf9o3"
			seedDisplay.SetText("Seed: " + seed)
			seedDisplay.Refresh()
		} else {
			seed = config.SeedId
			seedDisplay.SetText("Seed: " + seed)
			seedDisplay.Refresh()
		}
	}, w)
	newSeed.Show()
	return container.NewVBox(seedDisplay, widget.NewSeparator(), container.NewGridWithColumns(2, result, cmdSend))
}
