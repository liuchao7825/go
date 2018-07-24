/*Go语言允许想函数传递指针，只需在函数定义的参数上设置为指针类型即可。*/

package main
import "fmt"
func main(){
	/*定义局部变量*/
	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a 的值： %d\n", a )
	fmt.Printf("减缓前 b 的值： %d\n", b )
	/* 调用函数用于减缓值
	&a指向 a 的变量的地址
	&b指向 b 的变量的地址
	*/
	swap (&a, &b)
	fmt.Printf("交换后的 a 的值：%d\n", a )
	fmt.Printf("交换后的 b 的值：%d\n", b )
}
func swap(x *int,y *int){
	var temp int
	temp = *x	/*保存 x 地址的值*/
	*x = *y		/*将 y 复制给 x  */
	*y = temp	/*将temp复制给 y */
}