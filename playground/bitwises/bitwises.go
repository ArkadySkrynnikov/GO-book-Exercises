package bitwises

import "fmt"

func BitwiseTry(x uint64) int {
	var res int

	//2525
	//4096 2048 1024 512 256 128 64 32 16 8 4 2 1
	//  0    1   0    0   1   1   1  0  1 1 1 0 1

	// ShiftLeft()
	// ShiftRight()
	// BitwiseAnd()
	// BitwiseOr()
	// BitwiseXor()
	// BitwiseAndNot()
	BitwiseNot()
	return res

}

func BitwiseNot() {
	// Для беззнаковых чисел (unsigned integers)
	var a uint8 = 0b00001111 // Десятичное 15
	var resultUnsigned = ^a  // Инвертируем биты: 11110000
	fmt.Printf("Unsigned NOT: %08b (в десятичной: %d)\n", resultUnsigned, resultUnsigned)
	// Выведет: Unsigned NOT: 11110000 (в десятичной: 240)

	// Для знаковых чисел (signed integers)
	// В Go используется дополнительный код (two's complement) для знаковых чисел,
	// поэтому формула побитового НЕ для знаковых чисел: ^x = -x - 1
	var b int8 = 2
	var resultSigned = ^b
	fmt.Printf("Signed NOT: %d\n", resultSigned)
}

func BitwiseAndNot() {
	var c = 15
	fmt.Printf("c = bits %b, C AND NOT %t\n", c, c&^1 == 14) //c = bits 1111, C AND NOT true
	fmt.Printf("c = bits %b, C AND NOT %t\n", c, c&^5 == 10) //c = bits 1111, C AND NOT true
	var x = 742 &^ 15
	//   1011100110
	//   0000001111
	//   1011100000
	fmt.Println(x) //736
	// 1024 512 256 128 64 32 16 8 4 2 1
	//       1   0   1  1  1  0  0 0 0 0
}

func BitwiseXor() {
	var c = 15
	fmt.Printf("c = bits %b, C XOR %t\n", c, c^1 == 14) //c = bits 1111, C XOR true
	fmt.Printf("c = bits %b, C XOR %t\n", c, c^5 == 10) //c = bits 1111, C XOR true
	var x = 15 ^ 742
	//   0000001111
	//   1011100110
	//   1011101001
	fmt.Println(x) //745
	// 1024 512 256 128 64 32 16 8 4 2 1
	//       1   0   1  1  1  0  1 0 0 1
}

func BitwiseOr() {
	var c = 15
	fmt.Printf("c = bits %b, C OR %t\n", c, c|1 == 15) //c = bits 1111, C AND true
	fmt.Printf("c = bits %b, C OR %t\n", c, c|5 == 15) //c = bits 1111, C AND true
	var x = 15 | 742
	//   0000001111
	//   1011100110
	//   1011101111
	fmt.Println(x) //751
}

func BitwiseAnd() {
	var c = 15
	fmt.Printf("c = bits %b, C AND %t\n", c, c&1 == 1) //c = bits 1111, C AND true
	fmt.Printf("c = bits %b, C AND %t\n", c, c&5 == 5) //c = bits 1111, C AND true
	var x = 15 & 742
	//   0000001111
	//   1011100110
	//   0000000110
	fmt.Println(&x) //6
}

func ShiftLeft() {
	var a int = 7
	fmt.Printf("a = bits %b, digits %d\n", a, a) // a = bits 111, digits 7
	a = a << 2                                   // сдвиг влево на два разряда
	fmt.Printf("a = bits %b, digits %d\n", a, a) // a = bits 1110, digits 14
}

func ShiftRight() {
	var b int = 15
	fmt.Printf("b = bits %b, digits %d\n", b, b) // b = bits 1111, digits 7
	b = b >> 2                                   // сдвиг вправо на два разряда
	fmt.Printf("b = bits %b, digits %d\n", b, b) // b = bits 11, digits 3
}
