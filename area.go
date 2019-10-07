package main

import (
	"fmt"
)

func main() {
	var width, hight int
	fmt.Scanf("%d %d", &width, &hight)

	fmt.Println("area:", width*hight)
}