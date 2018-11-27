package gds

import (
	"fmt"
)

func Example_subName() {
	fmt.Println(subName("D_aa/bb/cc/dd", 4))
	fmt.Println(subName("D_aa", 2))

	// Output:
	// bb
	// aa
}
