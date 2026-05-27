package collectionx

import "errors"

var (
	// ErrElementIdUnknown is returned when the element id is not found.
	// ErrElementIdUnknown 在元素 id 不存在时返回。
	ErrElementIdUnknown = errors.New("OrderHash: ElementId Unknown. ")
	// ErrElementIdExists is returned when the element id already exists.
	// ErrElementIdExists 在元素 id 已存在时返回。
	ErrElementIdExists = errors.New("OrderHash: ElementId Exists. ")
	// ErrElementNil is returned when the element is nil.
	// ErrElementNil 在元素为 nil 时返回。
	ErrElementNil = errors.New("OrderHash: Element is nil. ")
	// ErrElementIndex is returned when the index is out of range.
	// ErrElementIndex 在索引越界时返回。
	ErrElementIndex = errors.New("OrderHash: Index out of range. ")
)

var Threshold = 1000 // threshold for switching lookup strategy / 查找策略切换阈值

// IOrderHashElement is an element with a unique string identifier.
// IOrderHashElement 是带唯一 string 标识的元素。
type IOrderHashElement interface {
	// Id returns the element identifier.
	// Id 返回元素标识。
	Id() string
	// SetId sets the element identifier.
	// SetId 设置元素标识。
	SetId(Id string)
}

// OrderHashElement is a default IOrderHashElement implementation.
// OrderHashElement 是 IOrderHashElement 的默认实现。
type OrderHashElement struct {
	id string
}

// String returns the element id as its string form.
// String 以 id 作为字符串表示。
func (s *OrderHashElement) String() string {
	return s.id
}

// Id returns the element identifier.
// Id 返回元素标识。
func (s *OrderHashElement) Id() string {
	return s.id
}

// SetId sets the element identifier.
// SetId 设置元素标识。
func (s *OrderHashElement) SetId(Id string) {
	s.id = Id
}

// IOrderHashGroup is an ordered collection of IOrderHashElement values keyed by Id.
// IOrderHashGroup 是按 Id 索引的有序 IOrderHashElement 集合。
type IOrderHashGroup interface {
	// Size returns the number of elements.
	// Size 返回元素数量。
	Size() int
	// Collection returns all elements in insertion order.
	// Collection 按插入顺序返回全部元素。
	Collection() []IOrderHashElement
	// Ids returns all element identifiers in order.
	// Ids 按顺序返回全部元素标识。
	Ids() []string

	// Get returns the element with the given id.
	// Get 按 id 获取元素。
	Get(eleId string) (ele IOrderHashElement, ok bool)
	// GetAt returns the element at the given index.
	// GetAt 按索引获取元素。
	GetAt(index int) (ele IOrderHashElement, ok bool)
	// Exists reports whether an element with the given id is present.
	// Exists 判断指定 id 的元素是否存在。
	Exists(eleId string) (ok bool)
	// Add inserts ele into the group.
	// Add 向分组中加入一个元素。
	// err: ErrElementNil, ErrElementIdExists
	Add(ele IOrderHashElement) error
	// AddAt inserts ele at the given index.
	// AddAt 在指定索引处加入元素。
	// err: ErrElementNil, ErrElementIndex, ErrElementIdExists
	AddAt(ele IOrderHashElement, index int) error
	// Adds inserts multiple elements.
	// Adds 批量加入多个元素。
	// count: number of elements successfully added.
	// count: 成功加入的元素数量。
	// err: errors produced while adding each element.
	// err: 每个元素加入时产生的错误。
	Adds(eles []IOrderHashElement) (count int, failArr []IOrderHashElement, err []error)
	// AddsAt inserts multiple elements starting at index.
	// AddsAt 从指定索引起批量加入元素。
	// count: number of elements successfully added.
	// count: 成功加入的元素数量。
	// err: errors produced while adding each element.
	// err: 每个元素加入时产生的错误。
	AddsAt(eles []IOrderHashElement, index int) (count int, failArr []IOrderHashElement, err []error)
	// Remove removes the element with the given id.
	// Remove 按 id 移除元素。
	// ele: the removed element.
	// ele: 被移除的元素。
	// err: ErrElementIdUnknown
	Remove(eleId string) (ele IOrderHashElement, err error)
	// RemoveAt removes the element at the given index.
	// RemoveAt 按索引移除元素。
	// ele: the removed element.
	// ele: 被移除的元素。
	// err: ErrElementIndex
	RemoveAt(index int) (ele IOrderHashElement, err error)
	// Removes removes multiple elements by id.
	// Removes 按 id 批量移除元素。
	// eles: removed elements.
	// eles: 被移除的元素列表。
	// err: ErrElementIdUnknown
	Removes(eleIdArr []string) (eles []IOrderHashElement, err []error)
	// RemovesAt removes count elements starting at index.
	// RemovesAt 从指定索引起移除 count 个元素。
	// eles: removed elements.
	// eles: 被移除的元素列表。
	// err: ErrElementIndex
	RemovesAt(index int, count int) (eles []IOrderHashElement, err error)
	// RemoveAll removes every element and returns them.
	// RemoveAll 移除全部元素并返回。
	// eles: all removed elements.
	// eles: 被移除的全部元素。
	RemoveAll() (eles []IOrderHashElement)
	// Update replaces the element with the same id, or adds it if missing.
	// Update 按 id 替换元素；不存在则直接加入。
	Update(ele IOrderHashElement) (replaced IOrderHashElement, err error)
	// Updates batch-updates elements by id.
	// Updates 按 id 批量替换元素。
	Updates(eles []IOrderHashElement) (replaced []IOrderHashElement, err []error)
	// ForEachElement iterates elements in order until f returns stop true.
	// ForEachElement 按顺序遍历元素，f 返回 stop 为 true 时停止。
	ForEachElement(f func(index int, ele IOrderHashElement) (stop bool))
}

