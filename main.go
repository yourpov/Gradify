package gradify

import (
	"fmt"
	"strings"
)

// ── Preset Color Palettes ─────────────────────────────

var (
	// pre-made color palettes
	Candy = []string{"ff6666", "ffcc99", "ff99cc", "cc99ff", "99ccff", "66ccff", "ff99cc", "ff6666"}
	Minty = []string{"66ffcc", "ccff99", "99ffcc", "99ccff", "cc99ff", "ff99cc", "66ffcc"}
	// Common use cases — clear, professional colors
	Error   = []string{"cc0000", "e60000", "ff3333", "ff6666", "ff9999"}
	Success = []string{"006600", "009900", "33cc33", "66ff66", "99ff99"}
	Warning = []string{"cc9900", "e6b800", "ffcc00", "ffdb4d", "fff2b3"}
	Info    = []string{"0055cc", "0077ff", "3399ff", "66b2ff", "99ccff"}
)

// ── Types ─────────────────────────────
type Preset struct {
	Name        string
	Description string
	Hex         []string
}

type Color struct {
	R, G, B int
}

// ── Convert Hex String to Color ─────────────────────────────
// parses a hex color string to a color
func Convert(h string) (Color, error) {
	var c Color
	switch len(h) {
	case 6:
		_, err := fmt.Sscanf(h, "%02x%02x%02x", &c.R, &c.G, &c.B)
		if err != nil {
			return c, err
		}
	case 3:
		_, err := fmt.Sscanf(h, "%1x%1x%1x", &c.R, &c.G, &c.B)
		if err != nil {
			return c, err
		}
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		return c, fmt.Errorf("invalid hex color: %s", h)
	}
	return c, nil
}

// ── Colorize Text ─────────────────────────────
// adds ANSI 24-bit foreground color codes to the input
func Colorize(text string, r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, text)
}

// ── Gradient Function ─────────────────────────────
// applies a smooth RGB gradient to each character
func Gradient(text string, rgb []string) string {
	text = strings.TrimSpace(text)
	if len(text) == 0 || len(rgb) == 0 {
		return text
	}

	colors := make([]Color, 0, len(rgb))
	for _, hex := range rgb {
		c, err := Convert(hex)
		if err == nil {
			colors = append(colors, c)
		}
	}
	if len(colors) == 0 {
		return text
	}

	n := len(text)
	r := make([]int, n)
	g := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		segmentCount := float64(len(colors) - 1)
		pos := float64(i) / float64(n-1) * segmentCount
		idx1 := int(pos)
		idx2 := idx1 + 1
		if idx2 >= len(colors) {
			idx2 = len(colors) - 1
		}

		fraction := pos - float64(idx1)
		r[i] = int(float64(colors[idx1].R)*(1-fraction) + float64(colors[idx2].R)*fraction)
		g[i] = int(float64(colors[idx1].G)*(1-fraction) + float64(colors[idx2].G)*fraction)
		b[i] = int(float64(colors[idx1].B)*(1-fraction) + float64(colors[idx2].B)*fraction)
	}

	var builder strings.Builder
	for i, ch := range text {
		builder.WriteString(Colorize(string(ch), r[i], g[i], b[i]))
	}

	return builder.String()
}
