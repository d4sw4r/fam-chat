package main

func main() {

	server := NewServer(":110")
	server.Run()
}
