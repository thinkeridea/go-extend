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

package exstrings

import "unsafe"

/*
UnsafeToBytes 把 string 转换为 []byte 没有多余的内存开销。

这个函数可以提升 string 转 []byte 的性能，并极大的降低内存开销，但是却相当的危险，对于不明确这个函数的人来说不建议使用该函数。

这个函数是 exbytes.ToString 反向操作版，但是不像 exbytes.ToString 那么的稳定安全，该函数使用不当会导致程序直接崩溃，且无法恢复。

	s := `{"k":"v"}`

	b := exstrings.UnsafeToBytes(s)
	// b[3] = 'k' // unexpected fault address 0x1118180
	data := map[string]string{}
	err := json.Unmarshal(b, &data)

	fmt.Println(data, err)

这是一个使用的例子，如果我们需要转换一个字符串很方便，且开销非常的低。
但是一定要注意，b[3] = 'k' 如果尝试修改获得的 []byte 将直接导致程序崩溃，并且不可能通过 recover() 恢复。

实际上我们可以突破这个限制，这就要了解字符串的一些规则，下面的例子可以完美运行，并修改字符串：

	s := strings.Repeat("A", 3)
	b := exstrings.UnsafeToBytes(s)
	b[1] = 'B'
	b[2] = 'C'

	fmt.Println(s, string(b))

非常完美，s和b变量的值都是 ABC， 为什么会这样呢？

这个就是 string 的内存分配方法， 字面量使用这种方式是没有办法修改的，因为这是在编译时就决定的，编译时会设定字符串的内存数据是只读数据。
如果程序运行时生成的数据，这种数据是可以安全使用该函数的，但是要当心你的字符串可能会被修改，
比如我们调用 json.Unmarshal(exstrings.UnsafeToBytes(s), &data)， 如果 json 包里面出现修改输入参数，那么原来的字符串就可能不是你想想的那样。

使用该函数要明确两个事情：

	- 确定字符串是否是字面量，或者内存被分配在只读空间上。
	- 确保访问该函数结果的函数是否会按照你的意愿访问或修改数据。

我公开该函数经过很多思考，虽然它很危险，但是有时间却很有用，如果我们需要大批量转换字符串的大小写，而且不再需要原字符串，我们可以原地安全的修改字符串。
当然还有更多的使用方法，可以极大的提升我们程序的性能。
*/
func UnsafeToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

/*
Bytes 把字符串转换成 []byte 类型，和 []byte(s) 操作结果一致，但是效率更高。

我进行了性能测试，它相比 []byte(s) 性能提升大约 14%，这仅仅是个实验的函数，它可能随着编译器优化而失去性能优势。
极端性能情况下依然可以使用它，它永远和 []byte(s) 的结果一致。

	BenchmarkStandardLibraryStringToBytes-8         18584335                58.4 ns/op           192 B/op          1 allocs/op
	BenchmarkExstringsStringToBytes-8               23752122                50.1 ns/op           192 B/op          1 allocs/op
*/
func Bytes(s string) []byte {
	buf := make([]byte, len(s))
	copy(buf, s)
	return buf
}
