package config

import (
	"errors"
	"strings"
	"time"
)

const (
	DefaultSplashEnabled = true
	DefaultSplashDelay   = 1 * time.Second

	DefaultScanInterval = 20 * time.Second
	DefaultScanDuration = 10 * time.Second
)

// ThemeConfig mirrors tview.Theme; all fields are optional strings that
// override the tview.Theme defaults when set.
type ThemeConfig struct {
	PrimitiveBackgroundColor    string `yaml:"primitive_background_color"`
	ContrastBackgroundColor     string `yaml:"contrast_background_color"`
	MoreContrastBackgroundColor string `yaml:"more_contrast_background_color"`
	BorderColor                 string `yaml:"border_color"`
	TitleColor                  string `yaml:"title_color"`
	GraphicsColor               string `yaml:"graphics_color"`
	PrimaryTextColor            string `yaml:"primary_text_color"`
	SecondaryTextColor          string `yaml:"secondary_text_color"`
	TertiaryTextColor           string `yaml:"tertiary_text_color"`
	InverseTextColor            string `yaml:"inverse_text_color"`
	ContrastSecondaryTextColor  string `yaml:"contrast_secondary_text_color"`
}

// Config captures runtime configuration values loaded from the YAML config file.
type Config struct {
	ScanInterval time.Duration `yaml:"scan_interval"`
	ScanDuration time.Duration `yaml:"scan_duration"`
	Splash       SplashConfig  `yaml:"splash"`
	Theme        ThemeConfig   `yaml:"theme"`
	Scanners     ScannerConfig `yaml:"scanners"`
}

// SplashConfig controls the splash screen visibility and timing.
type SplashConfig struct {
	Enabled bool          `yaml:"enabled"`
	Delay   time.Duration `yaml:"delay"`
}

// ScannerConfig groups scanner enablement flags.
type ScannerConfig struct {
	MDNS ScannerToggle `yaml:"mdns"`
	SSDP ScannerToggle `yaml:"ssdp"`
	ARP  ScannerToggle `yaml:"arp"`
}

// ScannerToggle lets users enable/disable a scanner.
type ScannerToggle struct {
	Enabled bool `yaml:"enabled"`
}

// DefaultConfig builds a Config pre-populated with baked-in defaults.
func DefaultConfig() *Config {
	return &Config{
		ScanInterval: DefaultScanInterval,
		ScanDuration: DefaultScanDuration,
		Splash: SplashConfig{
			Enabled: DefaultSplashEnabled,
			Delay:   DefaultSplashDelay,
		},
		Theme: ThemeConfig{
			PrimitiveBackgroundColor:    "#000a1a",
			ContrastBackgroundColor:     "#001a33",
			MoreContrastBackgroundColor: "#003366",
			BorderColor:                 "#0088ff",
			TitleColor:                  "#00ffff",
			GraphicsColor:               "#00ffaa",
			PrimaryTextColor:            "#cceeff",
			SecondaryTextColor:          "#6699ff",
			TertiaryTextColor:           "#ffaa00",
			InverseTextColor:            "#000a1a",
			ContrastSecondaryTextColor:  "#88ddff",
		},
		Scanners: ScannerConfig{
			MDNS: ScannerToggle{Enabled: true},
			SSDP: ScannerToggle{Enabled: true},
			ARP:  ScannerToggle{Enabled: true},
		},
	}
}

// validateAndNormalize validates the config and fixes up out-of-range values.
func (c *Config) validateAndNormalize() error {
	var errs []string

	if c.Splash.Delay < 0 {
		errs = append(errs, "splash.delay must be >= 0")
		c.Splash.Delay = DefaultSplashDelay
	}

	if c.ScanInterval <= 0 {
		errs = append(errs, "scan_interval must be > 0")
		c.ScanInterval = DefaultScanInterval
	}

	if c.ScanDuration <= 0 {
		errs = append(errs, "scan_duration must be > 0")
		c.ScanDuration = DefaultScanDuration
	}

	if c.ScanDuration > c.ScanInterval {
		errs = append(errs, "scan_duration must be <= scan_interval")
		c.ScanDuration = c.ScanInterval
	}

	if !c.Scanners.MDNS.Enabled && !c.Scanners.SSDP.Enabled && !c.Scanners.ARP.Enabled {
		errs = append(errs, "at least one scanner must be enabled")
		c.Scanners.MDNS.Enabled = true
		c.Scanners.SSDP.Enabled = true
		c.Scanners.ARP.Enabled = true
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}

	return nil
}
