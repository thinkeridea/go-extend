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
package expprof 这是从 net/http/pprof 包复制过来的，做了一些调整
我希望核心的逻辑不变，保持和标准库一致的功能，会调整使用的接口，使之易于控制。

我为什么会有这样的想法，主要源于自己想对程序的性能做分析，但是每次都需要调整程序并重新发布，
这极大的浪费了我的时间，在调整完之后我需要调整代码，去除 net/http/pprof 包，因为该包会暴露默认路由，
这容易导致三方服务来分析我的程序，是极度危险的情况。

我想实现两个功能来解决这一问题：

1、接口路由可以由用户提供，而不是在这个包里定义

2、接口可以设定开放和关闭，并且可以设定定时关闭，以避免调试结束接口还可以访问。
当然我想开放与关闭的接口可以由用户控制，并且可以限定只有内网可以设置（这个只能用户自己实现），这样该功能是安全的。

疑惑：我不知道这么做是否真的好，但是我感觉自己需要这个功能，暂时我还不清楚 pprof 包采集数据的原理，是否会一直采集数据，
采集数据是否会对程序运行有影响，这让我有些迷茫，我简单查看了源码，貌似有些数据程序是一直采集的，但是有些数据是访问过指定功能后开始持续采集的，
像 CPUProfile 是可以在程序内开关的，这导致我无法确定接口一直可用是否会对程序运行有影响，但是这个功能确实非常有用，随着我的学习这些疑惑应该会逐步解决，
所以是否使用该包由用户自己决定，如果你了解这些可以及时联系我进行交流。

原生 net/http 使用示例：

	expprof.RoutePrefix = "/debug/"
	http.HandleFunc(expprof.RoutePrefix, expprof.ServeHTTP)
	http.HandleFunc(expprof.RoutePrefix+"state", func(w http.ResponseWriter, r *http.Request) {
		if !exnet.HasLocalIPddr(exnet.ClientIP(r)) {
			w.WriteHeader(http.StatusNotFound)
		}

		s, err := strconv.Atoi(r.FormValue("s"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("错误的状态"))
		}

		switch s {
		case int(expprof.CloseApi):
			expprof.SetApiState(expprof.CloseApi)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("close api"))
		case int(expprof.OpenApi):
			expprof.SetApiState(expprof.OpenApi)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("open api"))
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("无效的状态"))
		}
	})

	log.Println(http.ListenAndServe("localhost:6060", nil))


gin 使用示例：

    expprof.RoutePrefix = "/debug/"
	g := gin.Default()
	g.GET(expprof.RoutePrefix+"*cmd", gin.WrapF(expprof.ServeHTTP))
	g.GET("/debug_state", func(c *gin.Context) {
		if !exnet.HasLocalIPddr(exnet.ClientIP(c.Request)) {
			c.Status(http.StatusNotFound)
			c.Abort()
		}
	}, func(c *gin.Context) {
		s, err := strconv.Atoi(c.Query("s"))
		if err != nil {
			c.String(http.StatusBadRequest, "错误的状态")
		}

		switch s {
		case int(expprof.CloseApi):
			expprof.SetApiState(expprof.CloseApi)
			c.String(http.StatusBadRequest, "close api")
		case int(expprof.OpenApi):
			expprof.SetApiState(expprof.OpenApi)
			c.String(http.StatusOK, "open api")
		default:
			c.String(http.StatusBadRequest, "错误的状态")
		}
	})

	g.Run(":6060")


Then use the pprof tool to look at the heap profile:

	go tool pprof http://localhost:6060/debug/pprof/heap

Or to look at a 30-second CPU profile:

	go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

Or to look at the goroutine blocking profile, after calling
runtime.SetBlockProfileRate in your program:

	go tool pprof http://localhost:6060/debug/pprof/block

Or to collect a 5-second execution trace:

	wget http://localhost:6060/debug/pprof/trace?seconds=5

Or to look at the holders of contended mutexes, after calling
runtime.SetMutexProfileFraction in your program:

	go tool pprof http://localhost:6060/debug/pprof/mutex

To view all available profiles, open http://localhost:6060/debug/pprof/
in your browser.

For a study of the facility in action, visit

	https://blog.golang.org/2011/06/profiling-go-programs.html
*/
package expprof
