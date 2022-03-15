package main

import "fmt"

type TrainGame struct {
	GroupName string
	Points    []int
	Average   float64
}

type x struct{}

func HogwartTrain() {

	var g TrainGame
	g.GroupName = "Gryffindor"
	g.Points = make([]int, 0)
	for i := 1; i <= 40; i++ {
		g.Points = append(g.Points, i)
	}
	fmt.Println("List of points scored by Gryffindor group students are : ", g.Points)
	sum1 := 0
	for i := 0; i < len(g.Points); i++ {
		sum1 = sum1 + g.Points[i]
	}
	fmt.Println("Sum of points gained by Gryffindor group students are : ", sum1)

	var s TrainGame
	s.GroupName = "Slytherin"
	s.Points = make([]int, 0)
	for i := 41; i <= 80; i++ {
		s.Points = append(s.Points, i)
	}
	fmt.Println("List of points scored by Slytherin group students are : ", s.Points)
	sum2 := 0
	for i := 0; i < len(s.Points); i++ {
		sum2 = sum2 + s.Points[i]
	}
	fmt.Println("Sum of points gained by Slytherin group students are : ", sum2)

	var r TrainGame
	r.GroupName = "Ravenclaw"
	r.Points = make([]int, 0)
	for i := 81; i <= 120; i++ {
		r.Points = append(r.Points, i)
	}
	fmt.Println("List of points scored by Ravenclaw group students are : ", r.Points)
	sum3 := 0
	for i := 0; i < len(r.Points); i++ {
		sum3 = sum3 + r.Points[i]
	}
	fmt.Println("Sum of points gained by Ravenclaw group students are : ", sum3)

	var h TrainGame
	h.GroupName = "Hufflepuff"
	h.Points = make([]int, 0)
	for i := 121; i <= 160; i++ {
		h.Points = append(h.Points, i)
	}
	fmt.Println("List of points scored by Hufflepuff group students are : ", h.Points)
	sum4 := 0
	for i := 0; i < len(h.Points); i++ {
		sum4 = sum4 + h.Points[i]
	}
	fmt.Println("Sum of points gained by Hufflepuff group students are : ", sum4)

	g.Points[8] = g.Points[8] + 4
	g.Points[26] = g.Points[26] + 4
	g.Points[14] = g.Points[14] + 4
	fmt.Println("List of points scored by Gryffindor group students are : ", g.Points)
	sum5 := 0
	for i := 0; i < len(g.Points); i++ {
		sum5 = sum5 + g.Points[i]
	}
	fmt.Println("Sum of points gained by Gryffindor group students are : ", sum5)

	s.Points[9] = s.Points[9] - 2
	s.Points[24] = s.Points[24] - 2
	s.Points[28] = s.Points[28] - 2
	fmt.Println("List of points scored by Slytherin group students are : ", s.Points)
	sum6 := 0
	for i := 0; i < len(s.Points); i++ {
		sum6 = sum6 + s.Points[i]
	}
	fmt.Println("Sum of points gained by Slytherin group students are : ", sum6)

	for i := 0; i < len(r.Points); i++ {
		r.Points[i] = r.Points[i] - 1
	}
	fmt.Println("List of points scored by Ravenclaw group students are : ", r.Points)
	sum7 := 0
	for i := 0; i < len(r.Points); i++ {
		sum7 = sum7 + r.Points[i]
	}
	fmt.Println("Sum of points gained by Ravenclaw group students are : ", sum7)

	for i := 0; i < len(h.Points); i++ {
		h.Points[i] = h.Points[i] + 2
	}
	fmt.Println("List of points scored by Hufflepuff group students are : ", h.Points)
	sum8 := 0
	for i := 0; i < len(h.Points); i++ {
		sum8 = sum8 + h.Points[i]
	}
	fmt.Println("Sum of points gained by Hufflepuff group students are : ", sum8)

	g.Average = float64(sum5 / len(g.Points))
	fmt.Println("Average points scored by Gryffindor group are : ", g.Average)

	s.Average = float64(sum6 / len(s.Points))
	fmt.Println("Average points scored by Slytherin group are : ", s.Average)

	r.Average = float64(sum7 / len(r.Points))
	fmt.Println("Average points scored by Ravenclaw group are : ", r.Average)

	h.Average = float64(sum8 / len(h.Points))
	fmt.Println("Average points scored by Hufflepuff group are : ", h.Average)

}
func main() {
	HogwartTrain()
}
