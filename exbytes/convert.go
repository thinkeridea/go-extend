// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package exbytes

import "unsafe"

// ToString 把 []byte 转换为 string 没有多余的内存开销。
//
// 使用该方法需要了解到 []byte 将和 string 公用一块内存， 修改 []byte 的数据将导致 string 发生变化，
// 这打破了字符串不可以修改的特性，如果你恰恰需要这么做，可能非常的有用。
// 要保证字符串不可变的特性，就必须保证 []byte 不会发生变化，或者立即消费 string,
// 往往这个非常的有用， 比如我们需要打印日志：
//
// b := []byte("hello word")
// log.Println(ToString(b))
//
// 尽快的消耗掉 string 是个好主意， 也可以遗忘掉 []byte 后面不在使用这个， 而只使用 string。
//
// 比较好的例子是 exstrings.UnsafePad 系列函数，在函数内部使用 []byte 作为字符串缓冲区，返回字符串通过该方法转换。
func ToString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}
