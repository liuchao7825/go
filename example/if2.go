package main
import "fmt"
func main(){
	/* 局部变量定义 */
	var a int = 100 ;
	/* 判断部位表达式 */
	if a < 20 {
		/* 如果条件为true 则执行一下语句 */
		fmt.Printf("a 小于 20 \n");
	}else {
		/*如果条件为false则执行以下语句*/
		fmt.Print("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 %d\n", a);

	}
