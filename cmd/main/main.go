package main

import (
	"encoding/json"
	"fmt"
)

type Output struct {
	Decoration     bool
	DurationSecond int
	Message        string
}

func main() {
	for i, v := range []string{"1", "2", "3", "4", "5"} {
		fmt.Printf("Hello World! %v %v\n", i, v)
	}

	output := Output{}
	output.Message = "Happy Birthday!"
	output.Decoration = true
	output.DurationSecond = 120

	fmt.Printf("Output: %v\n", output)
	fmt.Printf("Output: %+v\n", output)
	fmt.Printf("Output: %T\n", output)

	bytes, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
