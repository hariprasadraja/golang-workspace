package main

import "fmt"

// function which retuns a single value

func function1(a string) string {
	defer fmt.Println("this program runs at exit")
	for i := 0; i < 5; i++ {
		fmt.Println(a)
	}
	return "differ ends"
}
func main() {
	fmt.Println(function1("hi"))
}

// }

// funtion which retruns a pointer

// func pointerfunction(a *int) (*int){
// fmt.Println(a)
// return a
// }

// var number *int
// // number = 10, then error:"cannot use 10 (type int) as type *int in assignment"
// value := 10
// number = &value
// fmt.Println(number)
// fmt.Println(*number)
// *number =20
// fmt.Println(*number)
// fmt.Println(number)
//
// var newnumber *int
//
// // *newnumber = 90 , it will throw an error since newnumber does not point to any variable , So
// newnumber = new (int)
// *newnumber = 90

//.Println("new number",*newnumber)

//fmt.Printf("the number is %v", )

// calling the function with value
//pointerfunction(number)
//calling the function with address
//pointerfunciton(&number), then error:"cannot use &number (type **int) as type *int in argument to pointerfunction"
//pointerfunciton(*number), then error:"cannot use *number (type **int) as type *int in argument to pointerfunction"

// pointer returns
// i := int(42)
//     fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
//     p := &i
//     fmt.Printf("2. main  -- p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
//     byval(p)
//     fmt.Printf("5. main  -- p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
//     fmt.Printf("6. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
//
// }
// func byval(q *int) {
//     fmt.Printf("3. byval -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
//     *q = 4143
//     fmt.Printf("4. byval -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
//     fmt.Println("this is the q" , q)
//     q = nil
// }
