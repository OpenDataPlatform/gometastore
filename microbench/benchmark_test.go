package microbench

import (
	"fmt"
)

func ExampleMeasureSimple() {
	MeasureSimple(func() {
		fmt.Print("hello ")
	}, 1, 2)
	// Output: hello hello hello
}

func ExampleMeasure() {
	Measure(
		func() {
			fmt.Print(" pretest")
		},
		func() {
			fmt.Print(" test")
		},
		func() {
			fmt.Print(" cleanup")
		}, 1, 1)
	// Output: pretest test cleanup pretest test cleanup
}

func ExampleStats_Max() {
	stats := MakeStats()
	stats.Add(1.0)
	stats.Add(2.0)
	fmt.Println("max =", stats.Max())
	// Output: max = 2
}

func ExampleStats_Min() {
	stats := MakeStats()
	stats.Add(1.0)
	stats.Add(2.0)
	fmt.Println("min =", stats.Min())
	// Output: min = 1
}
