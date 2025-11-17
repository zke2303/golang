package main

type BilError struct {
	msg string
}

func (e *BilError) Error() string {
	return e.msg
}

func main() {

}
