package filterx

import (
	"fmt"
	"testing"
)

func TestFilterKernelSort(t *testing.T) {
	var kernel FilterKernel
	kernel = Emboss3Asymmetrical.Kernel
	fmt.Println(kernel)
	kernel = kernel.Rotate90(true)
	kernel.Sorted()
	fmt.Println(kernel)
	kernel = kernel.Rotate90(true)
	kernel.Sorted()
	fmt.Println(kernel)
	kernel = kernel.Rotate90(true)
	kernel.Sorted()
	fmt.Println(kernel)
	fmt.Println("--------------------------")
	kernel = Emboss3Asymmetrical.Kernel
	kernel = kernel.Rotate90(false)
	kernel.Sorted()
	fmt.Println(kernel)
	kernel = kernel.Rotate90(false)
	kernel.Sorted()
	fmt.Println(kernel)
	kernel = kernel.Rotate90(false)
	kernel.Sorted()
	fmt.Println(kernel)
}

func TestFor(t *testing.T) {
	s := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for index := range s {
		fmt.Println(index)
	}
}
