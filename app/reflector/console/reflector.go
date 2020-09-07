package console

import (
	"fmt"
	"github.com/koushamad/Enigma/app/reflector/service"
	"github.com/koushamad/Enigma/infra/repository/file"

	ui "github.com/VladimirMarkelov/clui"
)

type Reflector struct {
	service    service.Reflect
	repository file.Reflector
}

func (r Reflector) Run() {
	ui.InitLibrary()
	defer ui.DeinitLibrary()
	r.createView()

	ui.MainLoop()
}

func (r Reflector) createView() {
	view := ui.AddWindow(1, 1, 80, 7, "Reflector Generator")
	view.SetTitleButtons(ui.ButtonMaximize | ui.ButtonClose)

	frmViews := ui.CreateFrame(view, 8, 5, ui.BorderNone, ui.Fixed)
	frmViews.SetPack(ui.Horizontal)

	frmAction := ui.CreateFrame(frmViews, 8, 5, ui.BorderThin, ui.Fixed)
	frmAction.SetPack(ui.Vertical)
	frmAction.SetTitle("Actions")

	frmReflectors := ui.CreateFrame(frmViews, 70, 5, ui.BorderThin, ui.Fixed)
	frmReflectors.SetPack(ui.Vertical)
	frmReflectors.SetTitle("Reflectors")

	bch := ui.CreateTextDisplay(frmReflectors, 45, 24, 1)
	ui.ActivateControl(frmReflectors, bch)

	btn1 := ui.CreateButton(frmAction, ui.AutoSize, 4, "Show", ui.Fixed)
	btn2 := ui.CreateButton(frmAction, ui.AutoSize, 4, "Create", ui.Fixed)
	btn3 := ui.CreateButton(frmAction, ui.AutoSize, 4, "Quit", ui.Fixed)

	btn1.SetShadowType(ui.ShadowNone)
	btn2.SetShadowType(ui.ShadowNone)
	btn3.SetShadowType(ui.ShadowNone)

	btn1.OnClick(func(event ui.Event) {
		bch := ui.CreateTextDisplay(frmReflectors, 45, 24, 1)
		ui.ActivateControl(frmReflectors, bch)

		bch.SetLineCount(len(r.repository.GetAll()))
		bch.OnDrawLine(func(ind int) string {
			if reflector, err := r.repository.GetByIndex(ind); err == nil {
				return fmt.Sprintf("%03d %s", ind, reflector.Chars)
			}
			return ""
		})
	})

	btn1.OnClick(func(event ui.Event) {
		r.loadReflectors(bch)
	})

	btn2.OnClick(func(event ui.Event) {
		r.service.Generate()
		r.loadReflectors(bch)
	})

	btn3.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})

	ui.ActivateControl(view, btn1)
}

func (r Reflector) loadReflectors(bch *ui.TextDisplay) {

	bch.SetLineCount(len(r.repository.GetAll()))
	bch.OnDrawLine(func(ind int) string {
		if reflector, err := r.repository.GetByIndex(ind); err == nil {
			return fmt.Sprintf("%03d %s", ind, reflector.Chars)
		}
		return ""
	})
}
