package typesystem

import "fmt"

type I0 interface {
	NonPtrReceiver()
}

type I1 interface {
	PtrReceiver(a int32)
	NonPtrReceiver()
}

type S1 struct {
	a int32
}

func (s *S1) PtrReceiver(a int32) {
	s.a = a
}

func (s S1) NonPtrReceiver() {
	fmt.Println(s.a)
}

func NonPtrOrPtrReceiver() {
	var a S1 = S1{1}
	var b I1 = &a
	// var c I1 = a
	b.PtrReceiver(2)
	b.NonPtrReceiver()
	// c.NonPtrReceiver()
	var c I0 = b  // 接口也可以赋值给接口
	c.NonPtrReceiver()
}

func InterfaceNil() {
	var a interface{} = nil
	fmt.Println(a)  // <nil>
	fmt.Println(a == nil)  // true
	var b *int32 = nil
	a = b
	fmt.Println(a)  // <nil>
	fmt.Println(a == nil)  // false
}


type ParentI1 interface {
	A()
}

type ParentI2 interface {
	B()
}

type ChildI1 interface {
	ParentI1
	ParentI2
}

type S2 struct{
	a int32
}

func (f S2) A(){
	fmt.Println("a")
}

func (f S2) B(){
	fmt.Println("b")
}

func InterfaceNested() {
	var a S2 = S2{1}
	var b ChildI1 = a
	b.A()
	b.B()
}

type I2 interface {
	NeedCallB()
	B()
}

type S3 struct {}

func (self *S3) NeedCallB() {
	self.B()
}

func (self *S3) B() {
	fmt.Println("S3.B()")
}

type S3Child1 struct {
	S3
}

func (self *S3Child1) B() {
	fmt.Println("S3Child1.B()")
}

func StructNested() {
	var a I2 = &S3Child1{}
	a.NeedCallB()
	a.B()
}
