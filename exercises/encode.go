package main

import (
	"fmt"

	"bytes"
	"github.com/ugorji/go/codec"
	"log"
)

type status struct {
	A int
	S string
	B uint8
}

func main() {
	in := status{
		A: 12345,
		S: "testing",
		B: 3,
	}
	//in := map[string]interface{}{"rssi": uint8(12), "lqi": uint8(231), "ts": uint16(12312)}

	w := new(bytes.Buffer)
	enc := codec.NewEncoder(w,
		&codec.MsgpackHandle{RawToString: true, WriteExt: true})

	_ = enc.Encode(in)

	fmt.Println()
	log.Printf("% x", w.Bytes())

}

/*func main() {
	fmt.Println("Hello")

}*/
