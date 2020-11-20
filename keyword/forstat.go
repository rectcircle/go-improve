package keyword

import (
	"fmt"
	"strconv"
)

func ForExperiment() {
	{
		// 只会迭代 3 次
		arr := []int{1, 2, 3}
		for _, v := range arr {
			arr = append(arr, v)
		}
		// 返回 [1 2 3 1 2 3]
		fmt.Println(arr)
	}
	{
		// 永远输出3
		arr := []int{1, 2, 3}
		newArr := []*int{}
		for _, v := range arr {
			newArr = append(newArr, &v)
		}
		for _, v := range newArr {
			fmt.Println(*v)
		}
	}
	{
		// 优化为 runtime.memclrNoHeapPointers 调用
		arr := []int{1, 2, 3}
		for i := range arr {
			arr[i] = 0
		}
	}
	{
		// map 遍历具有随机性
		hash := map[string]int{
			"1": 1,
			"2": 2,
			"3": 3,
		}
		for k, v := range hash {
			println(k, v)
		}
		for k, v := range hash {
			println(k, v)
		}
	}
	{
		// 删除未遍历的元素，不会被遍历到
		fmt.Println("map 删除未遍历的元素")
		hash := map[string]int{
			"1": 1,
			"2": 2,
		}
		for k := range hash {
			fmt.Println(k)
			if k == "1" {
				delete(hash, "2")
			} else {
				delete(hash, "1")
			}
		}
	}
	{
		// 边遍历变添加元素是否会被遍历到：不确定，取决于元素是否会添加到未被遍历过的桶之中
		// https://segmentfault.com/q/1010000012242735
		fmt.Println("map 边遍历变添加元素是否会被遍历到")
		hash := map[string]int{
			"1": 1,
			"2": 2,
			"3": 3,
			"4": 4,
		}
		for k := range hash {
			fmt.Println(k)
			for i := 4; i < 1000; i++ {
				hash[strconv.Itoa(i)] = i
			}
		}
		fmt.Println("len(hash) = ", len(hash))
	}
}