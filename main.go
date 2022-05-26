package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()

	massege := flag.Arg(0)

	massege = strings.ReplaceAll(massege, ",", "")
	massege = strings.ReplaceAll(massege, "(", "")
	massege = strings.ReplaceAll(massege, ")", "")
	massege = strings.ReplaceAll(massege, "-", "")
	massege = strings.ReplaceAll(massege, " ", "")
	fmt.Println(massege)

}
