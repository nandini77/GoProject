package operations

import "fmt"

func MultiplyOperation() {

	var x1 int8 = 32
	var x2 int8 = 22
	r1 := x1 * x2
	fmt.Println("Result 1 is : ", r1)

	var x3 int16 = -2323
	var x4 int16 = -2218
	r2 := x3 * x4
	fmt.Println("Result 2 is : ", r2)

	var x5 int32 = -672323
	var x6 int32 = -232218
	r3 := x5 * x6
	fmt.Println("Result 3 is : ", r3)

	var x7 int64 = 5647672323
	var x8 int64 = 989232218
	r4 := x7 * x8
	fmt.Println("Result 4 is : ", r4)

	var y1 uint8 = 56
	var y2 uint8 = 98
	r5 := y1 * y2
	fmt.Println("Result 5 is : ", r5)

	var y3 uint16 = 7656
	var y4 uint16 = 8998
	r6 := y3 * y4
	fmt.Println("Result 6 is : ", r6)

	var y5 uint32 = 877956
	var y6 uint32 = 764598
	r7 := y5 * y6
	fmt.Println("Result 7 is : ", r7)

	var y7 uint64 = 45768956
	var y8 uint64 = 8765498
	r8 := y7 * y8
	fmt.Println("Result 8 is : ", r8)

	var z1 float32 = 324.654
	var z2 float32 = 123.987
	r9 := z1 * z2
	fmt.Println("Result 9 is : ", r9)

	var z3 float64 = 12324.6544
	var z4 float64 = 76123.9687
	r10 := z3 * z4
	fmt.Println("Result 10 is : ", r10)

}
