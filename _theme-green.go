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
// var ColorDefault = tcell.StyleDefault

// Green
var ColorDefault = tcell.StyleDefault.Background(tcell.ColorBlack.TrueColor()).Foreground(tcell.ColorLightGreen)

// Gray
// var ColorDefault = theme.ColorDefault.Background(tcell.NewRGBColor(0, 0, 0)).Foreground(tcell.NewRGBColor(0, 255, 200))

var ColorSpecialChar = ColorDefault.Foreground(tcell.ColorDarkSlateGray)

// var ColorModeline = ColorDefault.Foreground(tcell.ColorDarkCyan).Reverse(true)
var ColorModeline = ColorDefault.Reverse(true)

// var ColorCurrentLine = ColorDefault.Underline(true)

var ColorTab = ColorDefault.Foreground(tcell.ColorDarkSlateGray)

// var ColorSpace = ColorDefault.Background(tcell.ColorDarkOliveGreen)
var ColorSpace = ColorDefault.Background(tcell.ColorDarkSlateGrey)

// Popupmenu Normal
//var ColorPopupmenuForeground = ColorDefault.Foreground(tcell.ColorAntiqueWhite).Background(tcell.ColorDarkCyan)
//var ColorPopupmenuBackground = ColorDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorAntiqueWhite)

// Popupmenu Green
// var ColorPopupmenuBackground = tcell.StyleDefault.Background(tcell.ColorCadetBlue).Foreground(tcell.ColorAntiqueWhite)// green
var ColorPopupmenuBackground = tcell.StyleDefault.Background(tcell.Color17).Foreground(tcell.ColorAntiqueWhite)

// var ColorPopupmenuForeground = ColorPopupmenuBackground.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkCyan)//green
// var ColorPopupmenuForeground = ColorPopupmenuBackground.Foreground(tcell.ColorWhite).Background(tcell.Color45) // light blue
// var ColorPopupmenuForeground = ColorPopupmenuBackground.Foreground(tcell.ColorWhite).Background(tcell.Color45) // light blue
var ColorPopupmenuForeground = ColorPopupmenuBackground.Foreground(tcell.ColorWhite).Background(tcell.Color26) //

// var ColorPopupmenuForeground = ColorPopupmenuBackground.Reverse(true).Foreground(tcell.ColorAntiqueWhite).Background(tcell.ColorDarkCyan)

var ColorMarkContinue = ColorDefault.Foreground(tcell.ColorDarkSlateGray)
var ColorMarkLinefeed = ColorDefault.Foreground(tcell.ColorBlueViolet)
var ColorMarkEOF = ColorDefault.Foreground(tcell.ColorBlueViolet)
var ColorControlCode = ColorDefault.Foreground(tcell.ColorRed)
