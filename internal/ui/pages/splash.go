package pages

import (
	"fmt"
	"strings"

	"github.com/derailed/tview"
)

var LogoBig = []string{
	`Knock Knock..                                                     `,
	`                _               _   _                   ___       `,
	`      __      _| |__   ___  ___| |_| |__   ___ _ __ ___/ _ \      `,
	`      \ \ /\ / / '_ \ / _ \/ __| __| '_ \ / _ \ '__/ _ \// /      `,
	`       \ V  V /| | | | (_) \__ \ |_| | | |  __/ | |  __/ \/       `,
	`        \_/\_/ |_| |_|\___/|___/\__|_| |_|\___|_|  \___| ()       `,
	"\n",
	"\n",
	"\n",
}

// SplashPage adapts the splash logo into a Page.
type SplashPage struct {
	root *tview.Flex
}

func NewSplashPage() *SplashPage {
	s := &SplashPage{root: tview.NewFlex().SetDirection(tview.FlexRow)}

	logo := tview.NewTextView()
	logo.SetDynamicColors(true)
	logo.SetTextAlign(tview.AlignCenter)
	logo.SetTextColor(tview.Styles.SecondaryTextColor)

	logoText := strings.Join(LogoBig, "\n")
	_, err := fmt.Fprint(logo, logoText)
	if err != nil {
		return nil
	}

	logoLines := len(strings.Split(logoText, "\n"))

	centered := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(logo, logoLines, 0, false).
		AddItem(nil, 0, 1, false)

	s.root.AddItem(centered, 0, 1, false)

	return s
}

func (p *SplashPage) GetName() string { return "splash" }

func (p *SplashPage) GetPrimitive() tview.Primitive { return p.root }

func (p *SplashPage) FocusTarget() tview.Primitive { return p.root }
