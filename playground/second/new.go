package second

import "fmt"

func NewTry() {
	p := new(int)   // р, имеющий тип *int, указывает на неименованную переменную типа int
	fmt.Println(p)  // 0xc00000a0a0
	fmt.Println(*p) // "0"
	*p = 2          // Устанавливает значение этой переменной равным 2
	fmt.Println(*p) // "2"
	fmt.Println("----------------------")
	x := newInt()
	*x++
	fmt.Println(x)
	fmt.Println(*x)
	fmt.Println("----------------------")
	y := newDummyInt()
	*y++
	fmt.Println(y)
	fmt.Println(*y)
}

func newInt() *int {
	return new(int)
}
func newDummyInt() *int {
	var dummy int
	return &dummy
}
