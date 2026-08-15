package main

import (
	"fmt"

	"github.com/emcodest/go-core-basics/structs"

	"github.com/emcodest/go-core-basics/interfaces"
)

func main() {

	// interfaces & structs
	// rectangle := structs.Rectangle{}
	circle := structs.Circle{}
	res := interfaces.Calculate(circle)
	fmt.Println("Result: ", res.AreaOfShape, res.PerimeterOfShape)
	// enums
	//fmt.Println("result:", enums.FAILED)
	//fmt.Println(reflect.TypeOf(interfaces.MyVal))
}
