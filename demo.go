package share

import "fmt"

type demo struct {}

func New() *demo {
	return new(demo)
}

func (demo) SayHi() {
	fmt.Println("Hello, Go share!")
}
