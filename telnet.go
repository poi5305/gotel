package gotel

import(
    "io"
    "bytes"
)

const (
    // NULL null 
    NULL byte = '\x00'
    // SE End of subnegotiation parameters
    SE byte = '\xf0'
    // NOP No operation
    NOP byte = '\xf1'
    // DM DataMark The data stream portion of a Synch. This should always be accompanied by a TCP Urgent notification
    DM byte = '\xf2'
    // BRK Break NVT character BRK
    BRK byte = '\xf3'
    // IP Interrupt Process
    IP byte = '\xf4'
    // AO Abort output
    AO byte = '\xf5'
    // AYT Are you there
    AYT byte = '\xf6'
    // EC Erase character
    EC byte = '\xf7'
    // EL Erase line
    EL byte = '\xf8'
    // GA Go ahead signal
    GA byte = '\xf9'
    // SB Indicates that what follows is subnegotiation of the indicated option
    SB byte = '\xfa'
    // WILL Indicates the desire to begin performing, or confirmation that you are now performing, the indicated option
    WILL byte = '\xfb'
    // WONT Indicates the refusal to perform, or continue performing, the indicated option
    WONT byte = '\xfc'
    // DO Indicates the request that the other party perform, or confirmation that you are expecting the other party to perform, the indicated option
    DO byte = '\xfd'
    // DONT Indicates the demand that the other party stop performing, or confirmation that you are no longer expecting the other party to perform, the indicated option
    DONT byte = '\xfe'
    // IAC Interpret as command
    IAC byte = '\xff'
    //ECHO Echo
    ECHO byte = '\x01'
    // SGA Suppress go ahead
    SGA byte = '\x03'
    // TT Terminal type
    TT byte = '\x18'
    // IS comment
    IS byte = '\x00'
    // SEND comment
    SEND byte = '\x01'
    // WS Window size
    WS byte = '\x1f'
)

type telState int
const (
    stateData telState = iota
    stateIAC telState = iota
    stateSB telState = iota
    stateWill telState = iota
    stateWont telState = iota
    stateDo telState = iota
    stateDont telState = iota
)


// TelConfig comment
type TelConfig struct {
    CReadBuffer int
    CLogLevel LogLevel
    CTerminalType string
    CDoToWillCmdList map[byte] func(byte) bool
    CWillToDoCmdList map[byte] func(byte) bool
    CSubCmdListeners map[byte] func(byte, []byte) bool
}

// RegisterSubCmdListener Register sub command listener
func (t *TelConfig) RegisterSubCmdListener(code byte, listener func(byte, []byte) bool) {
    t.CSubCmdListeners[code] = listener
}

// UnregisterSubCmdListener UnRegister sub command listener
func (t *TelConfig) UnregisterSubCmdListener(code byte) {
    delete(t.CSubCmdListeners, code)
}

// GoTel comment
type GoTel struct {
    // public
    Config TelConfig
    
    // private
    state telState
    rw io.ReadWriter
    dataBuf *bytes.Buffer
    subData []byte
    err error

}

// New new GoTelnet implemention
func New(rw io.ReadWriter) *GoTel {
    t := GoTel{}
    t.Config = TelConfig{}
    t.UseDefaultConfig()
    
    t.state = stateData
    t.rw = rw
    t.dataBuf = bytes.NewBuffer(make([]byte, 0, t.Config.CReadBuffer))
    t.subData = make([]byte, 0, 32)
    t.err = nil
    return &t
}

// UseDefaultConfig comment
func (t *GoTel) UseDefaultConfig() {
    t.Config.CReadBuffer = 1024
    t.Config.CLogLevel = LogDebug
    t.Config.CTerminalType = "VT100"
    
    t.Config.CDoToWillCmdList = make(map[byte] func(byte) bool)
    t.Config.CWillToDoCmdList = make(map[byte] func(byte) bool)
    t.Config.CSubCmdListeners = make(map[byte] func(byte, []byte) bool)
    
    t.Config.CWillToDoCmdList[ECHO] = nil
    t.Config.CWillToDoCmdList[SGA] = nil
    t.Config.CDoToWillCmdList[TT] = nil
    t.Config.CDoToWillCmdList[WS] = nil // TODO check window size
    
    t.Config.RegisterSubCmdListener(TT, func(byte, []byte) bool {
        cmd := []byte {IAC, SB, TT, IS}
        cmd = append(cmd, t.Config.CTerminalType...)
        cmd = append(cmd, IAC, SB)
        t.SendCommand(cmd...)
        return true
    })
}

