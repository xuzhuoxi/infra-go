package collectionx

import "errors"

var (
	ErrSupportIdUnknown = errors.New("OrderHash: Id Unknown. ")
	ErrSupportIdExists  = errors.New("OrderHash: Id Exists. ")
	ErrSupportNil       = errors.New("OrderHash: Support is nil. ")
	ErrSupportIndex     = errors.New("OrderHash: Index out of range. ")
)

var Threshold = 1000 //阀值

// 唯一标识支持
type IOrderHashSupport interface {
	Id() string
	SetId(Id string)
}

type OrderHashSupport struct {
	id string
}

func (s *OrderHashSupport) String() string {
	return s.id
}

func (s *OrderHashSupport) Id() string {
	return s.id
}

func (s *OrderHashSupport) SetId(Id string) {
	s.id = Id
}

type IOrderHashGroup interface {
	// 数量
	Size() int
	// 元素列表
	Collection() []IOrderHashSupport
	// id列表
	Ids() []string

	// 取一个
	Get(supportId string) (support IOrderHashSupport, ok bool)
	// 判断
	Exists(supportId string) (ok bool)
	// 加入一个Id
	// err:
	//		ErrSupportNil,ErrSupportIdExists
	Add(support IOrderHashSupport) error
	// 加入一个Id
	// err:
	//		ErrSupportNil,ErrSupportIndex,ErrSupportIdExists
	AddAt(support IOrderHashSupport, index int) error
	// 加入多个Id
	// count: 成功加入的Id数量
	// err:
	//		每个加入时产生的错误
	Adds(supports []IOrderHashSupport) (count int, failArr []IOrderHashSupport, err []error)
	// 加入多个Id
	// count: 成功加入的Id数量
	// err:
	//		每个加入时产生的错误
	AddsAt(supports []IOrderHashSupport, index int) (count int, failArr []IOrderHashSupport, err []error)
	// 移除一个Id
	// support: 返回被移除的Id
	// err:
	//		ErrSupportIdUnknown
	Remove(supportId string) (support IOrderHashSupport, err error)
	// 移除一个Id
	// support: 返回被移除的Id
	// err:
	//		ErrSupportIndex
	RemoveAt(index int) (support IOrderHashSupport, err error)
	// 移除多个Id
	// supports: 返回被移除的Id数组
	// err:
	//		ErrSupportIdUnknown
	Removes(supportIdArr []string) (supports []IOrderHashSupport, err []error)
	// 移除多个Id
	// supports: 返回被移除的Id数组
	// err:
	//		ErrSupportIndex
	RemovesAt(index int, count int) (supports []IOrderHashSupport, err error)
	// 替换一个Id
	// 根据Id进行替换，如果找不到相同Id，直接加入
	Update(support IOrderHashSupport) (err error)
	// 替换一个Id
	// 根据Id进行替换，如果找不到相同Id，直接加入
	Updates(supports []IOrderHashSupport) (err []error)
}

type OrderHashGroup struct {
	supports   []IOrderHashSupport
	supportMap map[string]IOrderHashSupport
}

func (g *OrderHashGroup) Size() int {
	return len(g.supports)
}

func (g *OrderHashGroup) Collection() []IOrderHashSupport {
	return g.supports
}

func (g *OrderHashGroup) Ids() []string {
	ln := len(g.supports)
	rs := make([]string, ln, ln)
	for index, q := range g.supports {
		rs[index] = q.Id()
	}
	return rs
}

func (g *OrderHashGroup) Get(supportId string) (support IOrderHashSupport, ok bool) {
	support, ok = g.exists(supportId)
	return
}

func (g *OrderHashGroup) Exists(supportId string) (ok bool) {
	_, ok = g.exists(supportId)
	return
}

func (g *OrderHashGroup) Add(support IOrderHashSupport) error {
	if nil == support {
		return ErrSupportNil
	}
	_, ok := g.exists(support.Id())
	if ok {
		return ErrSupportIdExists
	}
	g.add(support)
	return nil

}

func (g *OrderHashGroup) AddAt(support IOrderHashSupport, index int) error {
	if nil == support {
		return ErrSupportNil
	}
	if index < 0 || index >= len(g.supports) {
		return ErrSupportIndex
	}
	_, ok := g.exists(support.Id())
	if ok {
		return ErrSupportIdExists
	}
	g.addAt(support, index)
	return nil
}

