// Package theme provides themes for the ge text editor.

package theme

import "github.com/gdamore/tcell/v3"

const (
	MarkTab       = '»'
	MarkNewline   = '¬'
	LF_WIDTH      = 1
	MarkEOF       = '~'
	MarkEOF_WIDTH = 1
	MarkContinue  = '⁃'
)

// Normal
// var ColorDefault = tcell.StyleDefault.Foreground(tcell.ColorNavajoWhite)
// var ColorDefault = tcell.StyleDefault.Foreground(tcell.ColorLightSeaGreen)
var ColorDefault = tcell.StyleDefault.Foreground(tcell.ColorLightGrey) // Might be good for Modeline
// var ColorDefault = tcell.StyleDefault.Foreground(tcell.ColorWhite)

// Modeline
var ColorModeLineActive = ColorDefault.Reverse(true)

// var ColorModeLineActive = tcell.StyleDefault.Foreground(tcell.ColorLightGrey).Reverse(true)        // Might be good for Modeline
var ColorModelineInactive = tcell.StyleDefault.Foreground(tcell.ColorLightSlateGray).Reverse(true) // Might be good for Modeline
// var ColorModeline = ColorDefault.Foreground(tcell.ColorDarkCyan).Reverse(true)

// var ColorRightbar = tcell.StyleDefault.Foreground(tcell.ColorLightSlateGray).Reverse(true) // Might be good for Modeline
var ColorRightbar = tcell.StyleDefault.Foreground(tcell.ColorDarkSlateGray).Reverse(true) // Might be good for Modeline
// var ColorRightbar = tcell.StyleDefault.Foreground(tcell.ColorGray).Reverse(true) // Might be good for Modeline
// var ColorRightbar = tcell.StyleDefault // Might be good for Modeline

// Green
//var ColorDefault = tcell.StyleDefault.Background(tcell.ColorBlack.TrueColor()).Foreground(tcell.ColorLightGreen)

// Gray
// var ColorDefault = theme.ColorDefault.Background(tcell.NewRGBColor(0, 0, 0)).Foreground(tcell.NewRGBColor(0, 255, 200))

// Popupmenu Normal
var ColorPopupmenuForeground = ColorDefault.Foreground(tcell.ColorAntiqueWhite).Background(tcell.ColorDarkCyan)

// var ColorPopupmenuBackground = ColorDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorAntiqueWhite)
var ColorPopupmenuBackground = ColorDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorLightSlateGrey)

// Popupmenu Green
// var ColorPopupmenuBackground = tcell.StyleDefault.Background(tcell.ColorCadetBlue).Foreground(tcell.ColorAntiqueWhite)// green
//var ColorPopupmenuBackground = tcell.StyleDefault.Background(tcell.Color17).Foreground(tcell.ColorAntiqueWhite)

// var ColorPopupmenuForeground = ColorPopupmenuBackground.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkCyan)//green
// var ColorPopupmenuForeground = ColorPopupmenuBackground.Foreground(tcell.ColorWhite).Background(tcell.Color45) // light blue
// var ColorPopupmenuForeground = ColorPopupmenuBackground.Foreground(tcell.ColorWhite).Background(tcell.Color45) // light blue
//var ColorPopupmenuForeground = ColorPopupmenuBackground.Foreground(tcell.ColorWhite).Background(tcell.Color26) //

// var ColorPopupmenuForeground = ColorPopupmenuBackground.Reverse(true).Foreground(tcell.ColorAntiqueWhite).Background(tcell.ColorDarkCyan)

// Special character

// var ColorSpecialChar = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
var ColorSpecialChar = ColorDefault.Foreground(tcell.ColorLightSlateGray)

// var ColorCurrentLine = ColorDefault.Underline(true)

var ColorTab = ColorDefault.Foreground(tcell.ColorDarkSlateGray)

// var ColorSpace = ColorDefault.Background(tcell.ColorDarkOliveGreen)
var ColorSpace = ColorDefault.Background(tcell.ColorDarkSlateGrey)
var ColorMarkContinue = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
var ColorMarkNewline = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
var ColorMarkEOF = ColorDefault.Foreground(tcell.ColorBlueViolet)
var ColorControlCode = ColorDefault.Foreground(tcell.ColorRed)
var ColorFind = ColorDefault.Foreground(tcell.ColorRed)
var ColorSearchFound = ColorDefault.Background(tcell.ColorDarkGreen)
var ColorSearchFoundOnCursor = ColorDefault.Background(tcell.ColorOrangeRed)

// ノードの種類に応じた色マッピング
var CodeColors = map[string]tcell.Style{
	"interpreted_string_literal": ColorDefault.Foreground(tcell.ColorYellow),    // "\033[1;33m", // 黄色
	"comment":                    ColorDefault.Foreground(tcell.ColorBlue),      // "\033[1;34m", // 青
	"url":                        ColorDefault.Foreground(tcell.ColorLightCyan), // "\033[1;36m", // シアン
	"package":                    ColorDefault.Foreground(tcell.ColorGreen),     // "\033[1;32m", // 緑
	"identifier":                 ColorDefault.Foreground(tcell.ColorRed),       // "\033[1;31m", // 赤
	"string":                     ColorDefault.Foreground(tcell.ColorPurple),    // "\033[1;35m", // 紫
	"number":                     ColorDefault.Foreground(tcell.ColorWhite),     // "\033[1;37m", // 白
	"int_literal":                ColorDefault.Foreground(tcell.ColorBrown),     // "\033[0;33m", // 茶色
	"slice_type":                 ColorDefault.Foreground(tcell.ColorDarkBlue),  // "\033[0;34m", // 濃い青
	"default":                    ColorDefault,                                  // "\033[0m",    // リセット
}
