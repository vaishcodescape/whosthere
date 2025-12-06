package theme

import (
	"strings"

	"github.com/derailed/tcell/v2"
	"github.com/derailed/tview"
	"github.com/ramonvermeulen/whosthere/internal/config"
)

// FromConfig starts from the current tview.Styles and applies
// any overrides from ThemeConfig. It returns the resulting tview.Theme.
func FromConfig(tc config.ThemeConfig) tview.Theme {
	th := tview.Styles

	if c := parseColor(tc.PrimitiveBackgroundColor); c != nil {
		th.PrimitiveBackgroundColor = *c
	}
	if c := parseColor(tc.ContrastBackgroundColor); c != nil {
		th.ContrastBackgroundColor = *c
	}
	if c := parseColor(tc.MoreContrastBackgroundColor); c != nil {
		th.MoreContrastBackgroundColor = *c
	}
	if c := parseColor(tc.BorderColor); c != nil {
		th.BorderColor = *c
	}
	if c := parseColor(tc.TitleColor); c != nil {
		th.TitleColor = *c
	}
	if c := parseColor(tc.GraphicsColor); c != nil {
		th.GraphicsColor = *c
	}
	if c := parseColor(tc.PrimaryTextColor); c != nil {
		th.PrimaryTextColor = *c
	}
	if c := parseColor(tc.SecondaryTextColor); c != nil {
		th.SecondaryTextColor = *c
	}
	if c := parseColor(tc.TertiaryTextColor); c != nil {
		th.TertiaryTextColor = *c
	}
	if c := parseColor(tc.InverseTextColor); c != nil {
		th.InverseTextColor = *c
	}
	if c := parseColor(tc.ContrastSecondaryTextColor); c != nil {
		th.ContrastSecondaryTextColor = *c
	}

	tview.Styles = th
	return th
}

// helper to transform user defined color strings into tcell.Color pointers.
// W3C color names and hex values are supported. Returns nil for empty strings.
func parseColor(s string) *tcell.Color {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	c := tcell.GetColor(s)
	return &c
}
