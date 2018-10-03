package main

import "os"

func main() {
	panic("A problem occured!")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
