package main

import "fmt"

const pi = 3.147

func main() {
	var radian, weight, height, area, base float32

	var function int
	fmt.Print("Enter Function \n1) triangle\n2) ractangle\n3) circle\n....")
	fmt.Scanf("%d", &function)
	switch function {
	case 1:
		{
			fmt.Printf("enter base:")
			fmt.Scanf("%f", &base)
			fmt.Printf("enter height:")
			fmt.Scanf("%f", &height)
			area = triangle(base, height)
			fmt.Println("triangle area : ", area)
		}
	case 2:
		{
			fmt.Printf("enter height:")
			fmt.Scanf("%f", &height)
			fmt.Printf("enter weight:")
			fmt.Scanf("%f", &weight)
			area = ractangle(height, weight)
			fmt.Println("ractangle area : ", area)
		}
	case 3:
		{
			fmt.Printf("enter radiant:")
			fmt.Scanf("%f", &radian)
			area = circle(radian)
			fmt.Println("circle area : ", area)
		}
	default:
		{
			fmt.Println("try again!")
		}
	}
}
func triangle(b, h float32) float32 {
	return 0.5 * b * h
}
func ractangle(h, w float32) float32 {
	return h * w
}
func circle(r float32) float32 {
	return pi * r * r
}
