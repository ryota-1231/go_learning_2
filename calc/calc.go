package calc

func Sum(a, b int) int {
	return a + b
}

// 外部パッケージにエクスポートするにはキャピタル（大文字）
type Person struct {
	Name string
	Age  int
}

var Public string = "Public"
var private string = "private"
