package main

import (
	"errors"
)

// This called contant block
// everytime you call iota, it gets incremented by one, it starts by 0
// iota is used only with constants
// const (
// 	first  = iota + 6  // this is called contant expression
// 	second = 2 << iota //  << is called bitshift operator, mutliplies by 2
// 	third              // it follows the above if not specifed
// )

// // You can declare many contant blocks and iota will reset
// const (
// 	fourth = iota
// )

func main() {
	// Variables decleration ways
	// There is value types and pointer types

	// var name string
	// name = "mohiey"

	// var name string = "mohiey"

	// name := "mohiey"  // implicit initialization syntax

	// num := complex(3, 4)
	// is the same as
	// r, im := real(3), imag(3)

	// var n float64 = 3.1234567
	// var n float64 = 3.123456789123456

	// var name *string = new(string)
	// *name = "mohiey"
	// // no pointer arithmetics are allowed
	// fmt.Println(name)
	// fmt.Println(*name)
	// println(&name)

	// name := "mohiey"
	// ptr := &name
	// fmt.Println(ptr, *ptr)

	// const i = 10
	// // the value of a constant has to determined at compile time, not runtime (assigning using a function)
	// println(i)

	// const c = 3
	// x := c + 4   //interpert c as int
	// z := c + 1.5 //interpert c as float
	// // but if we declared c like this const c int = 3 this will cause an error unless we make explicit conversion like float32(c)

	// fmt.Println(first, second, third, fourth)

	// var arr [3]int
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3

	// arr2 := [3]int{1, 2, 3}
	// arr3 := arr2
	// arr2[0] = 33342 // affects arr2 only
	// fmt.Println(arr2, arr3)

	// ex 1
	// arr := [3]int{1, 2, 3}
	// slice := arr[:] // : is called a slice operator
	// This slice is pointing to the data kept inside the array

	// fmt.Println(arr, slice)
	// [1 2 3] [1 2 3]

	// arr[0] = 77
	// slice[1] = 56
	// fmt.Println(arr, slice)
	// [77 56 3] [77 56 3]

	// ex 2
	// slice1 := []int{1, 2, 3}
	// slice2 := append(slice1, 4, 5)
	// slice1[0] = 3423 // affects slice1 only because append returns a new slice
	// fmt.Println(slice1, slice2)

	// ex 3
	// slice := []int{1, 2, 3, 4, 5}
	// s1 := slice[1:]  //start from index 1 till the end of the slice
	// s2 := slice[:4]  //start from index 0 till index 4 but not including
	// s3 := slice[1:3] //start from index 1 till index 3 but not including
	// s1[0] = 22       // will affect s2 and s3 and the original slice
	// fmt.Println(slice, s1, s2, s3)

	// m1 := map[string]int{"first": 1, "second": 2}
	// m2 := m1
	// fmt.Println(m1, m1["second"])
	// delete(m1, "first") //affects m1 and m2
	// fmt.Println(m1, m2)

	// type user struct {
	// 	id        int
	// 	firstName string
	// 	lastName  string
	// }
	// var u1 user
	// u1.id = 10
	// u1.firstName = "Mohamed"
	// u1.lastName = "Mohiey"

	// u2 := u1
	// u2.id = 979 // affects u2 only
	// println(u1.id, u2.id)

	// var u2 user = user{id: 23, firstName: "Ahamed", lastName: "Mohamed"}
	// u3 := user{id: 34, firstName: "Mostafa", lastName: "Mahmoud"}

	// fmt.Println(u)
	// fmt.Println(u2)
	// fmt.Println(u3)

	// i, e := test()
	// fmt.Println(i, e)

	// what if we need oly the the e from test function? use the write only variable
	// _, e := test()
	// fmt.Println(e)

	// we heve four different types of for loops

	// first type is loop till condition
	// var i int
	// for i < 5 {
	// 		println(i)
	// 		i++
	// }

	// second type is loop till condition with post clause
	// for i := 0; i < 5; i++ {
	// 		println(i)
	// 		if i == 3 {
	// 			break
	// 			// continue
	// 		}
	// }

	//  third type is infinite loop
	// var i int
	// for {
	// 	println(i)
	// 	i++
	// }

	// fourth type is loop over collections
	// nums := []int{1, 2, 3, 4}
	// for i, num := range nums {
	// 	println(i, num)
	// }

	// Branching Contructs
	// panic
	// panic("HELPPPPPP") // you can recover from panic, check the documentation

	// if statements
	// if 3 == 1 {
	// 	println("jjj")
	// }

	// if 3 == 2 {
	// 	println("jjj")

	// } else {
	// 	println("sss")

	// }

	// if 4 == 5 {
	// } else if 6 == 9 {
	// } else {
	// }

	// switch case
	// we have implicit break although we can use it explicitly
	// type User struct{ ID int }
	// u := User{ID: 2}
	// switch u.ID {
	// case 1:
	// 	println("The id of the user is 1")
	// case 2:
	// 	println("The id of the user is 2")
	// 	// fallthrough // explicit fallthrough
	// case 3:
	// 	println("The id of the user is 3")
	// default:
	// 	println("No Match found")
	// }

	// c := newCat()
	// c.talk()
	// fmt.Println(c.color)

}

// func callMe(age int, year int)  is the same as func callMe(age, year int)
// func callMe(age, year int) {

// }

// func hello(name string, age int) bool {
// 	return true
// }

// func start() error { //error is a pointer data type
// 	// return nil
// 	// or
// 	return errors.New("Something went wrong")
// }

// How to return multiple pieces of information
func test() (int, error) {
	return 10, errors.New("Test")
}
