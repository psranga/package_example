package main

import (
	"fmt"

	"github.com/psranga/package_example/formatter"
	"github.com/psranga/package_example/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
