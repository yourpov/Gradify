package main

import (
	"fmt"
	"github.com/yourpov/gradify"
)

func main() {
	fmt.Println(gradify.Gradient("x o x o | Error Preset!   | x o x o", gradify.Error)) // gradify.<preset name>
	fmt.Println(gradify.Gradient("x o x o | Success Preset! | x o x o", gradify.Success))
	fmt.Println(gradify.Gradient("x o x o | Warning Preset! | x o x o", gradify.Warning))
	fmt.Println(gradify.Gradient("x o x o | Info Preset!    | x o x o", gradify.Info))
	fmt.Println(gradify.Gradient("x o x o | Candy Preset!   | x o x o", gradify.Candy))
	fmt.Println(gradify.Gradient("x o x o | Minty Preset!   | x o x o", gradify.Minty))
}
