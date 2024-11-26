package tui

import "github.com/charmbracelet/lipgloss"

type StylesType = map[string]lipgloss.Style

const (
	themeColor = "#6731F1"
)

var Styles StylesType = StylesType{
	"themed": lipgloss.NewStyle().
		Foreground(lipgloss.Color(themeColor)),
	"muted": lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#aaaaaa", Dark: "#555555"}),
	"emphasized": lipgloss.NewStyle().
		Bold(true),
	"cell-selected": lipgloss.NewStyle().
		Background(lipgloss.Color("#FFFFFF")),
	"cell-flagged": lipgloss.NewStyle().
		Background(lipgloss.Color("#6731F1")),
	"tab-active": lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")),
	"tab-inactive": lipgloss.NewStyle().
		Foreground(lipgloss.Color("#555555")),
}
