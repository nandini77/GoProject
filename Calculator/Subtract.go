package operations

import "fmt"

func SubtractOperation() {

	var x1 int8 = 32
	var x2 int8 = 72
	r1 := x2 - x1
	fmt.Println("Result 1 is : ", r1)

	var x3 int16 = -2323
	var x4 int16 = -8218
	r2 := x4 - x3
	fmt.Println("Result 2 is : ", r2)

	var x5 int32 = -672323
	var x6 int32 = -932218
	r3 := x6 - x5
	fmt.Println("Result 3 is : ", r3)

	var x7 int64 = 5647672
	var x8 int64 = 989232218
	r4 := x8 - x7
	fmt.Println("Result 4 is : ", r4)

	var y1 uint8 = 56
	var y2 uint8 = 98
	r5 := y2 - y1
	fmt.Println("Result 5 is : ", r5)

	var y3 uint16 = 7656
	var y4 uint16 = 8998
	r6 := y4 - y3
	fmt.Println("Result 6 is : ", r6)

	var y5 uint32 = 877956
	var y6 uint32 = 964598
	r7 := y6 - y5
	fmt.Println("Result 7 is : ", r7)

	var y7 uint64 = 45768956
	var y8 uint64 = 8767655498
	r8 := y8 - y7
	fmt.Println("Result 8 is : ", r8)

	var z1 float32 = 324.654
	var z2 float32 = 523.987
	r9 := z2 - z1
	fmt.Println("Result 9 is : ", r9)

	var z3 float64 = 12324.6544
	var z4 float64 = 76123.9687
	r10 := z4 - z3
	fmt.Println("Result 10 is : ", r10)

}
