package innertype

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func StringExperiment() {
	// 单行字符串
	s1 := "123`中文❤😁\"\n\tabc"
	fmt.Println(s1)
	// 多行字符串，不支持转移字符
	s2 := `123中文❤😁"
	abc`
	fmt.Println(s2)

	p := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Println(p.Len, p.Data)

	b1 := []byte(s1);
	s3 := string(b1)
	fmt.Println(s3)

	utf8.RuneCountInString(s2)

	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}

	for _, c := range s2 {
		fmt.Println(string(c))
	}
}