func (t *GoTel) Read(p []byte) (int, error) {
    for t.dataBuf.Len() == 0 {
        b := make([]byte, t.Config.CReadBuffer, t.Config.CReadBuffer)
        n, err := t.rw.Read(b)
        t.err = err
        if err != nil {
            return 0, err
        }
        for i := 0; i < n; i++ {
            t.addByte(b[i])
        }
    }
    n, err := t.dataBuf.Read(p)
    t.err = err
    return n, err
}

func (t *GoTel) Write(p []byte) (int, error) {
    return t.rw.Write(p)
}

// SendCommand comment
func (t *GoTel) SendCommand(codes... byte) {
    cmd := make([]byte, len(codes) + 1)
    cmd[0] = byte(IAC)

	for i, code := range codes {
		cmd[i + 1] = code
	}
	t.rw.Write(cmd)
}

func (t *GoTel) addByte(b byte) {
    switch t.state {
    case stateData:
        t.state = t.inStateData(b)
    case stateIAC:
        t.state = t.inStateIAC(b)
    case stateSB:
        t.state = t.inStateSB(b)
    case stateWill:
        t.state = t.inStateWill(b)
    case stateWont:
        t.state = t.inStateWont(b)
    case stateDo:
        t.state = t.inStateDo(b)
    case stateDont:
        t.state = t.inStateDont(b)
    }
}

func (t *GoTel) inStateData(b byte) telState {
    newState := stateData
    if b == byte(IAC) {
        newState = stateIAC
    } else {
        t.dataBuf.WriteByte(b)
    }
    return newState
}

func (t *GoTel) inStateIAC(b byte) telState {
    newState := stateIAC
    switch (b) {
    case WILL:
        newState = stateWill
    case WONT:
        newState = stateWont
    case DO:
        newState = stateDo
    case DONT:
        newState = stateDont
    case SB:
        newState = stateSB
    default:
        newState = stateData
    }
    return newState
}

func (t *GoTel) inStateSB(b byte) telState {
    newState := stateSB
    t.subData = append(t.subData, b)
    subDataLen := len(t.subData)
    if len(t.subData) > 2 && t.subData[subDataLen - 2] == IAC && t.subData[subDataLen - 1] == SE {
        t.subCommand(t.subData[0], t.subData[1 : subDataLen - 3])
        t.subData = t.subData[ : 0]
        newState = stateData
    }
    return newState
}

func (t *GoTel) inStateWill(b byte) telState {
    if t.Config.CWillToDoCmdList != nil {
        do, ok := t.Config.CWillToDoCmdList[b]
        if ok {
            if do == nil || do(b) {
                t.SendCommand(IAC, DO, b)
                return stateData
            }
        }
    }
    t.SendCommand(IAC, DONT, b)
    return stateData
}

func (t *GoTel) inStateDo(b byte) telState {
    if t.Config.CDoToWillCmdList != nil {
        do, ok := t.Config.CDoToWillCmdList[b]
        if ok {
            if do == nil || do(b) {
                t.SendCommand(IAC, WILL, b)
                return stateData
            }
        }
    }
    t.SendCommand(IAC, DONT, b)
    return stateData
}

func (t *GoTel) inStateWont(b byte) telState {
    return stateData
}

func (t *GoTel) inStateDont(b byte) telState {
    return stateData
}

func (t *GoTel) subCommand(code byte, content []byte) {
    if t.Config.CSubCmdListeners != nil {
        listener, ok := t.Config.CSubCmdListeners[code]
        if ok {
            success := listener(code, content)
            if success {
                return
            } 
        }
    }
}