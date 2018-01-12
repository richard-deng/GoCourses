package main
import (
    "fmt"
    "time"
)
func main(){
	t := time.Now()
	fmt.Printf(t.Format("2006年01月02日"))
	switch t.Weekday().String() {
	case "Sunday":
		fmt.Printf("星期天")
	case "Monday":
		fmt.Printf("星期一")
	case "Tuesday":
		fmt.Printf("星期二")
	case "Wednesday":
		fmt.Printf("星期三")
	case "Thrsday":
		fmt.Printf("星期四")
	case "Friday":
		fmt.Printf("星期五")
	case "Saturday":
		fmt.Printf("星期六")
	}
}
