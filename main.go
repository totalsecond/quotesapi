package main

import (
	"fmt"
	"github.com/totalsecond/quotesapi/router"
)

func main() {
	fmt.Println("Starting the quotes internet machine")

	router.Start()
}
