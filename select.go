package main
import "fmt"

func main(){
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
			fmt.Println("0")
		case ch <- 1:
			fmt.Println("1")
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}

}