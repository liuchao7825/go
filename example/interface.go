package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("nokiaPhone")
}

type IPhone struct {
}

func (Iphone IPhone) call() {
	fmt.Println("Iphone")
}
func main() {
	var phone Phone
	phone = new(NokiaPhone)

	phone.call()
	phone = new(IPhone)
	phone.call()
}

/*在次此，定义了一个加快Phone，接口里面有一个方法call（）。然后我们在main函数里定义了一个Phone类型变量，并反别为之复制为NokiaPhone和Iphone。然后调用call（）方法*/
