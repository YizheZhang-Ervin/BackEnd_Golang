package main

func main() {
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(x int) {
			sendRPC(x)
			done <- true
		}(i)
	}
	for i := 0; i < 5; i++ {
		data := <-done
		println("Receive:", data)
	}
}

func sendRPC(i int) {
	println("Send:", i)
}
