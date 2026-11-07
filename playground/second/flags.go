package second

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", " ", "разделитель")

func EchoFlags() {
	flag.Parse()
	fmt.Println(*n)
	fmt.Println(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
