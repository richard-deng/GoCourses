package main 
import "fmt"
func main(){
	var slice = []int{1,2,3,4,5}
	for i:=0; i<=4; i++{
		fmt.Printf("slice[%d]=%d ", i, slice[i])
	}
	fmt.Printf("\n")
	for i, v := range slice {
		fmt.Printf("slice[%d]=%d ", i, v)
	}
}
