package main

//
//import "fmt"
//
//func plus(a, b int, potato string) (int, string) {
//	return a + b, potato
//}
//func allPlus(a ...int) int {
//	ans := 0
//	// range를 사용하면 index와 값을 꺼내오는데 index를 무시하기위해서 _를 사용
//	for _, item := range a {
//		ans += item
//	}
//	return ans
//}
//
//// structs
//type person struct {
//	name string
//	age  int
//}
//
//type myKoreaAge struct {
//	name string
//	age  int
//}
//
//func (p person) sayHello() {
//	fmt.Printf("Hello my name is %s and i'm %d\n", p.name, p.age)
//}
//
//func (a myKoreaAge) sayMyAge() {
//	fmt.Printf("안녕하세요. 저의 %s 입니다. \n저의 한국 나이는 %d입니다.", a.name, a.age)
//}
//
//func main() {
//	//struct
//	// pppp := person{"pppp",12}
//	pppp := person{name: "pppp", age: 12}
//	pppp.sayHello()
//	//fmt.Println(pppp)
//	myAge := myKoreaAge{"lee", 29}
//	myAge.sayMyAge()
//
//	fmt.Println("감자 코인을 만들자!")
//	ans, name := plus(2, 2, "potato")
//	p := allPlus(1, 2, 3, 4, 5, 6, 6, 7, 7, 8, 8)
//	fmt.Println(ans, name)
//	fmt.Println(p)
//	nam := "감자 코인을 만들자"
//	for index, letter := range nam {
//		fmt.Println(index, letter)
//		fmt.Println(index, string(letter))
//	}
//	x := 1234214125
//	xAsBinary := fmt.Sprintf("%b\n", x) // string으로 포멧 해주기
//	fmt.Println(x, xAsBinary)
//
//	// array 길이가 정해져있을때 java = arr 느낌, c++ = arr느낌
//	foods := [3]string{"potato", "pasta", "egg"}
//	for _, dish := range foods {
//		fmt.Println(dish)
//	}
//	//slice 길이 제한이 없을때. java = ArrayList느낌, c++ =vector느낌
//	foodss := []string{"potato", "pasta", "egg"}
//	for i := 0; i < len(foodss); i++ {
//		fmt.Println(foodss[i])
//	}
//
//	// Go랭에서는 append 방식이 좀 다름 c++ = foodss.push_back("apple");
//	fmt.Printf("%v\n", foodss)
//	foodss = append(foodss, "apple")
//	fmt.Printf("%v\n", foodss)
//
//}
