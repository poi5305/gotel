package gotel_test

import(
    //"golang.org/x/text/encoding/traditionalchinese"
    "github.com/poi5305/gotel"
    "net"
    "fmt"
    "testing"
    "log"
    "os"
    "io"
    "time"
)


func tcp() {
    conn, _ := net.Dial("tcp", "ptt.cc:23")
    b := make([]byte, 128)
    conn.Read(b)
    fmt.Println(b)
    conn.Close()
}

func reads(r io.Reader) {
    p := make([]byte, 1)
    for i := 0; i < 1610; i++ {
        n, err := r.Read(p)
        if err != nil {
            fmt.Print("\n",i , n, p, err, "\n")
            //return
        }
        time.Sleep(time.Millisecond)
        fmt.Print(p)
    }
}

// TestConn c
func TestConn(t *testing.T) {
    log.SetOutput(os.Stdout)
    conn, _ := net.Dial("tcp", "ptt.cc:23")
    telnet := gotel.New(conn)
    
    //r := traditionalchinese.Big5.NewDecoder().Reader(telnet)
    //reads(r)
    //fmt.Println("T")
    //bb := []byte {32, 32, 161, 253, 27, 91, 51, 48}
    //uu, _ := traditionalchinese.Big5.NewDecoder().Bytes(bb)
    //fmt.Println(bb, string(uu))
    reads(telnet)
    
    
// [32]2016/11/28 01:08:10 [Receive data length 27]
// [27]2016/11/28 01:08:10 [Receive data length 91]
// [91]2016/11/28 01:08:10 [Receive data length 59]
// [59]2016/11/28 01:08:10 [Receive data length 51]
// [51]2016/11/28 01:08:10 [Receive data length 55]
// [55]2016/11/28 01:08:10 [Receive data length 59]
// [59]2016/11/28 01:08:10 [Receive data length 52]
// [52]2016/11/28 01:08:10 [Receive data length 48]
// [48]2016/11/28 01:08:10 [Receive data length 109]
// [109]2016/11/28 01:08:10 [Receive data length 162]
// 2016/11/28 01:08:10 [Receive data length 172]
// [226][149][177]2016/11/28 01:08:10 [Receive data length 32]
// [32]2016/11/28 01:08:10 [Receive data length 32]
// [32]2016/11/28 01:08:10 [Receive data length 161]
// 2016/11/28 01:08:10 [Receive data length 27]
// 
// 1813 0 [32] traditionalchinese: invalid Big5 encoding
}