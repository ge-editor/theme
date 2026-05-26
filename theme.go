// Package theme provides color themes for the ge text editor (tcell/v3 compatible)
package theme

import "github.com/gdamore/tcell/v3"

// --- 特殊文字 ---
const (
	MarkTab       = '»'
	MarkNewline   = '¬'
	LF_WIDTH      = 1
	MarkEOF       = '~'
	MarkEOF_WIDTH = 1
	MarkContinue  = '⁃'
)

// --- 基本スタイル ---
var (
	// ColorDefault = tcell.StyleDefault.Foreground(tcell.ColorLightGrey) // 通常文字色
	ColorDefault = tcell.StyleDefault.Foreground(tcell.NewRGBColor(192, 192, 192)).Background(tcell.NewRGBColor(24, 24, 24))
)

// --- モードライン ---
var (
	// ColorModeLineActive = ColorDefault.Reverse(true) // Normal color
	// ColorModeLineActive = tcell.StyleDefault.Foreground(tcell.NewRGBColor(255, 255, 255)).Background(tcell.NewRGBColor(42, 123, 200)) // Blue line
	// ColorModeLineActive   = tcell.StyleDefault.Foreground(tcell.NewRGBColor(255, 255, 255)).Background(tcell.NewRGBColor(83, 94, 75)) // Green line
	ColorModeLineActive   = tcell.StyleDefault.Foreground(tcell.NewRGBColor(0, 0, 0)).Background(tcell.NewRGBColor(10, 173, 169)) // Green line
	ColorModelineInactive = ColorDefault.Foreground(tcell.ColorLightSlateGray).Reverse(true)
	ColorRightbar         = ColorDefault.Foreground(tcell.NewRGBColor(128, 128, 128)).Background(tcell.NewRGBColor(28, 28, 28))
	// 150,229,237 // cyan
	// 150,237,156 // light green
	// 216,184,138 //
)

// Linenumber
var (
	ColorLinenumber = ColorDefault.Foreground(tcell.NewRGBColor(128, 128, 128)).Background(tcell.NewRGBColor(32, 32, 32))
)

var (
	ColorEchoLine = ColorDefault
	// ColorEchoLine = ColorDefault.Background(tcell.NewRGBColor(51, 51, 51)).Foreground(tcell.ColorNames["white"])
)

// --- ポップアップメニュー ---
var (
	ColorPopupmenuForeground = ColorDefault.Foreground(tcell.ColorAntiqueWhite).Background(tcell.ColorDarkCyan)
	ColorPopupmenuBackground = ColorDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorLightSlateGrey)
)

// --- 特殊文字・空白・制御文字 ---
var (
	ColorSpecialChar  = ColorDefault.Foreground(tcell.ColorLightSlateGray)
	ColorTab          = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
	ColorSpace        = ColorDefault.Background(tcell.ColorDarkSlateGrey)
	ColorMarkContinue = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
	ColorMarkNewline  = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
	ColorMarkEOF      = ColorDefault.Foreground(tcell.ColorBlueViolet)
	ColorControlCode  = ColorDefault.Foreground(tcell.ColorRed)
)

// --- 検索関連 ---
var (
	ColorFind                = ColorDefault.Foreground(tcell.ColorRed)
	ColorSearchFound         = ColorDefault.Background(tcell.ColorDarkGreen)
	ColorSearchFoundOnCursor = ColorDefault.Background(tcell.ColorOrangeRed)
)

// --- ノード種類ごとの色 (Syntax Highlighting) ---
var CodeColors = map[string]tcell.Style{
	"interpreted_string_literal": ColorDefault.Foreground(tcell.ColorYellow),
	"comment":                    ColorDefault.Foreground(tcell.ColorBlue),
	"url":                        ColorDefault.Foreground(tcell.ColorLightCyan),
	"package":                    ColorDefault.Foreground(tcell.ColorGreen),
	"identifier":                 ColorDefault.Foreground(tcell.ColorRed),
	"string":                     ColorDefault.Foreground(tcell.ColorPurple),
	"number":                     ColorDefault.Foreground(tcell.ColorWhite),
	"int_literal":                ColorDefault.Foreground(tcell.ColorBrown),
	"slice_type":                 ColorDefault.Foreground(tcell.ColorDarkBlue),
	"default":                    ColorDefault,
}

// --- ヘルパー関数 (将来的に RGB カラー追加や動的テーマ変更に便利) ---
func Style(fg, bg tcell.Color, reverse, underline, bold bool) tcell.Style {
	s := ColorDefault.Foreground(fg).Background(bg)
	if reverse {
		s = s.Reverse(true)
	}
	if underline {
		s = s.Underline(true)
	}
	if bold {
		s = s.Bold(true)
	}
	return s
}
