# 性能测试报告

## [exstrings](exstrings/) 包

- [exstrings.Replace](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Replace)

```shell
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

- [exstrings.Repeat](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Repeat)

```shell
goos: darwin
goarch: amd64
pkg: github.com/thinkeridea/go-extend/exstrings/benchmark
BenchmarkRepeat-8                  	   50000	     28818 ns/op	  303104 B/op	       1 allocs/op
BenchmarkRepeatToBytes-8           	   50000	     28104 ns/op	  303104 B/op	       1 allocs/op
BenchmarkStandardLibraryRepeat-8   	   20000	     51968 ns/op	  606208 B/op	       2 allocs/op
PASS
ok  	github.com/thinkeridea/go-extend/exstrings/benchmark	6.200s
```

- [exstrings.Join](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Join)

```shell
goos: darwin
goarch: amd64
pkg: github.com/thinkeridea/go-extend/exstrings/benchmark
BenchmarkJoin-8                  	 5000000	       290 ns/op	      64 B/op	       1 allocs/op
BenchmarkJoinToBytes-8           	 5000000	       290 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLibraryJoin-8   	 5000000	       315 ns/op	     128 B/op	       2 allocs/op
PASS
ok  	github.com/thinkeridea/go-extend/exstrings/benchmark	5.406s
```

- [exstrings.Reverse](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Reverse)

`BenchmarkReverseUTF8DecodeRuneInString` 是 `BenchmarkExstringsReverse` 优化前的版本

```shell
goos: darwin
goarch: amd64
pkg: github.com/thinkeridea/go-extend/exstrings/benchmark
BenchmarkReverseRunes-8                           710557              1695 ns/op             480 B/op          2 allocs/op
BenchmarkReverseRange-8                          1404463               845 ns/op             192 B/op          1 allocs/op
BenchmarkReverseUTF8DecodeRuneInString-8         1658835               720 ns/op             192 B/op          1 allocs/op
BenchmarkExstringsReverse-8                      1738339               691 ns/op             192 B/op          1 allocs/op
PASS
ok      github.com/thinkeridea/go-extend/exstrings/benchmark    8.030s
```

- [exstrings.Bytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exstrings#Bytes)

```shell
goos: darwin
goarch: amd64
pkg: github.com/thinkeridea/go-extend/exstrings/benchmark
BenchmarkStandardLibraryStringToBytes-8         18468846                58.6 ns/op           192 B/op          1 allocs/op
BenchmarkExstringsStringToBytes-8               23291382                50.9 ns/op           192 B/op          1 allocs/op
PASS
ok      github.com/thinkeridea/go-extend/exstrings/benchmark    2.395s
```

## [exbytes](https://pkg.go.dev/github.com/thinkeridea/go-extend/exbytes) 包

- [exbytes.Replace](https://pkg.go.dev/github.com/thinkeridea/go-extend/exbytes#Replace)

```shell
goos: darwin
goarch: amd64
pkg: github.com/thinkeridea/go-extend/exbytes/benchmark
BenchmarkReplace-8                        545372              1974 ns/op             416 B/op          1 allocs/op
BenchmarkBytesReplace-8                   598182              1999 ns/op             736 B/op          2 allocs/op
BenchmarkStringsReplace-8                 518322              2112 ns/op            1056 B/op          3 allocs/op
BenchmarkUnsafeStringsReplace-8           618229              1991 ns/op             736 B/op          2 allocs/op
PASS
ok      github.com/thinkeridea/go-extend/exbytes/benchmark      4.695s
```

## [exutf8](https://pkg.go.dev/github.com/thinkeridea/go-extend/exunicode/exutf8) 包

- [exutf8.RuneSubString](https://pkg.go.dev/github.com/thinkeridea/go-extend/exunicode/exutf8#RuneSubString)

```shell
goos: darwin
goarch: amd64
pkg: github.com/thinkeridea/go-extend/exunicode/exutf8/benchmark
BenchmarkSubStrRunes-8                    876604              1351 ns/op             336 B/op          2 allocs/op
BenchmarkSubStrRange-8                  13053810                90.7 ns/op             0 B/op          0 allocs/op
BenchmarkSubStrDecodeRuneInString-8     11359845               103 ns/op               0 B/op          0 allocs/op
BenchmarkSubStrRuneIndexInString-8      14555875                81.6 ns/op             0 B/op          0 allocs/op
BenchmarkSubStrRuneSubString-8          14257446                83.0 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/thinkeridea/go-extend/exunicode/exutf8/benchmark     7.238s
```
