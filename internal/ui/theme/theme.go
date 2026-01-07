package theme

import (
	"sort"
	"strings"

	"github.com/derailed/tcell/v2"
	"github.com/derailed/tview"
	"github.com/ramonvermeulen/whosthere/internal/config"
	"github.com/ramonvermeulen/whosthere/internal/logging"
	"go.uber.org/zap"
)

var registry = map[string]tview.Theme{
	config.DefaultThemeName: {
		PrimitiveBackgroundColor:    tcell.GetColor("#000a1a"),
		ContrastBackgroundColor:     tcell.GetColor("#001a33"),
		MoreContrastBackgroundColor: tcell.GetColor("#003366"),
		BorderColor:                 tcell.GetColor("#0088ff"),
		TitleColor:                  tcell.GetColor("#00ffff"),
		GraphicsColor:               tcell.GetColor("#00ffaa"),
		PrimaryTextColor:            tcell.GetColor("#cceeff"),
		SecondaryTextColor:          tcell.GetColor("#6699ff"),
		TertiaryTextColor:           tcell.GetColor("#ffaa00"),
		InverseTextColor:            tcell.GetColor("#000a1a"),
		ContrastSecondaryTextColor:  tcell.GetColor("#88ddff"),
	},
	"dracula": {
		PrimitiveBackgroundColor:    tcell.GetColor("#282a36"),
		ContrastBackgroundColor:     tcell.GetColor("#343746"),
		MoreContrastBackgroundColor: tcell.GetColor("#44475a"),
		BorderColor:                 tcell.GetColor("#bd93f9"),
		TitleColor:                  tcell.GetColor("#f8f8f2"),
		GraphicsColor:               tcell.GetColor("#ff79c6"),
		PrimaryTextColor:            tcell.GetColor("#f8f8f2"),
		SecondaryTextColor:          tcell.GetColor("#8be9fd"),
		TertiaryTextColor:           tcell.GetColor("#50fa7b"),
		InverseTextColor:            tcell.GetColor("#282a36"),
		ContrastSecondaryTextColor:  tcell.GetColor("#ffb86c"),
	},
	"solarized-dark": {
		PrimitiveBackgroundColor:    tcell.GetColor("#002b36"),
		ContrastBackgroundColor:     tcell.GetColor("#073642"),
		MoreContrastBackgroundColor: tcell.GetColor("#586e75"),
		BorderColor:                 tcell.GetColor("#2aa198"),
		TitleColor:                  tcell.GetColor("#93a1a1"),
		GraphicsColor:               tcell.GetColor("#cb4b16"),
		PrimaryTextColor:            tcell.GetColor("#93a1a1"),
		SecondaryTextColor:          tcell.GetColor("#b58900"),
		TertiaryTextColor:           tcell.GetColor("#859900"),
		InverseTextColor:            tcell.GetColor("#002b36"),
		ContrastSecondaryTextColor:  tcell.GetColor("#268bd2"),
	},
	"gruvbox-dark": {
		PrimitiveBackgroundColor:    tcell.GetColor("#282828"),
		ContrastBackgroundColor:     tcell.GetColor("#3c3836"),
		MoreContrastBackgroundColor: tcell.GetColor("#504945"),
		BorderColor:                 tcell.GetColor("#fe8019"),
		TitleColor:                  tcell.GetColor("#ebdbb2"),
		GraphicsColor:               tcell.GetColor("#d3869b"),
		PrimaryTextColor:            tcell.GetColor("#ebdbb2"),
		SecondaryTextColor:          tcell.GetColor("#fabd2f"),
		TertiaryTextColor:           tcell.GetColor("#b8bb26"),
		InverseTextColor:            tcell.GetColor("#1d2021"),
		ContrastSecondaryTextColor:  tcell.GetColor("#83a598"),
	},
}

// Resolve returns the theme by name. Unknown names fall back to default; "custom" applies overrides atop default.
func Resolve(tc *config.ThemeConfig) tview.Theme {
	name := strings.ToLower(strings.TrimSpace(config.DefaultThemeName))
	if tc != nil {
		if n := strings.TrimSpace(tc.Name); n != "" {
			name = strings.ToLower(n)
		}
	}

	base, ok := registry[name]
	if name == config.CustomThemeName {
		defaultTheme := registry[config.DefaultThemeName]
		base = applyOverrides(&defaultTheme, tc)
	} else if !ok {
		logging.L().Warn("theme not found, falling back to default", zap.String("name", name))
		base = registry[config.DefaultThemeName]
	}

	tview.Styles = base
	return base
}

// applyOverrides starts from base and applies overrides from config.
func applyOverrides(base *tview.Theme, tc *config.ThemeConfig) tview.Theme {
	if base == nil {
		return registry[config.DefaultThemeName]
	}

	th := *base
	if tc == nil {
		return th
	}

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

	return th
}

// helper to transform user defined color strings into tcell.Color pointers.
func parseColor(s string) *tcell.Color {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	c := tcell.GetColor(s)
	return &c
}

// Register adds or replaces a theme in the registry.
func Register(name string, th *tview.Theme) {
	if th == nil {
		return
	}
	registry[strings.ToLower(strings.TrimSpace(name))] = *th
}

// Names returns the currently registered theme names (built-ins plus any custom registrations).
func Names() []string {
	names := make([]string, 0, len(registry))
	for k := range registry {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}
