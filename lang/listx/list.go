package listx

// IList is a Java-style ordered list of interface{} values.
// IList 是面向 interface{} 的类 Java 有序列表接口。
type IList interface {
	// Len returns the number of elements in the list.
	// Len 返回列表中的元素个数。
	Len() int
	// Swap swaps the elements at indexes i and j.
	// Swap 交换索引 i 与 j 处的元素。
	Swap(i, j int)
	// Clear removes all elements from the list.
	// Clear 移除列表中的全部元素。
	Clear()

	// Add appends elements to the end of the list.
	// Add 在列表末尾追加元素。
	Add(ele ...interface{}) (suc bool)
	// AddAt inserts elements at the specified index.
	// AddAt 在指定索引处插入元素。
	AddAt(index int, ele ...interface{}) (suc bool)
	// AddAll inserts all elements from list starting at index.
	// AddAll 从 index 起插入 list 中的全部元素。
	AddAll(index int, list IList) (suc bool)

	// RemoveAt removes the element at the specified index.
	// RemoveAt 移除指定索引处的元素。
	RemoveAt(index int) (ele interface{}, suc bool)
	// RemoveMultiAt removes amount elements starting at index.
	// RemoveMultiAt 从 index 起移除 amount 个元素。
	RemoveMultiAt(index int, amount int) (ele []interface{}, suc bool)
	// RemoveLast removes the last element.
	// RemoveLast 移除最后一个元素。
	RemoveLast() (ele interface{}, suc bool)
	// RemoveLastMulti removes amount elements from the end.
	// RemoveLastMulti 从末尾起移除 amount 个元素。
	RemoveLastMulti(amount int) (ele []interface{}, suc bool)
	// RemoveFirst removes the first element.
	// RemoveFirst 移除第一个元素。
	RemoveFirst() (ele interface{}, suc bool)
	// RemoveFirstMulti removes amount elements from the front.
	// RemoveFirstMulti 从开头起移除 amount 个元素。
	RemoveFirstMulti(amount int) (ele []interface{}, suc bool)

	// Get returns the element at the specified index.
	// Get 返回指定索引处的元素。
	Get(index int) (ele interface{}, ok bool)
	// GetMulti returns amount elements starting at index.
	// GetMulti 从 index 起返回 amount 个元素。
	GetMulti(index int, amount int) (ele []interface{}, ok bool)
	// GetMultiLast returns amount elements ending at lastIndex.
	// GetMultiLast 以 lastIndex 为终点返回 amount 个元素。
	GetMultiLast(lastIndex int, amount int) (ele []interface{}, ok bool)
	// GetAll returns all elements in the list.
	// GetAll 返回列表中的全部元素。
	GetAll() []interface{}
	// First returns the first element.
	// First 返回第一个元素。
	First() (ele interface{}, ok bool)
	// FirstMulti returns amount elements from the front.
	// FirstMulti 从开头返回 amount 个元素。
	FirstMulti(amount int) (ele []interface{}, ok bool)
	// Last returns the last element.
	// Last 返回最后一个元素。
	Last() (ele interface{}, ok bool)
	// LastMulti returns amount elements from the end.
	// LastMulti 从末尾返回 amount 个元素。
	LastMulti(amount int) (ele []interface{}, ok bool)

	// IndexOf returns the index of the first matching element.
	// IndexOf 返回首个匹配元素的索引。
	IndexOf(ele interface{}) (index int, ok bool)
	// LastIndexOf returns the index of the last matching element.
	// LastIndexOf 返回最后一个匹配元素的索引。
	LastIndexOf(ele interface{}) (index int, ok bool)
	// Contains reports whether the list contains ele.
	// Contains 判断列表是否包含 ele。
	Contains(ele interface{}) (contains bool)
	// ForEach iterates elements from front to back until stop is true.
	// ForEach 从前向后遍历元素，stop 为 true 时停止。
	ForEach(each func(index int, ele interface{}) (stop bool))
	// ForEachLast iterates elements from back to front until stop is true.
	// ForEachLast 从后向前遍历元素，stop 为 true 时停止。
	ForEachLast(each func(index int, ele interface{}) (stop bool))
}
