[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/thinkeridea/go-extend)
[![Build Status](https://travis-ci.org/thinkeridea/go-extend.svg?branch=master)](https://travis-ci.org/thinkeridea/go-extend)
[![codecov](https://codecov.io/gh/thinkeridea/go-extend/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkeridea/go-extend)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkeridea/go-extend)](https://goreportcard.com/report/github.com/thinkeridea/go-extend)

# go-extend 收集一些常用的操作函数，辅助更快的完成开发工作，并减少重复代码

它收集各种杂项函数，并进行归类，方便使用者查找，它可以大幅度提升开发效率和程序运行性能。它以保证性能为最大前提，提供有效的方法。
针对一些标准库中的函数或者库进行一些修改，使其性能大幅度提升，但它并不用来替换标准库函数，这些函数往往会在一些场景下有效，但有些函数可以用来替换标准库函数，它们保持一致的功能，且相当安全。

一些包或者函数使用示例及分析可以在我的 [博客(https://blog.thinkeridea.com)](https://blog.thinkeridea.com) 中找到。

## 规范:
	
- 与标准库包名一致的使用 `ex` 前缀， 避免与标准库包冲突
- 包目录下 `doc.go` 作为包说明文档

## 标准库函数改进列表

用来替换标准库的函数，它们和标准库函数功能保持一致，并拥有更好的性能：

- [exstrings.Join](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Join) 该方法是对标准库 strings.Join 修改，配合 unsafe 包能有效减少内存分配
- [exstrings.Repeat](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Repeat) 该方法是对标准库 strings.Repeat 修改，对于创建大字符串能有效减少内存分配
- [exstrings.Replace](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Replace) 替换字符串 该方法是对标准库 strings.Replace 修改，配合 unsafe 包能有效减少内存分配

用该改善标准库的函数，它们基本和标准库功能一致，但是它们都拥有更好的性能：

- [exbytes.Replace](https://godoc.org/github.com/thinkeridea/go-extend/exbytes#Replace) 它使用原地替换，直接修改输入的数据，获得更好的性能
- [exstrings.JoinToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.RepeatToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#RepeatToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.ReplaceToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#ReplaceToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.UnsafeReplaceToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#UnsafeReplaceToBytes) 它响应一个 `[]byte` 类型，并进行原地替换，它不能接收一个字面量字符串，否则会发生严重错误

## [API 列表](http://godoc.org/github.com/thinkeridea/go-extend)

### [exbytes](https://godoc.org/github.com/thinkeridea/go-extend/exbytes)

标准库 `bytes` 的扩展包，提供一些高效的 `[]byte` 操作方法。

- [exbytes.ToString](https://godoc.org/github.com/thinkeridea/go-extend/exbytes#ToString) 结合 `unsafe` 使 `[]byte` 转 `string` 没有成本
- [exbytes.Replace](https://godoc.org/github.com/thinkeridea/go-extend/exbytes#Replace) 它使用原地替换，直接修改输入的数据，获得更好的性能
- [exbytes.Reverse](https://godoc.org/github.com/thinkeridea/go-extend/exbytes#Reverse) 原地反转 []byte

### [exstrings](https://godoc.org/github.com/thinkeridea/go-extend/exstrings)

标准库 `strings` 的扩展包，提供一些高效的 `string` 操作方法。

- [exstrings.UnsafeToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#UnsafeToBytes) 结合 `unsafe` 使 `string` 转 `[]byte` 没有成本，但是字面量字符串转换后不可修改，否则会出现严重错误
- [exstrings.Copy](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Copy) 拷贝一个字符串，在截取字符串之后，避免得到大字符串引用导致内存泄漏
- [exstrings.Join](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Join) 该方法是对标准库 strings.Join 修改，配合 unsafe 包能有效减少内存分配
- [exstrings.JoinToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.Repeat](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Repeat) 该方法是对标准库 strings.Repeat 修改，对于创建大字符串能有效减少内存分配
- [exstrings.RepeatToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#RepeatToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.Replace](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Replace) 替换字符串 该方法是对标准库 strings.Replace 修改，配合 unsafe 包能有效减少内存分配
- [exstrings.ReplaceToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#ReplaceToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换
- [exstrings.UnsafeReplaceToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#UnsafeReplaceToBytes) 它响应一个 `[]byte` 类型，并进行原地替换，它不能接收一个字面量字符串，否则会发生严重错误
- [exstrings.Reverse](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Reverse) 一个高效反转字符串的方法，它支持 `UFT-8` 编码
- [exstrings.ReverseASCII](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#ReverseASCII) 一个高效反转字符串的方法，它支持 `ASCII` 编码
- [exstrings.UnsafeReverseASCII](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#UnsafeReverseASCII) 一个高效原地反转字符串的方法，它支持 `ASCII` 编码，它不能接收一个字面量字符串，否则会发生严重错误

#### [exstrings](https://godoc.org/github.com/thinkeridea/go-extend/exstrings) Pad 系列

一系列快速填充字符串的方法，可以做字符串格式化对齐相关操作。

- [exstrings.Pad](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Pad) 使用一个字符串填充另一个字符串，它有一系列方法，支持左、右、两边填充
- [exstrings.BothPad](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#BothPad) 使用另一个字符串从两端填充字符串为指定长度， 如果补充长度是奇数，右边的字符会更多一些
- [exstrings.LeftPad](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#LeftPad) 使用另一个字符串从左端填充字符串为指定长度
- [exstrings.RightPad](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#RightPad) 使用另一个字符串从右端填充字符串为指定长度

#### [exstrings](https://godoc.org/github.com/thinkeridea/go-extend/exstrings) JoinInts 系列

一些列快速高效的连接数值切片。

- [exstrings.JoinInts](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinInts) 使用 sep 连接 []int 并返回连接的字符串
- [exstrings.JoinUints](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinUints) 使用 sep 连接 []uint 并返回连接的字符串
- [exstrings.JoinInt8s](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinInt8s) 使用 sep 连接 []int8 并返回连接的字符串
- [exstrings.JoinInt16s](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinInt16s) 使用 sep 连接 []int16 并返回连接的字符串
- [exstrings.JoinInt32s](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinInt32s) 使用 sep 连接 []int32 并返回连接的字符串
- [exstrings.JoinInt64s](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinInt64s) 使用 sep 连接 []int64 并返回连接的字符串
- [exstrings.JoinUint8s](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinUint8s) 使用 sep 连接 []uint8 并返回连接的字符串
- [exstrings.JoinUint16s](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinUint16s) 使用 sep 连接 []uint16 并返回连接的字符串
- [exstrings.JoinUint32s](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinUint32s) 使用 sep 连接 []uint32 并返回连接的字符串
- [exstrings.JoinUint64s](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinUint64s) 使用 sep 连接 []uint64 并返回连接的字符串

### [exnet](https://godoc.org/github.com/thinkeridea/go-extend/exnet)

标准库 `exnet` 的扩展包，提供一些常用的操作函数。

- [exnet.HasLocalIPddr](https://godoc.org/github.com/thinkeridea/go-extend/exnet#HasLocalIPddr) 检查一个ip是否是内网ip
- [exnet.HasLocalIP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#HasLocalIP) 检查一个ip对象是否是内网ip
- [exnet.ClientIP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#ClientIP) 尽最大努力实现获取客户端 IP 的算法
- [exnet.ClientPublicIP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#ClientPublicIP) 尽最大努力实现获取客户端公网 IP 的算法
- [exnet.RemoteIP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#RemoteIP) 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法
- [exnet.IPString2Long](https://godoc.org/github.com/thinkeridea/go-extend/exnet#IPString2Long) 把ip字符串转为数值
- [exnet.IP2Long](https://godoc.org/github.com/thinkeridea/go-extend/exnet#IP2Long) 把net.IP转为数值
- [exnet.Long2IPString](https://godoc.org/github.com/thinkeridea/go-extend/exnet#Long2IPString) 把数值转为ip字符串
- [exnet.Long2IP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#Long2IP) 把数值转为net.IP

### [exatomic](https://godoc.org/github.com/thinkeridea/go-extend/exatomic)

一个浮点数的原子包， 支持 `float32` 和 `float64` 类型，鉴于浮点数的精度可能会有些意外，详细查看包文档或浮点数算法。

- [exatomic.AddFloat32](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#AddFloat32) atomically adds delta to *addr and returns the new value.
- [exatomic.CompareAndSwapFloat32](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#CompareAndSwapFloat32) executes the compare-and-swap operation for an float32 value.
- [exatomic.LoadFloat32](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#LoadFloat32) atomically loads *addr.
- [exatomic.StoreFloat32](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#StoreFloat32) atomically stores val into *addr.
- [exatomic.SwapFloat32](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#SwapFloat32) atomically stores new into *addr and returns the previous *addr value.
- [exatomic.AddFloat64](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#AddFloat64) atomically adds delta to *addr and returns the new value.
- [exatomic.CompareAndSwapFloat64](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#CompareAndSwapFloat64) executes the compare-and-swap operation for an float64 value.
- [exatomic.LoadFloat64](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#LoadFloat64) atomically loads *addr.
- [exatomic.StoreFloat64](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#StoreFloat64) atomically stores val into *addr.
- [exatomic.SwapFloat64](https://godoc.org/github.com/thinkeridea/go-extend/exatomic#SwapFloat64) atomically stores new into *addr and returns the previous *addr value.

### [pool](https://godoc.org/github.com/thinkeridea/go-extend/pool)

一个 `buffer` 的公共缓存池，有一些固定的长度可选，可以有效减少程序使用的 `sync.Pool` 数量, 提供多种预定义大小的缓存池。

- [pool.GetBuff64](https://godoc.org/github.com/thinkeridea/go-extend/pool#GetBuff64)  获取一个初始容量为 64 的 *bytes.Buffer Pool
- [pool.GetBuff128](https://godoc.org/github.com/thinkeridea/go-extend/pool#GetBuff128) 获取一个初始容量为 128 的 *bytes.Buffer Pool
- [pool.GetBuff512](https://godoc.org/github.com/thinkeridea/go-extend/pool#GetBuff512) 获取一个初始容量为 512 的 *bytes.Buffer Pool
- [pool.GetBuff1024](https://godoc.org/github.com/thinkeridea/go-extend/pool#GetBuff1024) 获取一个初始容量为 1024 的 *bytes.Buffer Pool
- [pool.GetBuff2048](https://godoc.org/github.com/thinkeridea/go-extend/pool#GetBuff2048) 获取一个初始容量为 2048 的 *bytes.Buffer Pool
- [pool.GetBuff4096](https://godoc.org/github.com/thinkeridea/go-extend/pool#GetBuff4096) 获取一个初始容量为 4096 的 *bytes.Buffer Pool
- [pool.GetBuff8192](https://godoc.org/github.com/thinkeridea/go-extend/pool#GetBuff8192) 获取一个初始容量为 8192 的 *bytes.Buffer Pool

该模块定义了一个 `BufferPool` 接口，所有的预分配池均返回该接口，该接口提供两个方法：

- `Get() *bytes.Buffer` 获取一个 `*bytes.Buffer`，且该实例已经被 Reset
- `Put(*bytes.Buffer)`  把 `*bytes.Buffer` 放回 Pool 中

### [helper](https://godoc.org/github.com/thinkeridea/go-extend/helper)
 
- [helper.Must](https://godoc.org/github.com/thinkeridea/go-extend/helper#Must) 快速的构建安全的方法，出现错误会 `panic`, 很适合程序初始化组件使用
- [helper.PanicRecover](https://godoc.org/github.com/thinkeridea/go-extend/helper#PanicRecover) 把 `panic` 转成 `error` ，并且打印栈信息到错误输出

### [datalog](https://godoc.org/github.com/thinkeridea/go-extend/datalog)

用于辅助拼接类 csv 格式化数据日志的组件。

#### [datalog](https://godoc.org/github.com/thinkeridea/go-extend/datalog).[Record](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record)

用于拼接日志

- [Record.NewRecord](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record.NewRecord) 创建一个固定长度的日志行记录器
- [Record.NewRecordPool](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record.NewRecordPool) 创建长度固定的日志记录缓存池 
- [Record.Clean](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record.Clean) 清空 Record 中的所有元素，如果使用 sync.Pool 在放回 Pool 之前应该清空 Record，避免内存泄漏
- [Record.Join](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record.Join) 使用 sep 连接 Record， 并在末尾追加 suffix 这个类似 strings.Join 方法
- [Record.ToBytes](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record.ToBytes) 使用 sep 连接 Record，并在末尾添加 newline 换行符
- [Record.UnsafeToBytes](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record.UnsafeToBytes)  使用 sep 连接 Record，并在末尾添加 newline 换行符，会使用 `unsafe` 包减少内存分配
- [Record.ArrayJoin](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record.ArrayJoin) 使用 sep 连接 Record，其结果作为数组字段的值
- [Record.ArrayFieldJoin](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record.ArrayFieldJoin) 使用 fieldSep 连接 Record，其结果作为一个数组的单元
- [Record.UnsafeArrayFieldJoin](https://godoc.org/github.com/thinkeridea/go-extend/datalog#Record.UnsafeArrayFieldJoin)   使用 fieldSep 连接 Record，其结果作为一个数组的单元，会使用 `unsafe` 包减少内存分配

## 标准库函数改进性能测试：

- [exstrings](exstrings/)

	- Replace 系列函数测试报告

	```
	goos: darwin
	goarch: amd64
	pkg: github.com/thinkeridea/go-extend/exstrings/benchmark
	BenchmarkReplace-8                  	  500000	      3256 ns/op	     960 B/op	      15 allocs/op
	BenchmarkReplaceToBytes-8           	  500000	      3283 ns/op	    1024 B/op	      16 allocs/op
	BenchmarkUnsafeReplaceToBytes-8     	  500000	      3041 ns/op	     960 B/op	      15 allocs/op
	BenchmarkStandardLibraryReplace-8   	  500000	      3679 ns/op	    1920 B/op	      30 allocs/op
	PASS
	ok  	github.com/thinkeridea/go-extend/exstrings/benchmark	7.769s
	```

	- Repeat 性能测试报告
	
	```
	goos: darwin
    goarch: amd64
    pkg: github.com/thinkeridea/go-extend/exstrings/benchmark
    BenchmarkRepeat-8                  	   50000	     28818 ns/op	  303104 B/op	       1 allocs/op
    BenchmarkRepeatToBytes-8           	   50000	     28104 ns/op	  303104 B/op	       1 allocs/op
    BenchmarkStandardLibraryRepeat-8   	   20000	     51968 ns/op	  606208 B/op	       2 allocs/op
    PASS
    ok  	github.com/thinkeridea/go-extend/exstrings/benchmark	6.200s
	```

	- Join 性能测试报告
	
	```
	goos: darwin
    goarch: amd64
    pkg: github.com/thinkeridea/go-extend/exstrings/benchmark
    BenchmarkJoin-8                  	 5000000	       290 ns/op	      64 B/op	       1 allocs/op
    BenchmarkJoinToBytes-8           	 5000000	       290 ns/op	      64 B/op	       1 allocs/op
    BenchmarkStandardLibraryJoin-8   	 5000000	       315 ns/op	     128 B/op	       2 allocs/op
    PASS
    ok  	github.com/thinkeridea/go-extend/exstrings/benchmark	5.406s
	```

## 许可

go-extend 根据 [GNU General Public License v3.0](https://www.gnu.org/licenses/) 许可证授权，有关完整许可证文本，请参阅 [LICENSE](LICENSE)
