package main

import (
	"fmt"

	"github.com/psranga/package_example/liba/formatter"
	"github.com/psranga/package_example/liba/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
