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

// TelWord Word property in TelWindow
type TelWord struct {
    IsBlod bool
    IsFaint bool
    IsUnderline bool
    IsBlink bool
    
    IsReverse bool
    IsNonDisplay bool
    TextForeColor uint8
    TextBackColor uint8
    
    TextWidth uint8
    TextIndex uint8
    
    TextCode [4]byte
}

// NewTelWindow Create new TelWindow
func NewTelWindow(c int, r int) *TelWindow {
    t := TelWindow {
        c, // cols
        r, // rows
        0, // point col
        0, // point row
        make([]TelWord, c * r),
        TelWord{},
    }
    return &t
}

// TelWindow Telnet virtual window manager
type TelWindow struct {
    cols int
    rows int
    pc int
    pr int
    words []TelWord
    tempWord TelWord
}

// AddByte Add new byte to TelWindow
func (t *TelWindow) AddByte(b byte) {
    
}