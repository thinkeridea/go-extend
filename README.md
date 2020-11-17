# go-extend

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/thinkeridea/go-extend)
[![Build Status](https://travis-ci.org/thinkeridea/go-extend.svg?branch=master)](https://travis-ci.org/thinkeridea/go-extend)
[![codecov](https://codecov.io/gh/thinkeridea/go-extend/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkeridea/go-extend)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkeridea/go-extend)](https://goreportcard.com/report/github.com/thinkeridea/go-extend)

go-extend 收集一些常用的操作函数，辅助更快的完成开发工作，并减少重复代码。

它收集各种杂项函数，并进行归类，方便使用者查找，它可以大幅度提升开发效率和程序运行性能。它以保证性能为最大前提，提供有效的方法。
针对一些标准库中的函数或者库进行一些修改，使其性能大幅度提升，但它并不用来替换标准库函数，这些函数往往会在一些场景下有效，但有些函数可以用来替换标准库函数，它们保持一致的功能，且相当安全。

一些包或者函数使用示例及分析可以在我的 [博客(https://blog.thinkeridea.com)](https://blog.thinkeridea.com) 中找到。

## 安装

```shell
$ go get  github.com/thinkeridea/go-extend/...
```

## 规范:

- 与标准库包名一致的使用 `ex` 前缀， 避免与标准库包冲突
- 包目录下 `doc.go` 作为包说明文档

## 性能测试

包中一些函数会进行性能测试，包括每次修订的性能对比，它们一般位于各自包下面的 `benchmark` 目录下，性能测试结果可以在 [benchmark.md](benchmark.md) 快速浏览。

## 标准库函数改进列表

用来替换标准库的函数，它们和标准库函数功能保持一致，并拥有更好的性能：

- [exstrings.Join](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Join) 该方法是对标准库 strings.Join 修改，配合 unsafe 包能有效减少内存分配
- [exstrings.Repeat](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Repeat) 该方法是对标准库 strings.Repeat 修改，对于创建大字符串能有效减少内存分配
- [exstrings.Replace](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Replace) 替换字符串 该方法是对标准库 strings.Replace 修改，配合 unsafe 包能有效减少内存分配

用该改善标准库的函数，它们基本和标准库功能一致，但是它们都拥有更好的性能：

- [exbytes.Replace](https://pkg.go.dev/github.com/thinkeridea/go-extend/exbytes#Replace) 它使用原地替换，直接修改输入的数据，获得更好的性能
- [exstrings.JoinToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.RepeatToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#RepeatToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.ReplaceToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#ReplaceToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.UnsafeReplaceToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#UnsafeReplaceToBytes) 它响应一个 `[]byte` 类型，并进行原地替换，它不能接收一个字面量字符串，否则会发生严重错误

## [API 列表](https://pkg.go.dev/github.com/thinkeridea/go-extend)

### [exbytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exbytes)

标准库 `bytes` 的扩展包，提供一些高效的 `[]byte` 操作方法。

- [exbytes.ToString](https://pkg.go.dev/github.com/thinkeridea/go-extend/exbytes#ToString) 结合 `unsafe` 使 `[]byte` 转 `string` 没有成本
- [exbytes.Replace](https://pkg.go.dev/github.com/thinkeridea/go-extend/exbytes#Replace) 它使用原地替换，直接修改输入的数据，获得更好的性能
- [exbytes.Reverse](https://pkg.go.dev/github.com/thinkeridea/go-extend/exbytes#Reverse) 原地反转 []byte
- [exbytes.Sub](https://pkg.go.dev/github.com/thinkeridea/go-extend/exbytes#Sub) 是 [exutf8.RuneSub](https://pkg.go.dev/github.com/thinkeridea/go-extend/exunicode/exutf8#RuneSub) 的别名，提供字符数量截取字节数组的方法

### [exstrings](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings)

标准库 `strings` 的扩展包，提供一些高效的 `string` 操作方法。

- [exstrings.UnsafeToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#UnsafeToBytes) 结合 `unsafe` 使 `string` 转 `[]byte` 没有成本，但是字面量字符串转换后不可修改，否则会出现严重错误
- [exstrings.Bytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Bytes) 把字符串转换成 []byte 类型，和 []byte(s) 操作结果一致，但是效率更高
- [exstrings.Copy](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Copy) 拷贝一个字符串，在截取字符串之后，避免得到大字符串引用导致内存泄漏
- [exstrings.SubString](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#SubString) 是 [exutf8.RuneSubString](https://pkg.go.dev/github.com/thinkeridea/go-extend/exunicode/exutf8#RuneSubString) 的别名，根据字符数量截取字符串的方法
- [exstrings.Join](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Join) 该方法是对标准库 strings.Join 修改，配合 unsafe 包能有效减少内存分配
- [exstrings.JoinToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.Repeat](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Repeat) 该方法是对标准库 strings.Repeat 修改，对于创建大字符串能有效减少内存分配
- [exstrings.RepeatToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#RepeatToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.Replace](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Replace) 替换字符串 该方法是对标准库 strings.Replace 修改，配合 unsafe 包能有效减少内存分配
- [exstrings.ReplaceToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#ReplaceToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.UnsafeReplaceToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#UnsafeReplaceToBytes) 它响应一个 `[]byte` 类型，并进行原地替换，它不能接收一个字面量字符串，否则会发生严重错误
- [exstrings.Reverse](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Reverse) 一个高效反转字符串的方法，它支持 `UFT-8` 编码
- [exstrings.ReverseASCII](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#ReverseASCII) 一个高效反转字符串的方法，它支持 `ASCII` 编码
- [exstrings.UnsafeReverseASCII](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#UnsafeReverseASCII) 一个高效原地反转字符串的方法，它支持 `ASCII` 编码，它不能接收一个字面量字符串，否则会发生严重错误

#### [exstrings](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings) Pad 系列

一系列快速填充字符串的方法，可以做字符串格式化对齐相关操作。

- [exstrings.Pad](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Pad) 使用一个字符串填充另一个字符串，它有一系列方法，支持左、右、两边填充
- [exstrings.BothPad](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#BothPad) 使用另一个字符串从两端填充字符串为指定长度， 如果补充长度是奇数，右边的字符会更多一些
- [exstrings.LeftPad](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#LeftPad) 使用另一个字符串从左端填充字符串为指定长度
- [exstrings.RightPad](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#RightPad) 使用另一个字符串从右端填充字符串为指定长度

#### [exstrings](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings) JoinInts 系列

一些列快速高效的连接数值切片。

- [exstrings.JoinInts](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinInts) 使用 sep 连接 []int 并返回连接的字符串
- [exstrings.JoinUints](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinUints) 使用 sep 连接 []uint 并返回连接的字符串
- [exstrings.JoinInt8s](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinInt8s) 使用 sep 连接 []int8 并返回连接的字符串
- [exstrings.JoinInt16s](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinInt16s) 使用 sep 连接 []int16 并返回连接的字符串
- [exstrings.JoinInt32s](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinInt32s) 使用 sep 连接 []int32 并返回连接的字符串
- [exstrings.JoinInt64s](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinInt64s) 使用 sep 连接 []int64 并返回连接的字符串
- [exstrings.JoinUint8s](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinUint8s) 使用 sep 连接 []uint8 并返回连接的字符串
- [exstrings.JoinUint16s](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinUint16s) 使用 sep 连接 []uint16 并返回连接的字符串
- [exstrings.JoinUint32s](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinUint32s) 使用 sep 连接 []uint32 并返回连接的字符串
- [exstrings.JoinUint64s](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#JoinUint64s) 使用 sep 连接 []uint64 并返回连接的字符串

### [exnet](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet)

标准库 `net` 的扩展包，提供一些常用的操作函数。

- [exnet.HasLocalIPddr](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet#HasLocalIPddr) 检查一个ip是否是内网ip
- [exnet.HasLocalIP](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet#HasLocalIP) 检查一个ip对象是否是内网ip
- [exnet.ClientIP](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet#ClientIP) 尽最大努力实现获取客户端 IP 的算法
- [exnet.ClientPublicIP](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet#ClientPublicIP) 尽最大努力实现获取客户端公网 IP 的算法
- [exnet.RemoteIP](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet#RemoteIP) 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法
- [exnet.IPString2Long](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet#IPString2Long) 把ip字符串转为数值
- [exnet.IP2Long](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet#IP2Long) 把net.IP转为数值
- [exnet.Long2IPString](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet#Long2IPString) 把数值转为ip字符串
- [exnet.Long2IP](https://pkg.go.dev/github.com/thinkeridea/go-extend/exnet#Long2IP) 把数值转为net.IP

### [exutf8](https://pkg.go.dev/github.com/thinkeridea/go-extend/exunicode/exutf8)

标准库 `utf8` 的扩展包，提供一些高效处理多字节字符的方法。

- [exutf8.RuneIndex](https://pkg.go.dev/github.com/thinkeridea/go-extend/exunicode/exutf8#RuneIndex) 计算指定字符数量的字节索引位置
- [exutf8.RuneIndexInString](https://pkg.go.dev/github.com/thinkeridea/go-extend/exunicode/exutf8#RuneIndexInString) 计算指定字符数量的字符串索引位置
- [exutf8.RuneSub](https://pkg.go.dev/github.com/thinkeridea/go-extend/exunicode/exutf8#RuneSub) 多字节字符截取方法
- [exutf8.RuneSubString](https://pkg.go.dev/github.com/thinkeridea/go-extend/exunicode/exutf8#RuneSubString) 多字节字符串截取

### [exatomic](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic)

一个浮点数的原子包， 支持 `float32` 和 `float64` 类型，鉴于浮点数的精度可能会有些意外，详细查看包文档或浮点数算法。

- [exatomic.AddFloat32](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#AddFloat32) atomically adds delta to *addr and returns the new value.
- [exatomic.CompareAndSwapFloat32](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#CompareAndSwapFloat32) executes the compare-and-swap operation for an float32 value.
- [exatomic.LoadFloat32](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#LoadFloat32) atomically loads *addr.
- [exatomic.StoreFloat32](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#StoreFloat32) atomically stores val into *addr.
- [exatomic.SwapFloat32](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#SwapFloat32) atomically stores new into *addr and returns the previous *addr value.
- [exatomic.AddFloat64](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#AddFloat64) atomically adds delta to *addr and returns the new value.
- [exatomic.CompareAndSwapFloat64](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#CompareAndSwapFloat64) executes the compare-and-swap operation for an float64 value.
- [exatomic.LoadFloat64](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#LoadFloat64) atomically loads *addr.
- [exatomic.StoreFloat64](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#StoreFloat64) atomically stores val into *addr.
- [exatomic.SwapFloat64](https://pkg.go.dev/github.com/thinkeridea/go-extend/exatomic#SwapFloat64) atomically stores new into *addr and returns the previous *addr value.


### [exsync](https://pkg.go.dev/github.com/thinkeridea/go-extend/exsync)

对 `sync` 的扩展，包装部分功能使其更加易用和实用。

#### [Once](https://pkg.go.dev/github.com/thinkeridea/go-extend/exsync#Once)

`sync.Once` 的扩展实现，由于每次使用 `sync.Once` 都需要保存两个字段，一个是 `sync.Once` 的实例，一个是数据本身，这带来一些混乱。
让 `Once` 自带数据保存，减少使用时需要定义多个字段，如果需要保存多个数据，可以使用 `[]interface{}` 或者自定义 `struct`。

- [Once.Do](https://pkg.go.dev/github.com/thinkeridea/go-extend/exsync#Once.Do) 和 `sync.Once.Do` 相似，响应 `interface{}`。

#### [OncePointer](https://pkg.go.dev/github.com/thinkeridea/go-extend/exsync#OncePointer)

`OncePointer` 性能方面略好于 `Once`，但不会有太大改善，依然落后于 `sync.Once`， 在某些场景下可以使用，更推荐使用 `Once`。

- [OncePointer.Do](https://pkg.go.dev/github.com/thinkeridea/go-extend/exsync#OncePointer.Do)和 `sync.Once.Do` 相似，响应 `unsafe.Pointer`。


### [pool](https://pkg.go.dev/github.com/thinkeridea/go-extend/pool)

一个 `buffer` 的公共缓存池，有一些固定的长度可选，可以有效减少程序使用的 `sync.Pool` 数量, 提供多种预定义大小的缓存池。

- [pool.GetBuff64](https://pkg.go.dev/github.com/thinkeridea/go-extend/pool#GetBuff64)  获取一个初始容量为 64 的 *bytes.Buffer Pool
- [pool.GetBuff128](https://pkg.go.dev/github.com/thinkeridea/go-extend/pool#GetBuff128) 获取一个初始容量为 128 的 *bytes.Buffer Pool
- [pool.GetBuff512](https://pkg.go.dev/github.com/thinkeridea/go-extend/pool#GetBuff512) 获取一个初始容量为 512 的 *bytes.Buffer Pool
- [pool.GetBuff1024](https://pkg.go.dev/github.com/thinkeridea/go-extend/pool#GetBuff1024) 获取一个初始容量为 1024 的 *bytes.Buffer Pool
- [pool.GetBuff2048](https://pkg.go.dev/github.com/thinkeridea/go-extend/pool#GetBuff2048) 获取一个初始容量为 2048 的 *bytes.Buffer Pool
- [pool.GetBuff4096](https://pkg.go.dev/github.com/thinkeridea/go-extend/pool#GetBuff4096) 获取一个初始容量为 4096 的 *bytes.Buffer Pool
- [pool.GetBuff8192](https://pkg.go.dev/github.com/thinkeridea/go-extend/pool#GetBuff8192) 获取一个初始容量为 8192 的 *bytes.Buffer Pool

该模块定义了一个 `BufferPool` 接口，所有的预分配池均返回该接口，该接口提供两个方法：

- `Get() *bytes.Buffer` 获取一个 `*bytes.Buffer`，且该实例已经被 Reset
- `Put(*bytes.Buffer)`  把 `*bytes.Buffer` 放回 Pool 中

### [helper](https://pkg.go.dev/github.com/thinkeridea/go-extend/helper)

- [helper.Must](https://pkg.go.dev/github.com/thinkeridea/go-extend/helper#Must) 快速的构建安全的方法，出现错误会 `panic`, 很适合程序初始化组件使用
- [helper.PanicRecover](https://pkg.go.dev/github.com/thinkeridea/go-extend/helper#PanicRecover) 把 `panic` 转成 `error` ，并且打印栈信息到错误输出

### [datalog](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog)

用于辅助拼接类 csv 格式化数据日志的组件。

#### [datalog](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog).[Record](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record)

用于拼接日志

- [Record.NewRecord](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record.NewRecord) 创建一个固定长度的日志行记录器
- [Record.NewRecordPool](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record.NewRecordPool) 创建长度固定的日志记录缓存池
- [Record.Clean](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record.Clean) 清空 Record 中的所有元素，如果使用 sync.Pool 在放回 Pool 之前应该清空 Record，避免内存泄漏
- [Record.Join](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record.Join) 使用 sep 连接 Record， 并在末尾追加 suffix 这个类似 strings.Join 方法
- [Record.ToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record.ToBytes) 使用 sep 连接 Record，并在末尾添加 newline 换行符
- [Record.UnsafeToBytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record.UnsafeToBytes)  使用 sep 连接 Record，并在末尾添加 newline 换行符，会使用 `unsafe` 包减少内存分配
- [Record.ArrayJoin](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record.ArrayJoin) 使用 sep 连接 Record，其结果作为数组字段的值
- [Record.ArrayFieldJoin](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record.ArrayFieldJoin) 使用 fieldSep 连接 Record，其结果作为一个数组的单元
- [Record.UnsafeArrayFieldJoin](https://pkg.go.dev/github.com/thinkeridea/go-extend/datalog#Record.UnsafeArrayFieldJoin)   使用 fieldSep 连接 Record，其结果作为一个数组的单元，会使用 `unsafe` 包减少内存分配

## 许可

go-extend 根据 [GNU General Public License v3.0](https://www.gnu.org/licenses/) 许可证授权，有关完整许可证文本，请参阅 [LICENSE](LICENSE)
