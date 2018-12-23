// Copyright (C) 2018  Qi Yin <qiyin@thinkeridea.com>
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

// Package exstrings 收集常规的 string 操作，作为 go 标准库 strings 的扩展。
// 避免重复编写 string 相关操作代码，集中在一起更有利于优化代码，保证代码质量。
//
// 这个包会使用一些特殊的操作来减少内存开销，已获得更好的程序性能。
//
// 我也会在这个包重写 标准库 strings 里面的一些方法，以适用于不同场景的性能优化。
package exstrings
