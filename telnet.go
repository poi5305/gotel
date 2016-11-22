package gotel

import(
    "io"
    "bytes"
)



// TelConfig comment
type TelConfig struct {
    CReadBuffer int
    CLogLevel LogLevel
}

// GoTel comment
type GoTel struct {
    // public
    Config TelConfig
    
    // private
    rw io.ReadWriter
    streamBuf []byte
    dataBuf *bytes.Buffer
    err error

}

// New new GoTelnet implemention
func New(rw io.ReadWriter) *GoTel {
    t := GoTel{}
    t.Config = TelConfig {
        1024, // CReadBuffer
        LogDebug, // LogLevel
    }
    
    t.rw = rw
    t.dataBuf = bytes.NewBuffer(make([]byte, 0, t.Config.CReadBuffer))
    t.err = nil
    return &t
}

func (t *GoTel) Read(p []byte) (int, error) {
    for t.dataBuf.Len() == 0 {
        b := make([]byte, t.Config.CReadBuffer, t.Config.CReadBuffer)
        t.rw.Read(b)
    }
    n, err := t.dataBuf.Read(p)
    t.err = err
    return n, err
}

func (t *GoTel) Write(p []byte) (int, error) {
    return t.rw.Write(p)
}

func (t *GoTel) addByte(b byte) {
    
}
