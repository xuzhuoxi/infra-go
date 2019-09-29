package filterx

import (
	"fmt"
	"testing"
)

func TestFilterKernelSort(t *testing.T) {
	kernel := Emboss3Asymmetrical.Kernel
	fmt.Println(kernel)
	filter2 := kernel.RotateClockwise90()
	fmt.Println(filter2)
	filter2.Sorted()
	fmt.Println(filter2)
}
