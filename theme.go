package theme

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v3"
)

// --- Special characters ---
const (
	MarkTab       = '»'
	MarkNewline   = '¬'
	LF_WIDTH      = 1
	MarkEOF       = '~'
	MarkEOF_WIDTH = 1
	MarkContinue  = '⁃'
)

const (
	DefaultColorForegroundRed   = 212
	DefaultColorForegroundGreen = 212
	DefaultColorForegroundBlue  = 212

	DefaultColorBackgroundRed   = 30
	DefaultColorBackgroundGreen = 30
	DefaultColorBackgroundBlue  = 30
)

// Restore the editor's default foreground and background colors
// for IME preedit rendering.
func RestoreTerminalAttributes() {
	fmt.Fprintf(os.Stdout,
		"\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm",
		DefaultColorForegroundRed, DefaultColorForegroundGreen, DefaultColorForegroundBlue,
		DefaultColorBackgroundRed, DefaultColorBackgroundGreen, DefaultColorBackgroundBlue,
	)
}

var (
	// --- Basic styles ---
	ColorDefault = tcell.StyleDefault.
			Foreground(tcell.NewRGBColor(
			DefaultColorForegroundRed,
			DefaultColorForegroundGreen,
			DefaultColorForegroundBlue,
		)).
		Background(tcell.NewRGBColor(
			DefaultColorBackgroundRed,
			DefaultColorBackgroundGreen,
			DefaultColorBackgroundBlue,
		)).
		Underline(tcell.NewRGBColor(124, 124, 124)) // Cursor row

	ColorColumnLimitOverflowBackground = tcell.NewRGBColor(44, 10, 33)

	// --- Mode line ---
	ColorModeLineActive = tcell.StyleDefault.
				Foreground(tcell.NewRGBColor(0, 0, 0)).
				Background(tcell.NewRGBColor(10, 173, 169))

	ColorModelineInactive = ColorDefault.
				Foreground(tcell.NewRGBColor(119, 136, 153)).
				Reverse(true)

	// Split leaf bar
	ColorRightbar = ColorDefault.
			Foreground(tcell.NewRGBColor(128, 128, 128)).
			Background(tcell.NewRGBColor(28, 28, 28))

	// Line numbers
	ColorLineNumber = ColorDefault.
			Foreground(tcell.NewRGBColor(128, 128, 128)).
			Background(tcell.NewRGBColor(40, 40, 40))
		// Background(tcell.NewRGBColor(51, 51, 51))
		// Background(tcell.NewRGBColor(32, 32, 32))

	ColorLineNumberOnEvenPage = ColorDefault.
					Foreground(tcell.NewRGBColor(152, 168, 164)).
					Background(tcell.NewRGBColor(40, 40, 40))
		// Background(tcell.NewRGBColor(51, 51, 51))
		// Background(tcell.NewRGBColor(32, 32, 32))

	ColorEchoLine = ColorDefault.
			Foreground(tcell.NewRGBColor(168, 168, 86)).
			Background(tcell.NewRGBColor(40, 40, 40))

	// --- Popup menu ---
	ColorPopupmenuForeground = ColorDefault.
					Foreground(tcell.NewRGBColor(250, 235, 215)).
					Background(tcell.NewRGBColor(0, 139, 139))

	ColorPopupmenuBackground = ColorDefault.
					Foreground(tcell.NewRGBColor(0, 0, 0)).
					Background(tcell.NewRGBColor(119, 136, 153))

	// --- Special characters, whitespace, and control characters ---
	ColorTab          = ColorDefault.Foreground(tcell.NewRGBColor(47, 79, 79))
	ColorSpace        = ColorDefault.Background(tcell.NewRGBColor(47, 79, 79))
	ColorMarkContinue = ColorDefault.Foreground(tcell.NewRGBColor(47, 79, 79))
	ColorMarkNewline  = ColorDefault.Foreground(tcell.NewRGBColor(47, 79, 79))
	ColorMarkEOF      = ColorDefault.Foreground(tcell.NewRGBColor(138, 43, 226))
	ColorControlCode  = ColorDefault.Foreground(tcell.NewRGBColor(255, 0, 0))

	// --- Search ---
	ColorFind                = ColorDefault.Foreground(tcell.NewRGBColor(255, 0, 0))
	ColorSearchFound         = ColorDefault.Background(tcell.NewRGBColor(0, 100, 0))
	ColorSearchFoundOnCursor = ColorDefault.Background(tcell.NewRGBColor(255, 69, 0))
)

// --- Colors by node type (syntax highlighting) ---
var CodeColors = map[string]tcell.Style{
	// Go
	"interpreted_string_literal": ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)), // String: #ce9178
	"comment":                    ColorDefault.Foreground(tcell.NewRGBColor(106, 153, 85)),  // Comment: #6a9955
	"url":                        ColorDefault.Foreground(tcell.NewRGBColor(78, 201, 176)),  // URL / type: #4ec9b0
	// "identifier":                 ColorDefault.Foreground(tcell.NewRGBColor(156, 220, 254)), // Identifier / variable: #9cdcfe
	"string": ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)), // String: #ce9178
	// "number":                     ColorDefault.Foreground(tcell.NewRGBColor(181, 206, 168)), // Number: #b5cea8
	// "int_literal":                ColorDefault.Foreground(tcell.NewRGBColor(181, 206, 168)), // Integer literal: #b5cea8
	// "slice_type":                 ColorDefault.Foreground(tcell.NewRGBColor(78, 201, 176)),  // Type: #4ec9b0

	// Go言語（Go言語仕様で定義されている25個の予約語
	"var":         ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"const":       ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"func":        ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"type":        ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"package":     ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"import":      ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"if":          ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"else":        ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"switch":      ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"case":        ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"default":     ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"fallthrough": ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"for":         ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"range":       ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"break":       ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"continue":    ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"goto":        ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"return":      ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"select":      ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"struct":      ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"interface":   ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"map":         ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"chan":        ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"go":          ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6
	"defer":       ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)), // #569CD6

	// その他 追加 キーワード
	// "keyword":          ColorDefault.Foreground(tcell.NewRGBColor(197, 134, 192)), // Keyword: #c586c0
	// "function":         ColorDefault.Foreground(tcell.NewRGBColor(220, 220, 170)), // Function: #dcdcaa
	// "boolean":          ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)),  // Boolean: #569cd6
	// "field_identifier": ColorDefault.Foreground(tcell.NewRGBColor(156, 220, 254)), // Field: #9cdcfe
	// "operator":         ColorDefault.Foreground(tcell.NewRGBColor(212, 212, 212)), // Operator: #d4d4d4

	// HUNG UP
	// "escape_sequence":  ColorDefault.Foreground(tcell.NewRGBColor(215, 186, 125)), // Escape: #d7ba7d

	// Markdown
	"atx_heading":       ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)).Bold(true), // String: #569CD6
	"setext_heading":    ColorDefault.Foreground(tcell.NewRGBColor(86, 156, 214)).Bold(true), // String: #569CD6
	"block_quote":       ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"list_marker":       ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"fenced_code_block": ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"code_span":         ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"emphasis":          ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"strong_emphasis":   ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"strikethrough":     ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"link":              ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"image":             ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"link_destination":  ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"link_title":        ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178
	"html_tag":          ColorDefault.Foreground(tcell.NewRGBColor(206, 145, 120)),           // String: #ce9178

	// "default": ColorDefault,
}
