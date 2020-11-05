package typesystem

import (
	"fmt"
	"reflect"
)

// NamedType 测试命名类型
func NamedType() {
	var a int32
	var b string
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a).Kind())
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind())
}

func UnnamedType() {
	var a [2]int
	var b []int 
	var c map[string]string
	var d *int32
	var e chan int
	var f struct { a int}
	// var g interface { }
	var h func() int
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a).Kind())
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind())
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(c).Kind())
	fmt.Println(reflect.TypeOf(d), reflect.TypeOf(d).Kind())
	fmt.Println(reflect.TypeOf(e), reflect.TypeOf(e).Kind())
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(f).Kind())
	// fmt.Println(reflect.TypeOf(g), reflect.TypeOf(g).Kind())
	fmt.Println(reflect.TypeOf(h), reflect.TypeOf(h).Kind())
}

type S struct { a int32}

func UnderlyingType() {
	type T1 string  // string
	type T2 T1  // string
	type T3 []string  // []string
	type T4 T3  // []string
	type T5 []T1  // []string
	type T6 T5  // []string

	var a = struct { a int32 } { a: 12 }  // struct { a int32 }
	var b S = a  // struct { a int32 }
	var c S = S(a)  // struct { a int32 }
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a).Kind())
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind())
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(c).Kind())
}

type InterfaceType interface { f() }
type MyStructType struct { a int32 }
type MyStructType2 struct { a int32 }

func Assignability() {
	var a1 int32 = 1
	var b1 int32 = a1
	fmt.Println(b1)

	var a2 = struct {a int32} {a: 1}
	var b2 MyStructType = a2
	var c2 struct {a int32} = b2
	fmt.Println(c2)

	var a3 InterfaceType
	var b3 interface{ f() } = a3
	var c3 InterfaceType = b3
	fmt.Println(c3)
}

func TypeConvert() {

	var a = "123"
	var b = []rune(a)
	var c = []byte(a)
	// var d = int32(a)
	var e = string(int32(65))
	var f = string(b)
	var g = string(c)
	fmt.Println(b)
	fmt.Println(c)
	// fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var s1 = MyStructType { a: 1}
	var s2 MyStructType2 = MyStructType2(s1)
	fmt.Println(s2)
}

func (self MyStructType) f() { fmt.Println(self.a) }
func (self MyStructType2) f() { fmt.Println(self.a) }

func TypeAssertion() {
	var a = MyStructType { a: 1 }
	var b InterfaceType = a

	c := b.(MyStructType)
	c.a = 3
	c.f()
	b.f() // 仍然打印 1 说明 上面 c 是 b 的一份拷贝

	if c, ok := b.(MyStructType); ok {
		fmt.Println("b.(MyStructType) success")
		c.f()
	} else {
		fmt.Println("b.(MyStructType) fail")
	}

	b.f()


	if c, ok := b.(MyStructType2); ok {
		fmt.Println("b.(MyStructType2) success")
		c.f()
	} else {
		fmt.Println(c)
		fmt.Println("b.(MyStructType2) fail")
	}

	switch c := b.(type) {
	case MyStructType: fmt.Println(c); fmt.Println("b is MyStructType")
	case MyStructType2: fmt.Println(c); fmt.Println("b is MyStructType2")
	default: fmt.Println(c); fmt.Println("b is Unknown")
	}
}