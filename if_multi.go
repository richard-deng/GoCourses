package main 
import "fmt"
func main(){
	var grade int 
	fmt.Scanln(&grade)
	if(grade >= 90) && (grade <= 100){
		fmt.Println("A")
	} else if(grade >=80 ) && (grade <90){
		fmt.Println("B")
	} else if(grade >=70 ) && (grade <80){
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}
