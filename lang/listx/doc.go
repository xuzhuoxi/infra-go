// Package listx provides IList and slice- or linked-list-backed list implementations.
// Package listx 提供 IList 接口及基于切片、链表的两类列表实现。
//
// Created by xuzhuoxi on 2019-04-03.
// @author xuzhuoxi
//
// Capabilities:
//   - IList: Java-style list operations (add, remove, get, index, foreach) on interface{}
//   - ArrayList: dynamic slice with optional maxSize and initial capacity
//   - LinkedList: doubly-linked list built on container/list with index-based access
//
// 主要职责：
//   - IList：面向 interface{} 的类 Java 列表操作（增删查、索引、遍历）
//   - ArrayList：可配置最大长度与初始容量的动态切片实现
//   - LinkedList：基于 container/list 的双向链表实现，支持按索引访问
package listx
