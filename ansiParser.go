package gotel

import(
    "log"
)

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

    // StateC0 In C0 state, 2 char cmd, 00 ~ 1F
    // StateC0 AnsiState = 2

    // StateIntermediate In Intermediate state, 3 char cmd, 20 ~ 2F
    StateIntermediate AnsiState = 2
    
    // StateParameter In Parameter state, 2 char cmd, 30 ~ 3F
    // StateParameter AnsiState = 3

    // StateC1 In C1 state (Uppercase), 2 char cmd, 40 ~ 5F (CSI special)
    // StateC1 AnsiState = 4

    // StateLowercase In Alphabetic state, 2 char cmd, 60 ~ 7E
    // StateLowercase AnsiState = 5

    // StateCSI In CSI state, n char cmd, 1B
    StateCSI AnsiState = 6
)

// AnsiCommand ANSI Command
type AnsiCommand uint8
const (
    // CmdText This byte is text
    CmdText AnsiCommand = 0
    // CmdParsing This byte is do nothing
    CmdParsing AnsiCommand = 1
)

// AnsiParser comment
type AnsiParser struct {
    ansiCmd AnsiCommand
    state AnsiState
    property AnsiProperty
    tempCommand []byte
}

// NewAnsiParser Create new ANSI parser
func NewAnsiParser() *AnsiParser {
    a := AnsiParser {
        CmdText,
        StateText,
        AnsiProperty{},
        make([]byte, 0, 16),
    }
    return &a
}

// AddByte Add a byte
func (a *AnsiParser) AddByte(b byte) AnsiCommand {
    switch a.state {
    case StateText:
        switch b {
        case 27: // esc
            a.state = StateESC
            a.ansiCmd = CmdParsing
        default:
            a.ansiCmd = CmdText
        }
    case StateESC:
        switch {
        case b == 0x1B || b == 0x9B: // CSI state (n char)
            a.state = StateCSI
            a.ansiCmd = CmdParsing
        case b >= 0x00 && b <= 0x1F: // C0 state (2 char)
            a.state = StateText
            a.ansiCmd = a.CommandC0(b)
        case b >= 0x20 && b <= 0x2F: // Intermediate state (only support 3 word now)
            a.state = StateIntermediate
            a.ansiCmd = CmdParsing
        case b >= 0x30 && b <= 0x3F: // Parameter state (2 char)
            a.state = StateText
            a.ansiCmd = a.CommandParameter(b)
        case b >= 0x40 && b <= 0x5F: // C1 state (2 char)
            a.state = StateText
            a.ansiCmd = a.CommandC1(b)
        case b >= 0x60 && b <= 0x7E: // Lowercase state (2 char)
            a.state = StateText
            a.ansiCmd = a.CommandLowercase(b)
        case b >= 0x80 && b <= 0x9F: // C1 state (2 char)
            a.state = StateText
            a.ansiCmd = a.CommandC1(b)
        }
    case StateIntermediate:
        a.tempCommand = append(a.tempCommand, b)
        if len(a.tempCommand) == 2 {
            a.ansiCmd = a.CommandIntermediate(a.tempCommand[0], a.tempCommand[1])
            a.tempCommand = a.tempCommand[:0]
            a.state = StateText
        }
    case StateCSI:
    }
    
    return a.ansiCmd
}

// CommandCSI comment
func (a *AnsiParser) CommandCSI(bs []byte) AnsiCommand {

}

// CommandC0 comment
func (a *AnsiParser) CommandC0(b byte) AnsiCommand {
    
}

// CommandIntermediate Only support 3 char command
func (a *AnsiParser) CommandIntermediate(b1 byte, b2 byte) AnsiCommand {
    
}

// CommandParameter comment
func (a *AnsiParser) CommandParameter(b byte) AnsiCommand {
    
}

// CommandC1 comment
func (a *AnsiParser) CommandC1(b byte) AnsiCommand {
    
}

// CommandLowercase comment
func (a *AnsiParser) CommandLowercase(b byte) AnsiCommand {
    
}



