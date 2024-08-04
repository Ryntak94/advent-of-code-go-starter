package main

import (
	"fmt"

	"github.com/ryntak94/advent-of-code-go-starter/test"
	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func main() {
	file, err := utils.FileAsString("other.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}

	println(test.AssertTestInput(file, partOne, file))

}

func partOne(file string) string {
	return file
}
