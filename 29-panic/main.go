package main

import "fmt"

const (
	manualPanic  = 1
	runtimePanic = 2
)

func badFunction() {
	fmt.Printf("Select Panic type (0=no panic, 1=int, 2=runtime panic)\n")

	var choice int

	fmt.Scanf("%d", &choice)

	switch choice {
	case manualPanic:
		panic(0)
	case runtimePanic:
		// The following code will panic
		var invalid func(int) int

		invalid(0)
	}
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
			default:
				panic(x)
			case int:
				fmt.Printf("Function panicked with a very unhelpful error: %d\n", x)
			}
		}
	}()
	badFunction()
	fmt.Printf("Program exited normally\n")
}
