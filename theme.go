// Package theme provides themes for the ge text editor.

package theme

import "github.com/gdamore/tcell/v2"

const (
	MarkTab       = '»'
	MarkLinefeed  = '¬'
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
var ColorMarkLinefeed = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
var ColorMarkEOF = ColorDefault.Foreground(tcell.ColorBlueViolet)
var ColorControlCode = ColorDefault.Foreground(tcell.ColorRed)
var ColorFind = ColorDefault.Foreground(tcell.ColorRed)
var ColorSearchFound = ColorDefault.Background(tcell.ColorDarkGreen)
var ColorSearchFoundOnCursor = ColorDefault.Background(tcell.ColorOrangeRed)
