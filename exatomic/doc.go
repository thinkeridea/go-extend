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

/*
Package exatomic 实现了 float32 与 float64 原子操作，
我在历史项目中实时在线统计一些价格，在多线程时往往非常有用，
简化程序逻辑，并提升程序性能。

float 原子操作可以设计很多个版本，历史项目中我用过 3 种不同的实现方式，
当前项目是我最满意的版本，对程序没有任何入侵，也没有定义新的类型，保持和标准库一致的接口，非常的简单。

使用该包你必须了解到，有时间你的操作可能并不像预期一样，受 float 类型算法实现，float 精度往往并没有想象中那么高，
如果连续对一个数加一个指定值，可能最后你会得到一个固定的值，它永远不会增加，例如下面的程序：

	var x float64
	for delta := float64(1); delta+delta > delta; delta += delta {
		x = delta
		if x+1 == x {
			fmt.Printf("%.64f\n", x)
		}
	}

输出的结果有点令人意外，  x+1以后居然不会有任何变化，而且这样值有非常多,
根据不同初始值，和增加不同大小的数据，出现这种情况的数值也不同。

以下是一部分输出：
	9007199254740992.0000000000000000000000000000000000000000000000000000000000000000
	18014398509481984.0000000000000000000000000000000000000000000000000000000000000000
	36028797018963968.0000000000000000000000000000000000000000000000000000000000000000
	72057594037927936.0000000000000000000000000000000000000000000000000000000000000000
	144115188075855872.0000000000000000000000000000000000000000000000000000000000000000
	288230376151711744.0000000000000000000000000000000000000000000000000000000000000000
	576460752303423488.0000000000000000000000000000000000000000000000000000000000000000
	1152921504606846976.0000000000000000000000000000000000000000000000000000000000000000
	2305843009213693952.0000000000000000000000000000000000000000000000000000000000000000
	4611686018427387904.0000000000000000000000000000000000000000000000000000000000000000
	9223372036854775808.0000000000000000000000000000000000000000000000000000000000000000
	..........

如果你的程序向下面这样一直增加一个固定值，也许会像这样，数值到一个固定值后再也不会发生任何变化。
	var x, last float64
	for x = 9007199254740992; x < math.MaxFloat64; x++ {
		if x == last {
			fmt.Printf("%.64f\n", x)
		}

		last = x
	}

但是我们也不用太过的担心，往常这些数值都非常大，至少我很少处理这么大的 float 数值。
我想你应该了解到这个问题，往往出现 float 操作无效的问题可能并不是这个包导致的，而是 float 精度的问题导致的。
*/
package exatomic
