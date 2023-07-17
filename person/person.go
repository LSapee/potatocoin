package person

import "fmt"

type Person struct {
	name string
	age  int
}

// 복사본을 원하지 않을때는 *포인터를 붙여준다.
func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println("SeeDetails' le:", p)
}

// copy값을 사용해도 상관없을 때
func (p Person) Name() string {
	return p.name
}

//main에 넣었던 값들
//le := person.Person{}
//le.SetDet때ails("lee", 99)
//fmt.Println("Main's lee", le)
