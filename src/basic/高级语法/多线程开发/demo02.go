package main

func main(){
	var ch chan int
	defer close(ch)
}
