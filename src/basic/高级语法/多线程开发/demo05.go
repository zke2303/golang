package main
import "fmt"
import "time"
func main(){
	go func(){
		fmt.Println(1)
	}()

	fmt.Println(2)
	time.Sleep(1)
}

