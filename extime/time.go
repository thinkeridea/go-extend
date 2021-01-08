package extime

import "time"

// 通用的时间单位持续时长，这里泛指国际基础单位制（民用日）所理解的时间，不考虑夏时制，不用作科学与天文学
// 例如 一天 24小时，一周 168小时，7天
const (
	// Nanosecond 纳秒，作为最基础的单位
	Nanosecond time.Duration = 1

	// Microsecond 微妙，表示1微妙持续的纳秒时长
	Microsecond = 1000 * Nanosecond

	// Millisecond 毫秒，表示1毫秒持续的纳秒时长
	Millisecond = 1000 * Microsecond

	// Second 秒，表示1秒持续的纳秒时长
	Second = 1000 * Millisecond

	// Minute 分钟，表示1分钟持续的纳秒时长
	Minute = 60 * Second

	// Hour 小时，表示1小时持续的纳秒时长
	Hour = 60 * Minute

	// Day 天，表示1天持续的纳秒时长
	// 这里不考虑夏时制问题，泛指国际基础单位制（民用日）所理解的时间
	Day = 24 * Hour

	// Week 周, 表示1周持续的纳秒时长
	// 这里不考虑夏时制问题，泛指国际基础单位制（民用日）所理解的时间
	Week = 7 * Day
)
