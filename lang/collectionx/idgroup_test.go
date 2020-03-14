package collectionx

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var index = 0

func newIdSupport() IIdSupport {
	now := time.Now().Nanosecond()
	return &IdSupport{id: strconv.Itoa(now)}
}

func newIdSupport2() IIdSupport {
	index++
	return &IdSupport{id: strconv.Itoa(index)}
}

func TestIdGroup_Collection(t *testing.T) {
	var list []IIdSupport
	group := IdGroup{supports: nil, supportMap: make(map[string]IIdSupport)}
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
	fmt.Println(group.supportMap)

}
