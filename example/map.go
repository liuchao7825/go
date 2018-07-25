package main
import "fmt"
func main(){
	var countryCapitalMap map[string]string /*创建集合*/
	countryCapitalMap = make (map[string]string)
	
	/*map插入key -value对， 各个国家对应的首都*/
	countryCapitalMap ["france"] = "Paris"
	countryCapitalMap ["Italy"]  = "罗马"
	countryCapitalMap ["Japan"]  = "东京"
	countryCapitalMap ["India"]  = "新德里"
	
	/*使用键值输出地图之*/
	for country := range countryCapitalMap{
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}
	
	/*检查元素在集合中是否存在*/
	captial, ok := countryCapitalMap ["美国"]   /*如果是真实的，则存在，否则不存在*/
    /*fmt.Println(captial) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("美国的首都是", captial)
    } else {
        fmt.Println("美国的首都不存在")
    }	
}
