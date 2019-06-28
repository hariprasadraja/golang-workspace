package main

func PipeLine() <-chan int {
	newChan := make(chan int)
	go func() {
		/* Do something here and put it in the channel */
		// newChan <- ANYTHING
		close(newChan)
	}()

	return newChan
}

func main() {

	for result := range PipeLine() {
		${0}
	}

}
