package main
import "fmt"
func main(){
	mySlice:=[10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(mySlice[:3])
	fmt.Println(mySlice[4:])
}