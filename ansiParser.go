package gotel

// ForeColor Text foreground color
const ForeColor int = 30
// BackColor Text background color
const BackColor int = 40

// Color Text color
type Color uint8
const (
    // Black Black
    Black Color = 0
    // Red Red
    Red Color = 1
    // Green Green
    Green Color = 2
    // Yellow Yellow
    Yellow Color = 3
    // Blue Blue
    Blue Color = 4
    // Magenta Megenta
    Magenta Color = 5
    // Cyan Cyan
    Cyan Color = 6
    // White White
    White Color = 7
)

// TextFormat Text format
type TextFormat uint8
const (
    // Normal Normal
    Normal TextFormat = 0
    // Bold Blod
    Bold TextFormat = 1
    // Faint Faint
    Faint TextFormat = 2
    // Italic Italic
    Italic TextFormat = 3
    // Underline Underline
    Underline TextFormat = 4
    // Blink Blink
    Blink TextFormat = 5
    // Reverse Reverse text color
    Reverse TextFormat = 7
    // NonDisplay NonDisplay
    NonDisplay TextFormat = 8
)

// AnsiProperty Ansi text property for record ansi escape codes
type AnsiProperty struct {
    IsBlod bool
    IsFaint bool
    IsUnderline bool
    IsBlink bool
    IsReverse bool
    IsNonDisplay bool
    TextForeColor uint8
    TextBackColor uint8
}

// AnsiState ANSI parser state
type AnsiState uint8
const (
    // StateText In text state
    StateText AnsiState = 0
    // StateESC In ESC state
    StateESC AnsiState = 1
    // StateC1 In C1 state
    StateC1 AnsiState = 2
    // StateCSI In CSI state
    StateCSI AnsiState = 3
)

// AnsiCommand ANSI Command
type AnsiCommand uint8
const (
    
)

// AnsiParser comment
type AnsiParser struct {
    state AnsiState
    property AnsiProperty
}

// NewAnsiParser Create new ANSI parser
func NewAnsiParser() *AnsiParser {
    a := AnsiParser {
        StateText,
        AnsiProperty{},
    }
    return &a
}

// AddByte Add a byte
func (a *AnsiParser) AddByte(b byte) (AnsiProperty, AnsiCommand) {
    
    switch a.state {
    case StateText:
        switch b {
        case 27: // esc
            a.state = StateESC
        default:
            
        }
    }
    
    
}



