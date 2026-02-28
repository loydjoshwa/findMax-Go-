package main

import "fmt"

func findmax(arr []int) int {
	max:=arr[0]

	for i := 1; i < len(arr); i++ {
		if max<arr[i]{
			max=arr[i]
		}
		
	}
	return max
}
func main() {
	fmt.Println("find maximum in an array")
	array:=[]int{20,30,30,40,900,40,1}
	a:=findmax(array)
	fmt.Println("max:",a)
}
