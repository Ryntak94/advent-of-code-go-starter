package main

import (
	"fmt"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func main() {
	file, err := utils.FileAsString("calibration.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	println(file)
}
