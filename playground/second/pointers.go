package second

import "fmt"

func Poinsters() {
	fmt.Println("---------------------------")
	init := 1
	one := &init       // one имеет тип *int и указывает на start
	two := &one        // two имеет тип **int и указывает на указатель one
	fmt.Println(two)   // "0xc00010e030" указывает на указатель one
	fmt.Println(*two)  // "0xc00010c060" указывает на указатель init
	fmt.Println(**two) // "1"
	**two = 2          // Эквивалентно присваиванию х = 2
	fmt.Println(init)  // "2"
	fmt.Println("---------------------------")
	fmt.Println(*f() == *f()) // true
	fmt.Println(f() == f())   // false
	fmt.Println("---------------------------")
	fmt.Println(myIncr(&init))   // 0xc00000a088
	fmt.Println(*myIncr(&init))  // 4
	fmt.Println(init)            // 4
	fmt.Println(bookIncr(&init)) // 5
	fmt.Println(init)            // 5
}

func f() *int {
	v := 1
	return &v
}

func bookIncr(variable *int) int {
	*variable++ // Увеличивает значение, на которое указывает р;
	// не изменяет значение р
	return *variable
}

func myIncr(variable *int) *int {
	*variable++ // Увеличивает значение, на которое указывает р;
	// не изменяет значение р
	return variable
}
