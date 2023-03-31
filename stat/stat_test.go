package stat

import "fmt"

type Group int

func ExampleMax() {
	ivals := []int{23, 9, 4, 42, 16, 15}
	fmt.Println(Max(ivals))
	fmt.Println(Max[int](nil))
	fvals := []float64{23, 9, 4, 42, 16, 15}
	fmt.Println(Max(fvals))
	gvals := []Group{23, 9, 4, 42, 16, 15}
	fmt.Println(Max(gvals))
	// Output:
	// 42 <nil>
	// 0 max of empty slice
	// 42 <nil>
	// 42 <nil>
}
