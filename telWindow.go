package gotel


import(
    "io"
)

// TelWord Word property in TelWindow
type TelWord struct {
    TextProperty AnsiProperty
    TextWidth uint8
    TextIndex uint8
    TextCode rune
}

// NewTelWindow Create new TelWindow
func NewTelWindow(c int, r int) *TelWindow {
    t := TelWindow {
        c, // cols
        r, // rows
        0, // point col
        0, // point row
        nil,
        AnsiProperty{},
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
    wordDecoder io.Reader
    cWrodProperty AnsiProperty
    words []TelWord
    tempWord TelWord
}

// SetWordDecoder Set decoder like big5 decoder (to utf8)
func (t *TelWindow) SetWordDecoder(r io.Reader) {
    t.wordDecoder = r
}

// AddByte Add new byte to TelWindow
func (t *TelWindow) AddByte(b byte) {
    
}