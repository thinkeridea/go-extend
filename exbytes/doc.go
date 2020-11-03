// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

// Package exbytes 收集常规的 []byte 操作，作为 go 标准库 bytes 的扩展。
// 避免重复编写 []byte 相关操作代码，集中在一起更有利于优化代码，保证代码质量。
//
// 这个包会使用一些特殊的操作来减少内存开销，已获得更好的程序性能。
//
// 我也会在这个包重写 标准库 bytes 里面的一些方法，以适用于不同场景的性能优化。
package exbytes
