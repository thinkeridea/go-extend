# go-extend

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/thinkeridea/go-extend)
[![Build Status](https://travis-ci.org/thinkeridea/go-extend.svg?branch=master)](https://travis-ci.org/thinkeridea/go-extend)
[![codecov](https://codecov.io/gh/thinkeridea/go-extend/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkeridea/go-extend)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkeridea/go-extend)](https://goreportcard.com/report/github.com/thinkeridea/go-extend)

go语言扩展包，收集一些常用的操作函数，辅助更快的完成开发工作，并减少重复代码。

规范:
	
- 包名统一使用 `ex` 前缀， 避免与官方包冲突。
- 包目录下 `doc.go` 作为包说明文档。

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

版权：

所有包统一使用 `GNU General Public License v3.0` 协议，每个源码文件开头必须包含以下版权描述：

```
// Copyright (C) <year>  <name of author> <email>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
```