// NewOrderHashGroup creates an empty OrderHashGroup.
// NewOrderHashGroup 创建一个空的 OrderHashGroup。
func NewOrderHashGroup() *OrderHashGroup {
	rs := &OrderHashGroup{}
	rs.eleMap = make(map[string]IOrderHashElement)
	rs.eles = make([]IOrderHashElement, 0, 32)
	return rs
}

// OrderHashGroup stores elements in slice order with a map for fast lookup.
// OrderHashGroup 以切片保持顺序，并以 map 加速查找。
type OrderHashGroup struct {
	eles   []IOrderHashElement
	eleMap map[string]IOrderHashElement
}

// Size returns the number of elements.
// Size 返回元素数量。
func (g *OrderHashGroup) Size() int {
	return len(g.eles)
}

// Collection returns all elements in insertion order.
// Collection 按插入顺序返回全部元素。
func (g *OrderHashGroup) Collection() []IOrderHashElement {
	return g.eles
}

// Ids returns all element identifiers in order.
// Ids 按顺序返回全部元素标识。
func (g *OrderHashGroup) Ids() []string {
	ln := len(g.eles)
	rs := make([]string, ln, ln)
	for index, q := range g.eles {
		rs[index] = q.Id()
	}
	return rs
}

// Get returns the element with the given id.
// Get 按 id 获取元素。
func (g *OrderHashGroup) Get(eleId string) (ele IOrderHashElement, ok bool) {
	ele, ok = g.exists(eleId)
	return
}

// GetAt returns the element at the given index.
// GetAt 按索引获取元素。
func (g *OrderHashGroup) GetAt(index int) (ele IOrderHashElement, ok bool) {
	if index < 0 || index >= g.Size() {
		return nil, false
	}
	return g.eles[index], ok
}

// Exists reports whether an element with the given id is present.
// Exists 判断指定 id 的元素是否存在。
func (g *OrderHashGroup) Exists(eleId string) (ok bool) {
	_, ok = g.exists(eleId)
	return
}

// Add inserts ele into the group.
// Add 向分组中加入一个元素。
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

// AddAt inserts ele at the given index.
// AddAt 在指定索引处加入元素。
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

// Adds inserts multiple elements.
// Adds 批量加入多个元素。
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

// AddsAt inserts multiple elements starting at index.
// AddsAt 从指定索引起批量加入元素。
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

// Remove removes the element with the given id.
// Remove 按 id 移除元素。
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

// RemoveAt removes the element at the given index.
// RemoveAt 按索引移除元素。
func (g *OrderHashGroup) RemoveAt(index int) (ele IOrderHashElement, err error) {
	if index < 0 || index >= len(g.eles) {
		return nil, ErrElementIndex
	}
	ele = g.removeAt(index)
	return
}

// Removes removes multiple elements by id.
// Removes 按 id 批量移除元素。
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

// RemovesAt removes count elements starting at index.
// RemovesAt 从指定索引起移除 count 个元素。
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

// RemoveAll removes every element and returns them.
// RemoveAll 移除全部元素并返回。
func (g *OrderHashGroup) RemoveAll() (eles []IOrderHashElement) {
	eles = g.eles
	g.eleMap = make(map[string]IOrderHashElement)
	g.eles = make([]IOrderHashElement, 0, 32)
	return
}

// Update replaces the element with the same id, or adds it if missing.
// Update 按 id 替换元素；不存在则直接加入。
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

// Updates batch-updates elements by id.
// Updates 按 id 批量替换元素。
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

// ForEachElement iterates elements in order until f returns stop true.
// ForEachElement 按顺序遍历元素，f 返回 stop 为 true 时停止。
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
