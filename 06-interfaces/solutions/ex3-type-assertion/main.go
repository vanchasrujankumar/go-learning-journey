package main

import "fmt"

func describeType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("%d is an integer\n", val)
	case float64:
		fmt.Printf("%.2f is a float64\n", val)
	case string:
		fmt.Printf("%s is a string of length %d\n", val, len(val))
	case bool:
		fmt.Printf("%v is a boolean\n", val)
	case []int:
		fmt.Printf("%v is a slice of ints with length %d\n", val, len(val))
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	values := []interface{}{42, 3.14, "hello", true, []int{1, 2, 3}, 3.14}
	for _, v := range values {
		describeType(v)
	}
}
