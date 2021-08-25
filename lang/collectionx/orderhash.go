package collectionx

import "errors"

var (
	ErrElementIdUnknown = errors.New("OrderHash: ElementId Unknown. ")
	ErrElementIdExists  = errors.New("OrderHash: ElementId Exists. ")
	ErrElementNil       = errors.New("OrderHash: Element is nil. ")
	ErrElementIndex     = errors.New("OrderHash: Index out of range. ")
)

var Threshold = 1000 //阀值

// 唯一标识支持
type IOrderHashElement interface {
	Id() string
	SetId(Id string)
}

type OrderHashElement struct {
	id string
}

func (s *OrderHashElement) String() string {
	return s.id
}

func (s *OrderHashElement) Id() string {
	return s.id
}

func (s *OrderHashElement) SetId(Id string) {
	s.id = Id
}

type IOrderHashGroup interface {
	// 数量
	Size() int
	// 元素列表
	Collection() []IOrderHashElement
	// id列表
	Ids() []string

	// 取一个Element
	Get(eleId string) (ele IOrderHashElement, ok bool)
	// 取一个Element
	GetAt(index int) (ele IOrderHashElement, ok bool)
	// 判断
	Exists(eleId string) (ok bool)
	// 加入一个Element
	// err:
	//		ErrElementNil,ErrElementIdExists
	Add(ele IOrderHashElement) error
	// 加入一个Element
	// err:
	//		ErrElementNil,ErrElementIndex,ErrElementIdExists
	AddAt(ele IOrderHashElement, index int) error
	// 加入多个Element
	// count: 成功加入的Element数量
	// err:
	//		每个加入时产生的错误
	Adds(eles []IOrderHashElement) (count int, failArr []IOrderHashElement, err []error)
	// 加入多个Element
	// count: 成功加入的Element数量
	// err:
	//		每个加入时产生的错误
	AddsAt(eles []IOrderHashElement, index int) (count int, failArr []IOrderHashElement, err []error)
	// 移除一个Element
	// ele: 返回被移除的Element
	// err:
	//		ErrElementIdUnknown
	Remove(eleId string) (ele IOrderHashElement, err error)
	// 移除一个Element
	// ele: 返回被移除的Element
	// err:
	//		ErrElementIndex
	RemoveAt(index int) (ele IOrderHashElement, err error)
	// 移除多个Element
	// eles: 返回被移除的Element数组
	// err:
	//		ErrElementIdUnknown
	Removes(eleIdArr []string) (eles []IOrderHashElement, err []error)
	// 移除多个Element
	// eles: 返回被移除的Element数组
	// err:
	//		ErrElementIndex
	RemovesAt(index int, count int) (eles []IOrderHashElement, err error)
	// 移除全部Element
	// eles: 返回被移除的Element数组
	RemoveAll() (eles []IOrderHashElement)
	// 替换一个Element
	// 根据Id进行替换，如果找不到相同Id，直接加入
	Update(ele IOrderHashElement) (replaced IOrderHashElement, err error)
	// 替换一个Element
	// 根据Id进行替换，如果找不到相同Id，直接加入
	Updates(eles []IOrderHashElement) (replaced []IOrderHashElement, err []error)
	// 遍历元素
	ForEachElement(f func(index int, ele IOrderHashElement) (stop bool))
}

func NewOrderHashGroup() *OrderHashGroup {
	rs := &OrderHashGroup{}
	rs.eleMap = make(map[string]IOrderHashElement)
	rs.eles = make([]IOrderHashElement, 0, 32)
	return rs
}

type OrderHashGroup struct {
	eles   []IOrderHashElement
	eleMap map[string]IOrderHashElement
}

func (g *OrderHashGroup) Size() int {
	return len(g.eles)
}

func (g *OrderHashGroup) Collection() []IOrderHashElement {
	return g.eles
}

func (g *OrderHashGroup) Ids() []string {
	ln := len(g.eles)
	rs := make([]string, ln, ln)
	for index, q := range g.eles {
		rs[index] = q.Id()
	}
	return rs
}

func (g *OrderHashGroup) Get(eleId string) (ele IOrderHashElement, ok bool) {
	ele, ok = g.exists(eleId)
	return
}

func (g *OrderHashGroup) GetAt(index int) (ele IOrderHashElement, ok bool) {
	if index < 0 || index >= g.Size() {
		return nil, false
	}
	return g.eles[index], ok
}

func (g *OrderHashGroup) Exists(eleId string) (ok bool) {
	_, ok = g.exists(eleId)
	return
}

func (g *OrderHashGroup) Add(ele IOrderHashElement) error {
	if nil == ele {
		return ErrElementNil
	}
	_, ok := g.exists(ele.Id())
	if ok {
		return ErrElementIdExists
	}
	g.add(ele)
	return nil

}

func (g *OrderHashGroup) AddAt(ele IOrderHashElement, index int) error {
	if nil == ele {
		return ErrElementNil
	}
	if index < 0 || index >= len(g.eles) {
		return ErrElementIndex
	}
	_, ok := g.exists(ele.Id())
	if ok {
		return ErrElementIdExists
	}
	g.addAt(ele, index)
	return nil
}

