// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package exsync

import (
	"sync"
	"unsafe"
)

// Once 是 sync.Once 的扩展实现，由于每次使用 sync.Once 都需要保存两个字段，一个是 sync.Once 的实例，一个是数据本身，这带来一些混乱
// 让 Once 自带数据保存，减少使用时需要定义多个字段，如果需要保存多个数据，可以使用 []interface{} 或者自定义 struct
//
// 以下是一个简单的示例：
// var db Once
// func DB() *mysql.Client{
// 		return db.Do(f func() interface{}{
// 			return mysql.NewClient(...)
// 		}).(*mysql.Client)
// }
//
// 当希望处理错误，可以响应 []interface{} 或者 自定义 struct， 如果在服务或程序初始化阶段可以考虑 panic 来报告错误，如下是使用 []interface{} 的示例：
// var db Once
// func DB() (*mysql.Client, error){
// 	res := db.Do(f func() interface{}{
// 		c, err:=mysql.NewClient(...)
// 		return []interface{}{c, err}
// 	}).([]interface{})
//
// 	return res[0].(*mysql.Client), res[1].(error)
// }
//
// 使用该方法需要一些取舍，它简单实用，相比 sync.Once 性能有所下降，不过它依然很快，这不会形成性能问题。
type Once struct {
	once sync.Once
	v    interface{}
}

// Do calls the function f if and only if Do is being called for the
// first time for this instance of Once. In other words, given
// 	var once Once
// if once.Do(f) is called multiple times, only the first call will invoke f,
// even if f has a different value in each invocation. A new instance of
// Once is required for each function to execute.
//
// Do is intended for initialization that must be run exactly once. Since f
// is niladic, it may be necessary to use a function literal to capture the
// arguments to a function to be invoked by Do:
// 	config.once.Do(func() { config.init(filename) })
//
// Because no call to Do returns until the one call to f returns, if f causes
// Do to be called, it will deadlock.
//
// If f panics, Do considers it to have returned; future calls of Do return
// without calling f.
//
func (o *Once) Do(f func() interface{}) interface{} {
	o.once.Do(func() {
		o.v = f()
	})

	return o.v
}

// OncePointer 性能方面略好于 Once，但不会有太大改善，依然落后于 sync.Once， 在某些场景下可以使用，更推荐使用 Once
type OncePointer struct {
	once sync.Once
	v    unsafe.Pointer
}

// Do calls the function f if and only if Do is being called for the
// first time for this instance of Once. In other words, given
// 	var once Once
// if once.Do(f) is called multiple times, only the first call will invoke f,
// even if f has a different value in each invocation. A new instance of
// Once is required for each function to execute.
//
// Do is intended for initialization that must be run exactly once. Since f
// is niladic, it may be necessary to use a function literal to capture the
// arguments to a function to be invoked by Do:
// 	config.once.Do(func() { config.init(filename) })
//
// Because no call to Do returns until the one call to f returns, if f causes
// Do to be called, it will deadlock.
//
// If f panics, Do considers it to have returned; future calls of Do return
// without calling f.
//
func (o *OncePointer) Do(f func() unsafe.Pointer) unsafe.Pointer {
	o.once.Do(func() {
		o.v = f()
	})

	return o.v
}
