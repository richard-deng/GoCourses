package main 
import "fmt"
func main(){
	var slice1 = make([]int, 3, 6)
    slice2 := append(slice1, 1, 2, 3)	
	slice3 := append(slice1, 1, 2, 3, 4)
	fmt.Printf("len = %d cap = %d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len = %d cap = %d %v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len = %d cap = %d %v\n", len(slice3), cap(slice3), slice3)
}
