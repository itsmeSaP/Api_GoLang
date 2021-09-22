// package main

// import (
// 	"fmt"
// )

// //Structs syntax
// type Animal struct {
// 	Name   string
// 	Origin string
// }

// //Inheritance--> bird class will inherit all properties from Animal
// type Bird struct {
// 	Animal
// 	topSpeed int
// 	canFly   bool
// }

// //METHOD
// func (b Bird) printBird() {
// 	fmt.Println(b.Name)
// 	fmt.Println(b.topSpeed)
// 	fmt.Println(b.canFly)
// }

// //FUNCTIONS
// //func name(argunemnt) return_type{....}
// func sum(a int, b int) int {
// 	return a + b

// }

// func main() {
// 	//VARIABLES AND DATATYPES
// 	// fmt.Println("Hello World")
// 	// var a int
// 	// a = 10
// 	// //:= auto declares variable b as string. We dont have to decalre it explicitly
// 	// b := "manjirou"
// 	// c := complex(5, 6)
// 	// //Syntax:- Println and Printf
// 	// fmt.Println("the variable  is: ", a, "and name is: ", b)
// 	// fmt.Printf("the variable a is %v and name is %v\n", a, b)
// 	// fmt.Printf("the variable c is %v and type  is %T", c, c)

// 	//ARRAYS-> Syntax
// 	// arr1 := []int{1, 2, 3}
// 	// arr2 := []string{"draken", "Manjirou", "Baji"}
// 	// fmt.Printf("arr1=%v %T \n", arr1, arr1)
// 	// fmt.Printf("arr2=%v %T \n", arr2, arr2)
// 	// //can also directly access array eleemnts
// 	// fmt.Println(arr1[1])
// 	// //array slicing [1:2]-> means elements from index 1 to index 2, including index 1 and excluding index 2
// 	// fmt.Println(arr1[1:2])

// 	//Maps--> like unordered maps
// 	//keys & values-->Syntax --> map[key_datatype]value_datatype{...}
// 	// map1 := map[string]int{
// 	// 	"a": 1,
// 	// 	"b": 2,
// 	// 	"c": 3,
// 	// }
// 	// fmt.Printf("The map is =%v %T", map1, map1)

// 	//Struct Objects--> declaring an object of type Animal
// 	// animal1 := Animal{
// 	// 	Name:   "Nine-tailed Fox",
// 	// 	Origin: "Japan",
// 	// }
// 	// fmt.Printf("The animal is %v and type is %T \n", animal1, animal1)
// 	// bird1 := Bird{}
// 	// bird1.Name = "crow"
// 	// bird1.Origin = "India"
// 	// bird1.topSpeed = 10
// 	// bird1.canFly = true
// 	// fmt.Printf("The Bird is %v and type is %T \n", bird1, bird1)

// 	//IF- condition
// 	// a := 1
// 	// b := 2
// 	// if a > b {
// 	// 	fmt.Println("a is larger")
// 	// } else {
// 	// 	fmt.Println("b is larger")
// 	// }
// 	// //interesting syntax , taking input from users adn printing
// 	// var inp1 int
// 	// if _, err := fmt.Scanf("%v", &inp1); err != nil {
// 	// 	log.Panicln("Error")
// 	// }
// 	// fmt.Printf("inp1 is %v and type is %T", inp1, inp1)

// 	//FOR - loop
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Print(i, " ")
// 	// 	//or we can also use Println instead of Print
// 	// }
// 	//directly iterating through arrays like in python
// 	// arr1 := []int{1, 2, 3, 4, 5}
// 	// for position, value := range arr1 {
// 	// 	fmt.Println(position, value)
// 	// }
// 	// i := 2
// 	// //FOR acts like WHILE
// 	// for i > 0 {
// 	// 	fmt.Println(i)
// 	// 	i--
// 	// }

// 	//Pointers
// 	// p1 := 10
// 	// //p2 is a pointer variable of type *int
// 	// p2 := &p1
// 	// fmt.Printf("%v %T \n", p1, p1)
// 	// fmt.Printf("%v %T", p2, p2)

// 	//Calling functions
// 	// result := sum(10, 20)
// 	// fmt.Println(result)
// 	// //directly write
// 	// fmt.Println(sum(5, 5))

// 	//METHODS
// 	// bird1 := Bird{}
// 	// bird1.Name = "crow"
// 	// bird1.topSpeed = 10
// 	// bird1.canFly = true
// 	// //calling the METHOD print()
// 	// bird1.printBird()

// }