func (g *OrderHashGroup) Adds(supports []IOrderHashSupport) (count int, failArr []IOrderHashSupport, err []error) {
	for idx, _ := range supports {
		e := g.Add(supports[idx])
		if nil != e {
			err = append(err, e)
			failArr = append(failArr, supports[idx])
		} else {
			count += 1
		}
	}
	return
}

func (g *OrderHashGroup) AddsAt(supports []IOrderHashSupport, index int) (count int, failArr []IOrderHashSupport, err []error) {
	for idx, _ := range supports {
		e := g.AddAt(supports[idx], index)
		if nil != e {
			err = append(err, e)
			failArr = append(failArr, supports[idx])
		} else {
			count += 1
			index += 1
		}
	}
	return
}

func (g *OrderHashGroup) Remove(supportId string) (support IOrderHashSupport, err error) {
	support, ok := g.exists(supportId)
	if !ok {
		return nil, ErrSupportIdUnknown
	}
	support = g.removeMapBy(supportId)
	_, index := g.findIndex(supportId)
	g.removeArrayAt(index)
	return
}

func (g *OrderHashGroup) RemoveAt(index int) (support IOrderHashSupport, err error) {
	if index < 0 || index >= len(g.supports) {
		return nil, ErrSupportIndex
	}
	support = g.removeAt(index)
	return
}

func (g *OrderHashGroup) Removes(supportIdArr []string) (supports []IOrderHashSupport, err []error) {
	for _, sId := range supportIdArr {
		support, e := g.Remove(sId)
		if nil != e {
			err = append(err, e)
			continue
		}
		supports = append(supports, support)
	}
	return
}

func (g *OrderHashGroup) RemovesAt(index int, count int) (supports []IOrderHashSupport, err error) {
	if index < 0 || index >= len(g.supports) {
		return nil, ErrSupportIndex
	}
	for count >= 0 {
		support, e := g.RemoveAt(index)
		if e != nil {
			err = e
			return
		}
		supports = append(supports, support)
		count -= 1
	}
	return
}

func (g *OrderHashGroup) Update(support IOrderHashSupport) (err error) {
	if nil == support {
		return ErrSupportNil
	}
	sId := support.Id()
	_, index := g.findIndex(sId)
	if -1 != index {
		g.supports[index] = support
	} else {
		g.supports = append(g.supports, support)
	}
	g.supportMap[sId] = support
	return
}

func (g *OrderHashGroup) Updates(supports []IOrderHashSupport) (err []error) {
	for idx, _ := range supports {
		e := g.Update(supports[idx])
		if nil != e {
			err = append(err, e)
		}
	}
	return
}

//-------------------

func (g *OrderHashGroup) add(support IOrderHashSupport) {
	g.supports = append(g.supports, support)
	g.supportMap[support.Id()] = support
}

func (g *OrderHashGroup) addAt(support IOrderHashSupport, index int) {
	g.supports = append(g.supports, support)
	g.supportMap[support.Id()] = support
}

func (g *OrderHashGroup) remove(supportId string) (support IOrderHashSupport) {
	support = g.removeMapBy(supportId)
	if nil != support {
		_, index := g.findIndex(supportId)
		g.removeArrayAt(index)
	}
	return support
}

func (g *OrderHashGroup) removeAt(index int) (support IOrderHashSupport) {
	support = g.removeArrayAt(index)
	if nil != support {
		g.removeMapBy(support.Id())
	}
	return
}

func (g *OrderHashGroup) removeArrayAt(index int) (support IOrderHashSupport) {
	if index >= 0 && index < len(g.supports) {
		support = g.supports[index]
		g.supports = append(g.supports[:index], g.supports[index+1:]...)
	}
	return
}

func (g *OrderHashGroup) removeMapBy(supportId string) (support IOrderHashSupport) {
	if q, ok := g.supportMap[supportId]; ok {
		delete(g.supportMap, supportId)
		return q
	} else {
		return nil
	}
}

func (g *OrderHashGroup) exists(supportId string) (support IOrderHashSupport, ok bool) {
	if len(g.supports) <= Threshold {
		s, index := g.findIndex(supportId)
		return s, -1 != index
	} else {
		support, ok = g.supportMap[supportId]
		return
	}
}

func (g *OrderHashGroup) findIndex(supportId string) (support IOrderHashSupport, index int) {
	for index, q := range g.supports {
		if q.Id() == supportId {
			return q, index
		}
	}
	return nil, -1
}
