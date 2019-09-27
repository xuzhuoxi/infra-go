package gaussx

import (
	"fmt"
	"testing"
)

func TestGetGaussKernel(t *testing.T) {
	temp := CreateGaussKernel(2, 1.4)
	fmt.Println(temp)
	temp = CreateGaussKernel(1, 1)
	fmt.Println(temp)
}

func TestCreateGaussKernel2(t *testing.T) {
	temp := CreateGaussKernel2(2, 1.4)
	fmt.Println(temp)
	temp = CreateGaussKernel2(2, 1.0)
	fmt.Println(temp)
	temp = CreateGaussKernel2(1, 1)
	fmt.Println(temp)
}

func TestGetGaussKernelInteger(t *testing.T) {
	temp := CreateGaussKernelInt(2, 1, 340)
	fmt.Println(temp)
	temp = CreateGaussKernelInt(1, 1, 16)
	fmt.Println(temp)
}

func TestGetGaussKernelInteger2(t *testing.T) {
	temp := CreateGaussKernelInt2(2, 1, 300)
	fmt.Println(temp)
	temp = CreateGaussKernelInt2(1, 1, 16)
	fmt.Println(temp)
	temp = CreateGaussKernelInt2(1, 0.8, 0)
	fmt.Println(temp)
}

func TestGetAvgArr(t *testing.T) {
	temp := GetAvgArr(2, 1.4)
	fmt.Println(temp)
	temp = GetAvgArr(1, 1)
	fmt.Println(temp)
}
