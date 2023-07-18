// Package binaryx
// Created by xuzhuoxi
// on 2019-03-19.
// @author xuzhuoxi
//
package binaryx

import (
	"encoding/binary"
	"errors"
	"io"
)

// ReadLen
// 从Reader中读取一个长度数据
// 按uint16格式读取，强制转换为int值返回
// r: Reader
// order: 大小端设定
func ReadLen(r io.Reader, order binary.ByteOrder) (ln int, err error) {
	return ReadLen2(r, order)
}

// ReadLen1
// 从Reader中读取一个长度数据
// r: Reader
// order: 大小端设定
func ReadLen1(r io.Reader, order binary.ByteOrder) (ln int, err error) {
	val := uint8(0)
	if err := binary.Read(r, order, &val); nil != err {
		return 0, err
	}
	return int(val), nil
}

// ReadLen2
// 从Reader中读取一个长度数据
// r: Reader
// order: 大小端设定
func ReadLen2(r io.Reader, order binary.ByteOrder) (ln int, err error) {
	val := uint16(0)
	if err := binary.Read(r, order, &val); nil != err {
		return 0, err
	}
	return int(val), nil
}

// ReadLen4
// 从Reader中读取一个长度数据
// r: Reader
// order: 大小端设定
func ReadLen4(r io.Reader, order binary.ByteOrder) (ln uint, err error) {
	val := uint32(0)
	if err := binary.Read(r, order, &val); nil != err {
		return 0, err
	}
	return uint(val), nil
}

// ReadLenTo
// 从Reader中读取一个长度数据
// r: Reader
// order: 大小端设定
// ln: 长度值，只支持int8, int16, int32, int64, uint8, uint16, uint32, uint64及其指针类型
func ReadLenTo(w io.Reader, order binary.ByteOrder, ln interface{}) error {
	switch ln.(type) {
	case int8, int16, int32, int64, *int8, *int16, *int32, *int64:
		return binary.Read(w, order, ln)
	case uint8, uint16, uint32, uint64, *uint8, *uint16, *uint32, *uint64:
		return binary.Read(w, order, ln)
	}
	return errors.New("ln type error! ")
}

// WriteLen
// 把长度写入到Writer中
// w: Writer
// order: 大小端设定
// ln: 长度值，按uint16写入
func WriteLen(w io.Writer, order binary.ByteOrder, ln int) error {
	return WriteLen2(w, order, ln)
}

// WriteLen1
// 把长度写入到Writer中
// w: Writer
// order: 大小端设定
// ln: 长度值，按uint16写入
func WriteLen1(w io.Writer, order binary.ByteOrder, ln int) error {
	return binary.Write(w, order, uint8(ln))
}

// WriteLen2
// 把长度写入到Writer中
// w: Writer
// order: 大小端设定
// ln: 长度值，按uint16写入
func WriteLen2(w io.Writer, order binary.ByteOrder, ln int) error {
	return binary.Write(w, order, uint16(ln))
}

// WriteLen4
// 把长度写入到Writer中
// w: Writer
// order: 大小端设定
// ln: 长度值，按uint16写入
func WriteLen4(w io.Writer, order binary.ByteOrder, ln uint) error {
	return binary.Write(w, order, uint32(ln))
}

// WriteLenTo
// 把长度写入到Writer中
// w: Writer
// order: 大小端设定
// ln: 长度值，只支持int8, int16, int32, int64, uint8, uint16, uint32, uint64及其指针类型
func WriteLenTo(w io.Writer, order binary.ByteOrder, ln interface{}) error {
	switch ln.(type) {
	case int8, int16, int32, int64, *int8, *int16, *int32, *int64:
		return binary.Write(w, order, ln)
	case uint8, uint16, uint32, uint64, *uint8, *uint16, *uint32, *uint64:
		return binary.Write(w, order, ln)
	}
	return errors.New("ln type error! ")
}
