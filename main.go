package main

import (
	"fmt"
	"os"
	"strings"

	echo "./echo"
)

func main() {
	fmt.Println(echo.Echo(strings.Join(os.Args[1:], " ")))
}
