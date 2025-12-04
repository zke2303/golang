package main
import "fmt"
func main(){
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2
	ch3 <- 3
		select {
		case n, ok := <- ch1:
			fmt.Println(n, ok)
		case n, ok := <- ch2:
			fmt.Println(n, ok)
		case n, ok := <- ch3:
			fmt.Println(n, ok)
		default:
			fmt.Println("default")
		}

}
