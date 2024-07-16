package main

import (
	"fmt"

	"github.com/robilsharma38/coingecko-cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}

}
