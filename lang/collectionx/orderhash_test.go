package collectionx

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var index = 0

func newIdSupport() IOrderHashElement {
	now := time.Now().Nanosecond()
	return &OrderHashElement{id: strconv.Itoa(now)}
}

func newIdSupport2() IOrderHashElement {
	index++
	return &OrderHashElement{id: strconv.Itoa(index)}
}

func TestIdGroup_Collection(t *testing.T) {
	var list []IOrderHashElement
	group := OrderHashGroup{eles: nil, eleMap: make(map[string]IOrderHashElement)}
	for i := 0; i < 20; i++ {
		support := newIdSupport2()
		group.add(support)
		list = append(list, support)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println(group.Collection())
	group.Remove("10")
	group.Removes([]string{"1", "5", "15", "20"})
	fmt.Println(group.Collection())
	fmt.Println(group.eleMap)

}
