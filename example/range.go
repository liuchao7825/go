package main

import "fmt"

func main() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {

		sum += num
	}
	fmt.Println("sum:", sum)
	//再数组上使用range将转入index和值两个变量。上面那个例子不需要使用该元素的序号，所以用空白符”_“省略了。有时候需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range 也可以用在map的键值上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来美剧Unicode字符串。第一个参数是字符的索引，第二是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
