package target

import "bytes"

type MyBuffer struct {
	bytes.Buffer // 埋め込みフィールド
	Name string
	Size int
}


func main() {
	var b MyBuffer
	_ = b
}