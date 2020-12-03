package typesystem

import "fmt"

type MyInt int32

func (a MyInt) add(b MyInt) int32 {
	return int32(a) + int32(b)
}

type MyInt2 MyInt // 等价于 type MyInt2 int32

func (a MyInt2) subtract(b MyInt2) int32 {
	return int32(a) - int32(b)
}

type MyInt3 = MyInt

func (a MyInt3) Multiply(b MyInt3) int32 {
	return int32(a) * int32(b)
}

type MyInt4 = int32

// 以下等价于：func (a int32) Divide(b int32) int32 {
// 因此报错 invalid receiver int32 (basic or unnamed type)compiler
// func (a MyInt4) Divide(b MyInt4) int32 {
// 	return a / b
// }

func TypeSyntax() {
	var a MyInt = 1
	fmt.Println(a + a)
	fmt.Println(a.add(a))

	var b MyInt2 = 2
	fmt.Println(b + b)
	// fmt.Println(b.add(b))
	fmt.Println(b.subtract(b))

	var c MyInt3 = a
	fmt.Println(c.add(c))
	fmt.Println(c.Multiply(c))
	fmt.Println(a.Multiply(a))
	fmt.Println(a.add(a))
}
