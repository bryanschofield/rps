package main

import (
	"fmt"
	"os"
	"strings"

	echo "./echo"
)

func main() {
	echoer := echo.NewEchoer()
	fmt.Println(echoer.Echo(strings.Join(os.Args[1:], " ")))
}
