package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// ASCII文字 'A'の無限ストリームを出力するReader型を実装
func (mr MyReader) Read(b []byte) (int, error) {
	// ASCII文字'A'は、引数に格納する。どのように呼び出し元に返しているのか..
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
