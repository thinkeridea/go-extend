[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/thinkeridea/go-extend)
[![Build Status](https://travis-ci.org/thinkeridea/go-extend.svg?branch=master)](https://travis-ci.org/thinkeridea/go-extend)
[![codecov](https://codecov.io/gh/thinkeridea/go-extend/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkeridea/go-extend)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkeridea/go-extend)](https://goreportcard.com/report/github.com/thinkeridea/go-extend)

# go-extend 收集一些常用的操作函数，辅助更快的完成开发工作，并减少重复代码。

它收集各种杂项函数，并进行归类，方便使用者查找，它可以大幅度提升开发效率和程序运行性能。它以保证性能为最大前提，提供有效的方法。
针对一些标准库中的函数或者库进行一些修改，使其性能大幅度提升，但它并不用来替换标准库函数，这些函数往往会在一些场景下有效，但有些函数可以用来替换标准库函数，它们保持一致的功能，且相当安全。

一些包或者函数使用示例及分析可以在我的 [博客(https://blog.thinkeridea.com)](https://blog.thinkeridea.com) 中找到。

## 规范:
	
- 与标准库包名一致的使用 `ex` 前缀， 避免与标准库包冲突。
- 包目录下 `doc.go` 作为包说明文档。

## 标准库函数改进列表

用来替换标准库的函数，它们和标准库函数功能保持一致，并拥有更好的性能：

- [exstrings.Join](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Join)
- [exstrings.Repeat](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Repeat)
- [exstrings.Replace](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Replace)

用该改善标准库的函数，它们基本和标准库功能一致，但是它们都拥有更好的性能：

- [exbytes.Replace](https://godoc.org/github.com/thinkeridea/go-extend/exbytes#Replace) 它使用原地替换，直接修改输入的数据，获得更好的性能。
- [exstring.JoinToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换。
- [exstrings.RepeatToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#RepeatToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换。
- [exstrings.ReplaceToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#ReplaceToBytes) 它响应一个 `[]byte` 类型，有效避免类型转换。
- [exstrings.UnsafeReplaceToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#UnsafeReplaceToBytes) 它响应一个 `[]byte` 类型，并进行原地替换，它不能接收一个字面量字符串，否则会发生严重错误。

## 很有用的方法

- [exbytes.ToString](https://godoc.org/github.com/thinkeridea/go-extend/exbytes#ToString) 结合 `unsafe` 使 `[]byte` 转 `string` 没有成本。
- [exstrings.UnsafeToBytes](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#UnsafeToBytes) 结合 `unsafe` 使 `string` 转 `[]byte` 没有成本，但是字面量字符串转换后不可修改，否则会出现严重错误。
- [exstrings.JoinInts](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#JoinInts) 连接数字切片，这有一系列方法，针对各种数值类型。
- [exstrings.Pad](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Pad) 使用一个字符串填充另一个字符串，它有一系列方法，支持左、右、两边填充，`Unsafe` 系列前缀方法已被弃用。
- [exbytes.Reverse](https://godoc.org/github.com/thinkeridea/go-extend/exbytes#Reverse) 原地反转 []byte。
- [exstrings.Reverse](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#Reverse) 一个高效反转字符串的方法，它支持 `UFT-8` 编码。
- [exstrings.ReverseASCII](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#ReverseASCII) 一个高效反转字符串的方法，它支持 `ASCII` 编码。
- [exstrings.UnsafeReverseASCII](https://godoc.org/github.com/thinkeridea/go-extend/exstrings#UnsafeReverseASCII) 一个高效原地反转字符串的方法，它支持 `ASCII` 编码，它不能接收一个字面量字符串，否则会发生严重错误。
- [exnet.HasLocalIPddr](https://godoc.org/github.com/thinkeridea/go-extend/exnet#HasLocalIPddr) 检查一个ip是否是内网ip。
- [exnet.HasLocalIP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#HasLocalIPddr) 检查一个ip对象是否是内网ip。
- [exnet.ClientIP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#ClientIP) 尽最大努力实现获取客户端 IP 的算法。
- [exnet.ClientPublicIP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#ClientPublicIP) 尽最大努力实现获取客户端公网 IP 的算法。
- [exnet.RemoteIP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#RemoteIP) 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
- [exnet.IPString2Long](https://godoc.org/github.com/thinkeridea/go-extend/exnet#IPString2Long) 把ip字符串转为数值。
- [exnet.Long2IPString](https://godoc.org/github.com/thinkeridea/go-extend/exnet#Long2IPString) 把数值转为ip字符串。
- [exnet.IP2Long](https://godoc.org/github.com/thinkeridea/go-extend/exnet#IP2Long) 把net.IP转为数值。
- [exnet.Long2IP](https://godoc.org/github.com/thinkeridea/go-extend/exnet#Long2IP) 把数值转为net.IP。
- [helper.Must](https://godoc.org/github.com/thinkeridea/go-extend/helper#Must) 快速的构建安全的方法，出现错误会 `panic`, 很适合程序初始化组件使用。
- [helper.PanicRecover](https://godoc.org/github.com/thinkeridea/go-extend/helper#PanicRecover) 把 `panic` 转成 `error` ，并且打印栈信息到错误输出。

## 非常有用的包

- [pool](https://godoc.org/github.com/thinkeridea/go-extend/pool) 一个 `buffer` 的公共缓存池，有一些固定的长度可选，可以有效减少程序使用的 `sync.Pool` 数量。
- [exatomic](https://godoc.org/github.com/thinkeridea/go-extend/exatomic) 一个浮点数的原子包， 支持 `float32` 和 `float64` 类型，鉴于浮点数的精度可能会有些意外，详细查看包文档或浮点数算法。
- [datalog](https://godoc.org/github.com/thinkeridea/go-extend/datalog) 用于辅助拼接类 csv 格式化数据日志的组件。

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

go-extend 根据 [GNU General Public License v3.0](https://www.gnu.org/licenses/) 许可证授权，有关完整许可证文本，请参阅 [LICENSE](LICENSE)。
