package main

import (
	"fmt"

	TEngine ".."
)

func main() {
	_, err := TEngine.NewApp()
	if err != nil {
		fmt.Println(err)
		return
	}
}
