package theme

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v3"
)

// --- 特殊文字 ---
const (
	MarkTab       = '»'
	MarkNewline   = '¬'
	LF_WIDTH      = 1
	MarkEOF       = '~'
	MarkEOF_WIDTH = 1
	MarkContinue  = '⁃'
)

const (
	DefaultColorForgroundRed   = 192
	DefaultColorForgroundGreen = 192
	DefaultColorForgroundBlue  = 192

	DefaultColorBackgroundRed   = 24
	DefaultColorBackgroundGreen = 24
	DefaultColorBackgroundBlue  = 24
)

// Restore the editor's default foreground and background colors
// for IME preedit rendering.
func RestoreTerminalAttributes() {
	fmt.Fprintf(os.Stdout,
		"\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm",
		DefaultColorForgroundRed, DefaultColorForgroundGreen, DefaultColorForgroundBlue,
		DefaultColorBackgroundRed, DefaultColorBackgroundGreen, DefaultColorBackgroundBlue,
	)
}

var (
	// --- 基本スタイル ---
	ColorDefault = tcell.StyleDefault.
			Foreground(tcell.NewRGBColor(DefaultColorForgroundRed, DefaultColorForgroundGreen, DefaultColorForgroundBlue)).
			Background(tcell.NewRGBColor(DefaultColorBackgroundRed, DefaultColorBackgroundGreen, DefaultColorBackgroundBlue)).
			Underline(tcell.NewRGBColor(124, 124, 124)) // Cursor row

	ColorColumnLimitOverflowBackground = tcell.NewRGBColor(44, 10, 33)

	// --- Mode line ---
	ColorModeLineActive   = tcell.StyleDefault.Foreground(tcell.NewRGBColor(0, 0, 0)).Background(tcell.NewRGBColor(10, 173, 169)) // Green line
	ColorModelineInactive = ColorDefault.Foreground(tcell.ColorLightSlateGray).Reverse(true)

	// Split leaf bar
	ColorRightbar = ColorDefault.Foreground(tcell.NewRGBColor(128, 128, 128)).Background(tcell.NewRGBColor(28, 28, 28))

	// Line Number
	ColorLineNumber           = ColorDefault.Foreground(tcell.NewRGBColor(128, 128, 128)).Background(tcell.NewRGBColor(32, 32, 32))
	ColorLineNumberOnEvenPage = ColorDefault.Foreground(tcell.NewRGBColor(152, 168, 164)).Background(tcell.NewRGBColor(32, 32, 32))

	ColorEchoLine = ColorDefault.Foreground(tcell.NewRGBColor(168, 168, 86)).Background(tcell.NewRGBColor(40, 40, 40))

	// --- ポップアップメニュー ---
	ColorPopupmenuForeground = ColorDefault.Foreground(tcell.ColorAntiqueWhite).Background(tcell.ColorDarkCyan)
	ColorPopupmenuBackground = ColorDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorLightSlateGrey)

	// --- 特殊文字・空白・制御文字 ---
	ColorTab          = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
	ColorSpace        = ColorDefault.Background(tcell.ColorDarkSlateGrey)
	ColorMarkContinue = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
	ColorMarkNewline  = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
	ColorMarkEOF      = ColorDefault.Foreground(tcell.ColorBlueViolet)
	ColorControlCode  = ColorDefault.Foreground(tcell.ColorRed)

	// --- 検索関連 ---
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
/*
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
*/
