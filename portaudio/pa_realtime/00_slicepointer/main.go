package main

import "fmt"

func main() {
	in := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}
	out := make([]int16, 9)

	//fmt.Println("1111111111111111111111111")
	//for i , _ := range in{
	//	fmt.Println("in: ", &in[i], "out: ", &out[i])
	//}
	//fmt.Println("2222222222222222222222222")
	// pointer 変わらない
	//for i , _ := range in{
	//	fmt.Println("in: ", &in[i], "out: ", &out[i])
	//}
	//fmt.Println("af")
	//for i , _ := range in{
	//	out[i] = in[i]
	//	fmt.Println("in: ", &in[i], "out: ", &out[i])
	//}
	//fmt.Println("3333333333333333333333333")
	//pointer 変わる (新しいやつできる)
	//for i , _ := range in{
	//	fmt.Println("in: ", &in[i], "out: ", &out[i])
	//}
	//fmt.Println("af")
	//out = append([]int16{}, in...)
	//for i , _ := range in{
	//	fmt.Println("in: ", &in[i], "out: ", &out[i])
	//}
	//fmt.Println("4444444444444444444444444")
	//pointer 変わる (inと同じポインタになる)
	//for i, _ := range in {
	//	fmt.Println("in: ", &in[i], "out: ", &out[i])
	//}
	//fmt.Println("af")
	//out = in
	//for i, _ := range in {
	//	fmt.Println("in: ", &in[i], "out: ", &out[i])
	//}
	fmt.Println("5555555555555555555555555")
	//pointer 変わらない
	for i, _ := range in {
		fmt.Println("in: ", &in[i], "out: ", &out[i])
	}
	fmt.Println("af")
	copy(out, in)
	for i, _ := range in {
		fmt.Println("in: ", &in[i], "out: ", &out[i])
	}

}
