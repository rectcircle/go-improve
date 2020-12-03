package innertype

import "fmt"

func MapExperiment() {
	// 字面量方式创建
	h1 := map[string]int{
		"1": 2,
		"3": 4,
		"5": 6,
	}
	fmt.Println(h1)

	// 创建并指定容量
	h2 := make(map[string]int, 3)
	h2["1"] = 2
	h2["3"] = 4
	h2["5"] = 6

	// 访问
	fmt.Println(h2["1"]) // 不存在将返回零值
	fmt.Println(h2["2"]) // 不存在将返回零值
	// 访问并判断是否存在
	if e, ok := h2["1"]; ok {
		fmt.Println(ok, e)
	}
	// 写入
	h2["7"] = 8
	fmt.Println(h2)
	// 删除
	delete(h2, "7")
	fmt.Println(h2)
	// 遍历 https://golang.org/ref/spec#For_statements
	for k, v := range h2 {
		fmt.Println(k, v)
		// 迭代中删除、创建、修改都是是安全的
		delete(h2, k)
	}
	fmt.Println(h2)

	// 线层安全 map sync.Map
}
