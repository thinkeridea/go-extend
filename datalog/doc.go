// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

/*
Package datalog 用于辅助拼接类 csv 格式化数据日志的组件。

Record 作为一行记录或者一组数据，可以简化数据字段拼接过程及避免低级错误，
详细可以参考博客：https://blog.thinkeridea.com/201907/go/csv_like_data_logs.html

使用示例：

const (
	LogVersion = "v1.0.0"
)
const (
	LogVer = iota
	LogTime
	LogUid
	LogUserName
	LogFriends

	LogFieldNumber
)

func main() {
	var w bytes.Buffer
	pool := datalog.NewRecordPool(LogFieldNumber)

	r := pool.Get().(datalog.Record)
	r[LogVer] = LogVersion
	r[LogTime] = time.Now().Format("2006-01-02 15:04:05")
	r[LogUid] = "Uid"
	r[LogUserName] = "UserNmae"

	data := r.Join(datalog.FieldSep, datalog.NewLine)
	r.Clean()
	pool.Put(r)

	if _, err := w.Write(data); err != nil {
		panic(err)
	}
}

这会输出：（因为分隔符是不可见字符，下面使用,代替字段分隔符，使用;\n代替换行符， 使用/代替数组字段分隔符，是-代替数组分隔符）

'v1.0.0,2019-07-18,11:39:09,Uid,UserNmae,;\n'

往往我们会先定义一组常量， 用来标记字段的顺序，常量标识数据位置，保证我们不会因为一些检查或者忘记记录某个字段而导致整个日志格式混乱，
没有赋值的 LogFriends 依旧有一个占位符，确保日志格式完好。

使用 pool 可以很好的利用内存，不会带来过多的内存分配，而且 Record 的每个字段值都是字符串，简单的赋值并不会带来太大的开销。
使用 Record.Join 可以高效的连接一行日志记录，便于我们快速的写入的日志文件中。


有时候也并非都是记录一些单一的值，比如上面 LogFriends 会记录当前记录相关的朋友信息，这可能是一组数据，
datalog 也提供了一些简单的辅助函数，可以结合下面的实例实现：

const (
	LogVersion = "v1.0.0"
)
const (
	LogVer = iota
	LogTime
	LogUid
	LogUserName
	LogFriends

	LogFieldNumber
)

const (
	LogFriendUid = iota
	LogFriendUserName

	LogFriendFieldNumber
)

func main() {
	var w bytes.Buffer
	pool := datalog.NewRecordPool(LogFieldNumber)
	frPool := datalog.NewRecordPool(LogFriendFieldNumber)

	r := pool.Get().(datalog.Record)
	r[LogVer] = LogVersion
	r[LogTime] = time.Now().Format("2006-01-02 15:04:05")
	r[LogUid] = "Uid"
	r[LogUserName] = "UserNmae"

	r[LogFriends] = GetLogFriends(rand.Intn(3), frPool)
	data := r.Join(datalog.FieldSep, datalog.NewLine)
	r.Clean()
	pool.Put(r)

	if _, err := w.Write(data); err != nil {
		panic(err)
	}

	fmt.Println("'" + w.String() + "'")
}

func GetLogFriends(friendNum int, pool *sync.Pool) string {
	fs := datalog.NewRecord(friendNum)
	fr := pool.Get().(datalog.Record)
	for i := 0; i < friendNum; i++ {
		fr[LogFriendUid] = "FUid"
		fr[LogFriendUserName] = "FUserName"
		fs[i] = fr.ArrayFieldJoin(datalog.ArrayFieldSep, datalog.ArraySep)
	}
	fr.Clean()
	pool.Put(fr)

	return fs.ArrayJoin(datalog.ArraySep)
}

这会输出：（因为分隔符是不可见字符，下面使用,代替字段分隔符，使用;\n代替换行符， 使用/代替数组字段分隔符，是-代替数组分隔符）

'v1.0.0,2019-07-18,11:39:09,Uid,UserNmae,FUid/FUserName-FUid/FUserName;\n'

这样在解析时可以把某一字段当做数组解析，这极大的极大的提高了数据日志的灵活性，
但是并不建议使用过多的层级，数据日志应当清晰简洁，但是有些特殊场景可以使用一层嵌套。
*/
package datalog
