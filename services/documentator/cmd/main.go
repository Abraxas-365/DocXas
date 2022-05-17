package main

import (
	"documentator/reader/core/models"
	"fmt"
)

func main() {
	documentation := models.Documentation{}
	documentation.New()
	fmt.Println(documentation)
}
