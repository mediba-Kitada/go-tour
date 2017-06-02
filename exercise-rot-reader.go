package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(p byte) (b byte) {
	switch {
	case (p >= 'A' && p < 'N') || (p >= 'a' && p < 'n'):
		b = p + 13
	case (p > 'M' && p <= 'Z') || (p > 'm' && p <= 'z'):
		b = p - 13
	default:
		b = p
	}
	return
}

// ROT13換字式暗号をすべてのアルファベット文字に適用して読み出し
// Read関数を実装し、io.Readerインターフェースを満たす必要がある
func (rot *rot13Reader) Read(bs []byte) (int, error) {

	n, e := rot.r.Read(bs)
	for i := 0; i < n; i++ {
		bs[i] = rot13(bs[i])
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
