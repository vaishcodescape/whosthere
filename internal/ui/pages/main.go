package pages

import (
	"github.com/derailed/tview"
	"github.com/ramonvermeulen/whosthere/internal/state"
	"github.com/ramonvermeulen/whosthere/internal/ui/components"
)

// DashboardPage is the dashboard showing discovered devices.
type DashboardPage struct {
	root        *tview.Flex
	deviceTable *components.DeviceTable
	spinner     *components.Spinner
	state       *state.AppState

	onShowDetails func()
}

func NewDashboardPage(s *state.AppState, onShowDetails func()) *DashboardPage {
	t := components.NewDeviceTable()
	spinner := components.NewSpinner()

	main := tview.NewFlex().SetDirection(tview.FlexRow)
	main.AddItem(
		tview.NewTextView().
			SetText("whosthere").
			SetTextAlign(tview.AlignCenter),
		0, 1, false,
	)
	main.AddItem(t, 0, 18, true)

	status := tview.NewFlex().SetDirection(tview.FlexColumn)
	status.AddItem(spinner.View(), 0, 1, false)
	status.AddItem(
		tview.NewTextView().
			SetText("j/k: up/down  g/G: top/bottom  Enter: details").
			SetTextAlign(tview.AlignRight),
		0, 1, false,
	)
	main.AddItem(status, 1, 0, false)

	dp := &DashboardPage{
		root:          main,
		deviceTable:   t,
		spinner:       spinner,
		state:         s,
		onShowDetails: onShowDetails,
	}

	t.SetSelectedFunc(func(row, col int) {
		ip := t.SelectedIP()
		if ip == "" || onShowDetails == nil {
			return
		}
		s.SetSelectedIP(ip)
		onShowDetails()
	})

	return dp
}

func (p *DashboardPage) GetName() string { return "main" }

func (p *DashboardPage) GetPrimitive() tview.Primitive { return p.root }

func (p *DashboardPage) FocusTarget() tview.Primitive { return p.deviceTable }

func (p *DashboardPage) Spinner() *components.Spinner { return p.spinner }

func (p *DashboardPage) RefreshFromState() {
	devices := p.state.DevicesSnapshot()
	p.deviceTable.ReplaceAll(devices)
}
