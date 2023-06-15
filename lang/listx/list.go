// Package listx
// Created by xuzhuoxi
// on 2019-04-03.
// @author xuzhuoxi
//
package listx

type IList interface {
	// Len
	//the number of elements in the collection.
	Len() int
	// Swap
	// swaps the elements with indexes i and j.
	Swap(i, j int)
	// Clear
	// Removes all of the elements from this list (optional operation).
	Clear()

	// Add
	// Appends the elements to the end of this list (optional operation).
	Add(ele ...interface{}) (suc bool)
	// AddAt
	// Inserts the specified elements at the specified position in this list (optional operation).
	AddAt(index int, ele ...interface{}) (suc bool)
	// AddAll
	// Appends all of the elements in the specified collection to the end of this list, in the order that they are returned by the specified collection's iterator (optional operation)
	AddAll(index int, list IList) (suc bool)

	// RemoveAt
	// Removes the element at the specified position in this list (optional operation).
	RemoveAt(index int) (ele interface{}, suc bool)
	// RemoveMultiAt
	// Removes some elements at the specified position in this list (optional operation).
	RemoveMultiAt(index int, amount int) (ele []interface{}, suc bool)
	// RemoveLast
	// Removes the last element at from this list (optional operation).
	RemoveLast() (ele interface{}, suc bool)
	// RemoveLastMulti
	// Removes some elements from this list started at the last position.
	RemoveLastMulti(amount int) (ele []interface{}, suc bool)
	// RemoveFirst
	// Removes the first element at from this list (optional operation).
	RemoveFirst() (ele interface{}, suc bool)
	// RemoveFirstMulti
	// Removes some elements from this list started at the first position.
	RemoveFirstMulti(amount int) (ele []interface{}, suc bool)

	// Get
	// Returns the element at the specified position in this list.
	Get(index int) (ele interface{}, ok bool)
	// GetMulti
	// Returns some elements in this list started at the specified position.
	GetMulti(index int, amount int) (ele []interface{}, ok bool)
	// GetMultiLast
	// Returns some elements in this list started at the specified position.
	GetMultiLast(lastIndex int, amount int) (ele []interface{}, ok bool)
	// GetAll
	// Returns all elements in this list.
	GetAll() []interface{}
	// First
	// Returns the first element of this list.
	First() (ele interface{}, ok bool)
	// FirstMulti
	// Returns some elements in this list started at the first position.
	FirstMulti(amount int) (ele []interface{}, ok bool)
	// Last
	// Returns the last element of this list.
	Last() (ele interface{}, ok bool)
	// LastMulti
	// Returns some elements in this list started at the last position.
	LastMulti(amount int) (ele []interface{}, ok bool)

	// IndexOf
	// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
	IndexOf(ele interface{}) (index int, ok bool)
	// LastIndexOf
	// Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
	LastIndexOf(ele interface{}) (index int, ok bool)
	// Contains
	// Returns true if this list contains the specified element.
	Contains(ele interface{}) (contains bool)
	// ForEach
	// Performs the given action for each element of the list
	ForEach(each func(index int, ele interface{}) (stop bool))
	// ForEachLast
	// Performs the given action for each element of the list
	ForEachLast(each func(index int, ele interface{}) (stop bool))
}
