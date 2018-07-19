package main
import "fmt"
func main(){
	var x interface{}
	switch i :=x.(type){
		case nil:
			fmt.Printf("x 的类型 ：%T, import")
		case int:
			fmt.Printf("x 是整型")
		case flaot64:
			fmt.Printf("x 是float64型")
		case bool,string:
			fmt.Printf("x 是bool或string型")
	}
	}

