package innertype

import (
	"fmt"
	"reflect"
	"unsafe"
)

func SliceExperiment() {
	// 基本使用
	a := [...]int32 {1, 2, 3}
	// 创建的三种方式
	// 方式1：从数组创建
	s1 := a[0:3]
	s2 := s1[1:2:3]
	// 方式2：字面量创建
	s3 := []int32{1, 2, 3}
	// 方式3：make 创建
	s4 := make([]int32, 10)
	fmt.Println(s1, s2, s3, s4)

	// 访问、修改、删除切片元素，拷贝切片
	s1[1] = -2
	s1[100] = 1
	// 可以发现从数组和切片创建的切片数据在没有append操作之前共享底层数据
	fmt.Println(a, s1, s2)
	// append 之后 且 超过容量后，会脱离共享
	s1_2 := append(s1, int32(-3))
	fmt.Println(a, s1, s1_2, s2)
	// append 之后 但 不超过容量，不会脱离共享
	fmt.Println(len(s2), cap(s2))
	s2_2 := append(s2, int32(-3))
	fmt.Println(a, s1, s1_2, s2, s2_2)
	// 删除元素，利用append + slice 实现
	s3_2 := append(s3[:1], s3[2:]...)
	fmt.Println(s3, s3_2)
	// 拷贝切片 dest 与 source 将脱离共享
	var s5 []int32 = make([]int32, 3, 3)
	copy(s5, a[:])
	s5[1] = 12
	fmt.Println(a, s5)

	// slice 运行时底层类型
	// type SliceHeader struct {
	// 	Data uintptr
	// 	Len  int
	// 	Cap  int
	// }

	fmt.Println(len(s3), cap(s3))
	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&s3))
	fmt.Println(sh1)
}