func (g *OrderHashGroup) Adds(eles []IOrderHashElement) (count int, failArr []IOrderHashElement, err []error) {
	for idx := range eles {
		e := g.Add(eles[idx])
		if nil != e {
			err = append(err, e)
			failArr = append(failArr, eles[idx])
		} else {
			count += 1
		}
	}
	return
}

func (g *OrderHashGroup) AddsAt(eles []IOrderHashElement, index int) (count int, failArr []IOrderHashElement, err []error) {
	for idx := range eles {
		e := g.AddAt(eles[idx], index)
		if nil != e {
			err = append(err, e)
			failArr = append(failArr, eles[idx])
		} else {
			count += 1
			index += 1
		}
	}
	return
}

func (g *OrderHashGroup) Remove(eleId string) (ele IOrderHashElement, err error) {
	ele, ok := g.exists(eleId)
	if !ok {
		return nil, ErrElementIdUnknown
	}
	ele = g.removeMapBy(eleId)
	_, index := g.findIndex(eleId)
	g.removeArrayAt(index)
	return
}

func (g *OrderHashGroup) RemoveAt(index int) (ele IOrderHashElement, err error) {
	if index < 0 || index >= len(g.eles) {
		return nil, ErrElementIndex
	}
	ele = g.removeAt(index)
	return
}

func (g *OrderHashGroup) Removes(eleIdArr []string) (eles []IOrderHashElement, err []error) {
	for _, sId := range eleIdArr {
		ele, e := g.Remove(sId)
		if nil != e {
			err = append(err, e)
			continue
		}
		eles = append(eles, ele)
	}
	return
}

func (g *OrderHashGroup) RemovesAt(index int, count int) (eles []IOrderHashElement, err error) {
	if index < 0 || index >= len(g.eles) {
		return nil, ErrElementIndex
	}
	for count >= 0 {
		ele, e := g.RemoveAt(index)
		if e != nil {
			err = e
			return
		}
		eles = append(eles, ele)
		count -= 1
	}
	return
}

func (g *OrderHashGroup) RemoveAll() (eles []IOrderHashElement) {
	eles = g.eles
	g.eleMap = make(map[string]IOrderHashElement)
	g.eles = make([]IOrderHashElement, 0, 32)
	return
}

func (g *OrderHashGroup) Update(ele IOrderHashElement) (replaced IOrderHashElement, err error) {
	if nil == ele {
		return nil, ErrElementNil
	}
	sId := ele.Id()
	_, index := g.findIndex(sId)
	if -1 != index {
		replaced = g.eles[index]
		g.eles[index] = ele
	} else {
		g.eles = append(g.eles, ele)
	}
	g.eleMap[sId] = ele
	return
}

func (g *OrderHashGroup) Updates(eles []IOrderHashElement) (replaced []IOrderHashElement, err []error) {
	for idx := range eles {
		r, e := g.Update(eles[idx])
		if nil != e {
			err = append(err, e)
		} else {
			replaced = append(replaced, r)
		}
	}
	return
}

func (g *OrderHashGroup) ForEachElement(f func(index int, ele IOrderHashElement) (stop bool)) {
	for idx := range g.eles {
		if f(idx, g.eles[idx]) {
			return
		}
	}
}

//-------------------

func (g *OrderHashGroup) add(ele IOrderHashElement) {
	g.eles = append(g.eles, ele)
	g.eleMap[ele.Id()] = ele
}

func (g *OrderHashGroup) addAt(ele IOrderHashElement, index int) {
	g.eles = append(g.eles, ele)
	g.eleMap[ele.Id()] = ele
}

func (g *OrderHashGroup) remove(eleId string) (ele IOrderHashElement) {
	ele = g.removeMapBy(eleId)
	if nil != ele {
		_, index := g.findIndex(eleId)
		g.removeArrayAt(index)
	}
	return ele
}

func (g *OrderHashGroup) removeAt(index int) (ele IOrderHashElement) {
	ele = g.removeArrayAt(index)
	if nil != ele {
		g.removeMapBy(ele.Id())
	}
	return
}

func (g *OrderHashGroup) removeArrayAt(index int) (ele IOrderHashElement) {
	if index >= 0 && index < len(g.eles) {
		ele = g.eles[index]
		g.eles = append(g.eles[:index], g.eles[index+1:]...)
	}
	return
}

func (g *OrderHashGroup) removeMapBy(eleId string) (ele IOrderHashElement) {
	if q, ok := g.eleMap[eleId]; ok {
		delete(g.eleMap, eleId)
		return q
	} else {
		return nil
	}
}

func (g *OrderHashGroup) exists(eleId string) (ele IOrderHashElement, ok bool) {
	if len(g.eles) <= Threshold {
		s, index := g.findIndex(eleId)
		return s, -1 != index
	} else {
		ele, ok = g.eleMap[eleId]
		return
	}
}

func (g *OrderHashGroup) findIndex(eleId string) (ele IOrderHashElement, index int) {
	for index, q := range g.eles {
		if q.Id() == eleId {
			return q, index
		}
	}
	return nil, -1
}
