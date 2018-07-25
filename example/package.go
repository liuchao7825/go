package main
import "fmt"
func main(){
	resule := getNum()
	fmt.Println(resule())
}
func getNum() func() int{
	i :=0
	return func () int{
		i+=1
		return i
	}
}

/*getNum()这个方法在文档中写的有点简化，完整写法是：

	func getNum() (func() (int)) {//这里返回值是一个方法
			i:=0
			return func() (int){//返回值是一个int类型
				return i
			}	
	}
	
从上边的代码来看当我们在main函数中调用result :=getNum()的时候，获取到result其实是方法getNum()的返回值，也是一个方法
这时候result :=getNum()等价于
	result :=func() int{
		i :=0
		return func()(int) {//返回值是一个int类型
			return i
		}
	}
	
这个时候result 气质只是一个方法，如果打印 fmt.Println(result)的话得到的是内存地址，并不是我们想要的结果，想获取文档中i的值我们还需要这样写
		result : getNum()
		fmt.Println(result())
	
*/