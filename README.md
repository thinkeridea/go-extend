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

## 许可

go-extend 根据 MIT License 许可证授权，有关完整许可证文本，请参阅 [LICENSE](LICENSE)
