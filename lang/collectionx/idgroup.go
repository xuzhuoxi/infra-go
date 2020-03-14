package collectionx

import "errors"

var (
	ErrSupportIdUnknown = errors.New("IdGroup: Id Unknown. ")
	ErrSupportIdExists  = errors.New("IdGroup: Id Exists. ")
	ErrSupportNil       = errors.New("IdGroup: IdSupport is nil. ")
)

// 唯一标识支持
type IIdSupport interface {
	Id() string
	SetId(Id string)
}

type IdSupport struct {
	id string
}

func (s *IdSupport) String() string {
	return s.id
}

func (s *IdSupport) Id() string {
	return s.id
}

func (s *IdSupport) SetId(Id string) {
	s.id = Id
}

type IIdGroup interface {
	// 数量
	Size() int
	// 元素列表
	Collection() []IIdSupport
	// id列表
	Ids() []string

	// 加入一个Id
	// err:
	//		Id重复时返回 ErrSupportIdExists
	Add(support IIdSupport) error
	// 加入多个Id
	// count: 成功加入的Id数量
	// err1:
	//		元素为nil时返回 ErrSupportNil
	// err2:
	//		Id重复时返回 ErrSupportIdExists
	Adds(supports []IIdSupport) (count int, failArr []IIdSupport, err error, err2 error)
	// 移除一个Id
	// support: 返回被移除的Id
	// err:
	//		Id不存在时返回 ErrSupportIdUnknown
	Remove(supportId string) (support IIdSupport, err error)
	// 移除多个Id
	// supports: 返回被移除的Id数组
	// err:
	//		Id不存在时返回 ErrSupportIdUnknown
	Removes(supportIdArr []string) (supports []IIdSupport, err error)
	// 替换一个Id
	// 根据Id进行替换，如果找不到相同Id，直接加入
	Update(support IIdSupport)
	// 替换一个Id
	// 根据Id进行替换，如果找不到相同Id，直接加入
	Updates(supports []IIdSupport)
}

type IdGroup struct {
	supports   []IIdSupport
	supportMap map[string]IIdSupport
}

func (g *IdGroup) Size() int {
	return len(g.supports)
}

func (g *IdGroup) Collection() []IIdSupport {
	return g.supports
}

func (g *IdGroup) Ids() []string {
	ln := len(g.supports)
	rs := make([]string, ln, ln)
	for index, q := range g.supports {
		rs[index] = q.Id()
	}
	return rs
}

func (g *IdGroup) Add(support IIdSupport) error {
	if g.isIdExists(support.Id()) {
		return ErrSupportIdExists
	}
	g.add(support)
	return nil

}

func (g *IdGroup) Adds(supports []IIdSupport) (count int, failArr []IIdSupport, err1 error, err2 error) {
	for _, support := range supports {
		if nil == support {
			err1 = ErrSupportNil
			continue
		}
		if g.isIdExists(support.Id()) {
			err2 = ErrSupportIdExists
			failArr = append(failArr, support)
			continue
		}
		g.add(support)
		count += 1
	}
	return
}

func (g *IdGroup) Remove(supportId string) (support IIdSupport, err error) {
	if !g.isIdExists(supportId) {
		return nil, ErrSupportIdUnknown
	}
	support = g.removeMapBy(supportId)
	index := g.findIndex(supportId)
	g.removeArrayAt(index)
	return
}

func (g *IdGroup) Removes(supportIdArr []string) (supports []IIdSupport, err error) {
	for _, sId := range supportIdArr {
		if !g.isIdExists(sId) {
			err = ErrSupportIdUnknown
			continue
		}
		queue := g.removeMapBy(sId)
		index := g.findIndex(sId)
		g.removeArrayAt(index)
		supports = append(supports, queue)
	}
	return
}

func (g *IdGroup) Update(support IIdSupport) {
	sId := support.Id()
	if g.isIdExists(sId) {
		index := g.findIndex(sId)
		g.supports[index] = support
	} else {
		g.supports = append(g.supports, support)
	}
	g.supportMap[sId] = support
}

func (g *IdGroup) Updates(supports []IIdSupport) {
	for idx, _ := range supports {
		if nil != supports[idx] {
			continue
		}
		g.Update(supports[idx])
	}
	return
}

func (g *IdGroup) add(support IIdSupport) {
	g.supports = append(g.supports, support)
	g.supportMap[support.Id()] = support
}

func (g *IdGroup) removeArrayAt(index int) (support IIdSupport) {
	if index >= 0 && index < len(g.supports) {
		support = g.supports[index]
		g.supports = append(g.supports[:index], g.supports[index+1:]...)
	}
	return
}

func (g *IdGroup) removeMapBy(supportId string) (support IIdSupport) {
	if q, ok := g.supportMap[supportId]; ok {
		delete(g.supportMap, supportId)
		return q
	} else {
		return nil
	}
}

func (g *IdGroup) findIndex(supportId string) int {
	for index, q := range g.supports {
		if q.Id() == supportId {
			return index
		}
	}
	return -1
}

func (g *IdGroup) isIdExists(supportId string) bool {
	if _, ok := g.supportMap[supportId]; ok {
		return true
	}
	return false
}
