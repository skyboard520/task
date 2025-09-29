package task2

import (
	"fmt"
	"math"
)

/**
1.题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func getAreaAndPerimeter(shape Shape) (float64, float64) {
	return shape.Area(), shape.Perimeter()
}

func Struct1() {
	rectangle := Rectangle{width: 5, height: 10}
	circle := Circle{radius: 7}
	fmt.Println("Rectangle:")
	fmt.Println(getAreaAndPerimeter(rectangle))
	fmt.Println("Circle:")
	fmt.Println(getAreaAndPerimeter(circle))
}

/**
2.题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.Age, e.EmployeeID)
}
func Struct2() {
	employee := Employee{Person{"John", 30}, 1234}
	employee.PrintInfo()
}
