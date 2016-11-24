package gotel_test

import(
    "github.com/poi5305/gotel"
    "net"
    "fmt"
    "testing"
    "log"
    "os"
)


func tcp() {
    conn, _ := net.Dial("tcp", "ptt.cc:23")
    b := make([]byte, 128)
    conn.Read(b)
    fmt.Println(b)
    conn.Close()
}

// TestConn c
func TestConn(t *testing.T) {
    log.SetOutput(os.Stdout)
    conn, _ := net.Dial("tcp", "ptt.cc:23")
    telnet := gotel.New(conn)
    p := make([]byte, 1024)
    telnet.Read(p)
    fmt.Println(string(p))
    telnet.Read(p)
    fmt.Println(string(p))
    telnet.Read(p)
    fmt.Println(string(p))
    telnet.Read(p)
    fmt.Println(string(p))
    telnet.Read(p)
    fmt.Println(string(p))
    telnet.Read(p)
    fmt.Println(string(p))
    fmt.Println(p)
}