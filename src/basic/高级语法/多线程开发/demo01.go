package main
import "fmt"
func main(){
	sh := make(chan int)
	defer close(sh)

	go func(){
		sh <- 10
	}()

	// 从无缓存管道中读取数据
	num := <- sh
	fmt.Println(num)

}
