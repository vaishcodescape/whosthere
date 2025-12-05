package views

import (
	"github.com/derailed/tview"
	"github.com/ramonvermeulen/whosthere/internal/state"
	"github.com/ramonvermeulen/whosthere/internal/ui/components"
	"github.com/ramonvermeulen/whosthere/internal/ui/table"
)

// MainPage is the dashboard showing discovered devices.
type MainPage struct {
	root        *tview.Flex
	deviceTable *table.DeviceTable
	spinner     *components.Spinner
	state       *state.AppState
}

func NewMainPage(s *state.AppState) *MainPage {
	t := table.NewDeviceTable()
	spinner := components.NewSpinner()

	mp := &MainPage{
		root:        nil,
		deviceTable: t,
		spinner:     spinner,
		state:       s,
	}

	main := tview.NewFlex().SetDirection(tview.FlexRow)
	main.AddItem(tview.NewTextView().SetText("whosthere").SetTextAlign(tview.AlignCenter), 0, 1, false)
	main.AddItem(t, 0, 18, true)

	status := tview.NewFlex().SetDirection(tview.FlexColumn)
	status.AddItem(mp.spinner.View(), 0, 1, false)
	status.AddItem(tview.NewTextView().SetText("jK up/down - gG top/bottom").SetTextAlign(tview.AlignRight), 0, 1, false)
	main.AddItem(status, 1, 0, false)

	mp.root = main
	return mp
}

func (p *MainPage) GetName() string { return "main" }

func (p *MainPage) GetPrimitive() tview.Primitive { return p.root }

func (p *MainPage) Spinner() *components.Spinner { return p.spinner }

// RefreshFromState reloads the device table from the shared AppState.
func (p *MainPage) RefreshFromState() {
	devices := p.state.DevicesSnapshot()
	p.deviceTable.ReplaceAll(devices)
}

// Refresh forces a redraw of the devices table.
func (p *MainPage) Refresh() { p.RefreshFromState() }
