package main 
import "fmt"
func main(){
	var slice1 = []int{1,2,3,4,5,6,7,8,9,10}
	var slice2 = make([]int, 3, 5)
	var n int 
	n = copy(slice2, slice1)
	fmt.Println(n, slice2, len(slice2), cap(slice2))
	slice3 := slice1[3:6]
	n = copy(slice3, slice1[1:5])
	fmt.Println(n, slice1, slice3)
}
